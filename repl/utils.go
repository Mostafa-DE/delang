package repl

import (
	"fmt"
	"strings"

	"github.com/Mostafa-DE/delang/parser"
)

func parserErrors(p *parser.Parser) bool {
	if len(p.Errors()) != 0 {
		fmt.Println("Error parsing program:")
		fmt.Println(p.Errors()[0])

		return true
	}

	return false
}

func clearCurrentLine() {
	fmt.Print("\r")
	fmt.Print(strings.Repeat(" ", 80))
	fmt.Print("\r")
}

func moveCursorLeft(n int) {
	fmt.Printf("\033[%dD", n)
}

func moveCursorRight(n int) {
	fmt.Printf("\033[%dC", n)
}
