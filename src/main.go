package main

import (
	"log"
	"os"

	"github.com/Mostafa-DE/delang/execFile"
	"github.com/Mostafa-DE/delang/repl"
	"github.com/Mostafa-DE/delang/server"
	"github.com/joho/godotenv"
)

func getStartedEnv() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return os.Getenv("MODE")
}

func main() {
	mode := getStartedEnv()

	if mode == "server" {
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
