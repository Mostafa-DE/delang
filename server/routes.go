package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

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
	res := make(chan map[string]string)

	go func() {
		res <- codeExec(resW, req)
	}()

	timeout := time.After(5 * time.Second)

	select {
	case <-timeout:
		result := map[string]string{
			"error": "Program execution timeout due to the 5 seconds limit",
		}

		resW.Header().Set("Content-Type", "application/json")
		json.NewEncoder(resW).Encode(result)

		time.AfterFunc(1*time.Second, func() {
			// This is to make sure that the response is sent before exiting.
			os.Exit(0)
		})

	case result := <-res:
		resW.Header().Set("Content-Type", "application/json")
		json.NewEncoder(resW).Encode(result)
	}
}

func codeExec(resW http.ResponseWriter, req *http.Request) map[string]string {
	var response map[string]string

	fileName := createFileToExecFromReqBody(req)

	if fileName == "" {
		response = map[string]string{
			"error": "Something went wrong while executing the code",
		}

		os.Remove(fileName)
		return response
	}

	fileContents, err := ioutil.ReadFile(fileName)

	if err != nil {
		response = map[string]string{
			"error": "Something went wrong while executing the code",
		}

		os.Remove(fileName)

		return response
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
		return response
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

	os.Remove(fileName)

	return response
}
