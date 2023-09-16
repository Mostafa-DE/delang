package main

import (
	"os"

	"github.com/Mostafa-DE/delang/execFile"
	"github.com/Mostafa-DE/delang/repl"
	"github.com/Mostafa-DE/delang/server"
)

func main() {
	if env := os.Getenv("ENV"); env == "PROD" {
		server.StartServer()
		return
	}

	if len(os.Args) == 1 {
		repl.StartSession()
		return
	} else {
		execFile.Run()
	}
}
