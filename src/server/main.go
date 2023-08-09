package server

import (
	"bytes"
	"encoding/json"
	"evaluator"
	"fmt"
	"lexer"
	"net/http"
	"parser"
	"strings"
)

type RequestBody struct {
	Code string `json:"code"`
}

func StartServer() {
	http.HandleFunc("/api/exec", codeExecHandler)

	port := 8000
	fmt.Printf("Server started on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func codeExecHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	var response map[string]string

	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	cleanedBody := strings.ReplaceAll(buf.String(), "\n", "")
	json.Unmarshal([]byte(cleanedBody), &requestBody)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	l := lexer.New(requestBody.Code)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		response = map[string]string{
			"error": p.Errors()[0],
		}

		json.NewEncoder(w).Encode(response)
	}

	evaluated := evaluator.Eval(program)

	response = map[string]string{
		"data": evaluated.Inspect(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
