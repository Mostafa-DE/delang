package server

import (
	"encoding/json"
	"evaluator"
	"lexer"
	"net/http"
	"parser"
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
	}

	evaluated := evaluator.Eval(program)

	response = map[string]string{
		"data": evaluated.Inspect(),
	}

	resW.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resW).Encode(response)
}
