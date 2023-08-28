package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/object"
	"github.com/Mostafa-DE/delang/parser"

	"github.com/Mostafa-DE/delang/evaluator"
)

type RequestBody struct {
	Code string `json:"code"`
}

func initAppRoutes() {
	http.HandleFunc("/api/exec", codeExecHandler)
}

func codeExecHandler(resW http.ResponseWriter, req *http.Request) {
	var response map[string]string

	reqBody := prepareReqBody(req)

	if strings.TrimSpace(reqBody.Code) == "" {
		response = map[string]string{
			"error": "code is required",
		}

		json.NewEncoder(resW).Encode(response)
		return
	}

	l := lexer.New(reqBody.Code)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		response = map[string]string{
			"error": p.Errors()[0],
		}

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
}
