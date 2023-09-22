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
	var response map[string]string

	fileName := createFileToExecFromReqBody(req)

	if fileName == "" {
		response = map[string]string{
			"error": "Something went wrong while executing the code",
		}

		os.Remove(fileName)
		json.NewEncoder(resW).Encode(response)
		return
	}

	fileContents, err := ioutil.ReadFile(fileName)

	if err != nil {
		response = map[string]string{
			"error": "Something went wrong while executing the code",
		}

		os.Remove(fileName)
		json.NewEncoder(resW).Encode(response)
		return
	}

	fileContentString := string(fileContents)

	l := lexer.New(fileContentString)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		response = map[string]string{
			"error": p.Errors()[0],
		}
		os.Remove(fileName)
		json.NewEncoder(resW).Encode(response)
		return
	}

	env := object.NewEnvironment()

	eval := evaluator.Eval(program, env)

	logs, ok := env.Get("bufferLogs")

	if !ok {
		logs = &object.Buffer{}
	}

	response = map[string]string{
		"logs": logs.Inspect(),
		"data": eval.Inspect(),
	}

	resW.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resW).Encode(response)
	os.Remove(fileName)
}
