package main

import (
	"fmt"
	"os"
	"os/user"
	"repl"
)

func main() {
	user, err := user.Current()

	if err != nil {panic(err)}

	fmt.Printf("Hi %s! Welcome to Test v1.0.0\n", user.Username)

	repl.StartSession(os.Stdin, os.Stdout)
}
