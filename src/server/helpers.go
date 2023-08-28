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

	// Remove the new lines
	cleanedBody := strings.ReplaceAll(buf.String(), "\n", "")

	json.Unmarshal([]byte(cleanedBody), &requestBody)

	// Replace the single quotes with double quotes
	// This is because single quotes in Go used to represent runes (characters) not strings
	// But in our language we want to allow single double quotes to represent strings
	requestBody.Code = strings.ReplaceAll(requestBody.Code, "'", "\"")

	return requestBody
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	return port
}

func enableCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Continue with the actual handler
		handler.ServeHTTP(w, r)
	})
}
