package main

import (
	"os"

	"github.com/Mostafa-DE/delang/execFile"
	"github.com/Mostafa-DE/delang/repl"
)

func main() {
	if len(os.Args) == 1 {
		repl.StartSession()
	} else {
		execFile.Run()
	}
}
