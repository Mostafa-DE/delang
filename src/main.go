package main

import (
	"fmt"
	"os"
	"os/user"
	"repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hi %s! Welcome to DE v0.0.9\n", user.Username)
	fmt.Printf("Type '.exit' to exit the REPL.\n")
	fmt.Printf("Type '.help' to see a list of commands.\n")

	repl.StartSession(os.Stdin, os.Stdout)
}
