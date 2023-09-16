package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Mostafa-DE/delang/evaluator"
	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/object"
	"github.com/Mostafa-DE/delang/parser"
)

type RequestBody struct {
	Code string `json:"code"`
}

func initAppRoutes() {
	http.HandleFunc("/api/exec", codeExecHandler)
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
