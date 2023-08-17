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
