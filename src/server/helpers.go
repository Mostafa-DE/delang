package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

func prepareReqBody(req *http.Request) RequestBody {
	var requestBody RequestBody
	var buf bytes.Buffer

	_, err := buf.ReadFrom(req.Body)

	if err != nil {
		panic(err)
	}

	cleanedBody := strings.ReplaceAll(buf.String(), "\n", "")
	json.Unmarshal([]byte(cleanedBody), &requestBody)

	return requestBody
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	return port
}
