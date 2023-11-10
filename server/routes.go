package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Mostafa-DE/delang/evaluator"
	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/object"
	"github.com/Mostafa-DE/delang/parser"
)

func initAppRoutes() {
	http.HandleFunc("/api/exec", codeExecHandler)
	http.HandleFunc("/api/examples/", examplesHandler)
}

func examplesHandler(resW http.ResponseWriter, req *http.Request) {
	pathname := req.URL.Path
	exampleNumber := pathname[len("/api/examples/"):]

	absPath, err := filepath.Abs(fmt.Sprintf("server/examples/%s.md", exampleNumber))
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		json.NewEncoder(resW).Encode(map[string]string{
			"error": "Something went wrong while getting the file, please try again later",
		})
		return
	}

	fileContents, err := ioutil.ReadFile(absPath)

	if err != nil {
		fmt.Println("Error reading file:", err)
		json.NewEncoder(resW).Encode(map[string]string{
			"error": "Something went wrong while getting the file, please try again later",
		})
		return
	}

	mds := string(fileContents)
	html := mdToHTML([]byte(mds))

	respose := map[string]string{
		"html": string(html),
	}

	resW.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resW).Encode(respose)
}

func codeExecHandler(resW http.ResponseWriter, req *http.Request) {
	var res map[string]string

	fileName := createFileToExecFromReqBody(req)

	if fileName == "" {
		res = map[string]string{
			"error": "Something went wrong while executing the code",
		}

		os.Remove(fileName)

		resW.Header().Set("Content-Type", "application/json")
		json.NewEncoder(resW).Encode(res)

		return
	}

	fileContents, err := ioutil.ReadFile(fileName)

	if err != nil {
		res = map[string]string{
			"error": "Something went wrong while executing the code",
		}

		os.Remove(fileName)

		resW.Header().Set("Content-Type", "application/json")
		json.NewEncoder(resW).Encode(res)

		return
	}

	fileContentString := string(fileContents)

	l := lexer.New(fileContentString)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		res = map[string]string{
			"error": p.Errors()[0],
		}
		os.Remove(fileName)

		resW.Header().Set("Content-Type", "application/json")
		json.NewEncoder(resW).Encode(res)

		return
	}

	env := object.NewEnvironment()
	env.Set("timeoutLoop", &object.Boolean{Value: true}, false)

	eval := evaluator.Eval(program, env)

	if eval == nil {
		eval = &object.Null{}
	}

	if eval.Type() == object.ERROR_OBJ {
		res = map[string]string{
			"error": eval.Inspect(),
		}

		os.Remove(fileName)

		resW.Header().Set("Content-Type", "application/json")
		json.NewEncoder(resW).Encode(res)

		return
	}

	logs, logsOk := env.Get("bufferLogs")
	timeOutExceeded, timeoutOk := env.Get("timeoutExceeded")

	if !logsOk {
		logs = &object.Buffer{}
	}

	if !timeoutOk {
		timeOutExceeded = &object.Boolean{Value: false}
	}

	res = map[string]string{
		"logs":    logs.Inspect(),
		"data":    eval.Inspect(),
		"timeout": fmt.Sprintf("%t", timeOutExceeded.(*object.Boolean).Value),
	}

	os.Remove(fileName)

	resW.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resW).Encode(res)
}
