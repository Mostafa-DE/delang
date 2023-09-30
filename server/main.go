package server

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func StartServer() {
	port := getPort()

	initAppRoutes()

	fmt.Printf("Server started on port %s\n", port)

	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), enableCORS(http.DefaultServeMux))

	if err != nil {
		fmt.Printf("Something went wrong while starting the server: %s", err)
	}
}
