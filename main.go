package main

import (
	"os"
	"strings"

	"github.com/Mostafa-DE/delang/execFile"
	"github.com/Mostafa-DE/delang/repl"
)

func main() {
	if len(os.Args) == 1 {
		repl.StartSession()
	} else {
		pathArr := strings.Split(os.Args[1], "/")

		if pathArr[len(pathArr)-1] == "--version" {
			println("DE v0.0.9")
			println(DE)
			return
		}

		execFile.Run()
	}
}

const DE = `
  ____       ________
 |  _ \     /|_______|
 | | | |   | |
 | | | |   | |_______
 | | | |   | |_______|
 | | | |   | |
 | |_| /   | |_______
 |____/     \|_______|
 `
