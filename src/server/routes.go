package server

import (
	"encoding/json"
	"net/http"

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

	requestBody := prepareReqBody(req)

	l := lexer.New(requestBody.Code)
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

	evaluated := evaluator.Eval(program, env)

	response = map[string]string{
		"data": evaluated.Inspect(),
	}

	resW.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resW).Encode(response)
}
