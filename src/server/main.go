package server

import (
	"fmt"
	"net/http"
)

func initApp() {
	port := getPort()

	initAppRoutes()

	fmt.Printf("Server started on port %s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), enableCORS(http.DefaultServeMux))

	if err != nil {
		fmt.Printf("Something went wrong while starting the server: %s", err)
	}
}

func StartServer() {
	initApp()
}
