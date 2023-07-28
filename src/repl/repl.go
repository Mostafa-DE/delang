package repl

import (
	"bufio"
	"fmt"
	"io"
	"lexer"
	"parser"
)

const PROMPT = ">> "

func StartSession(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		lineText := scanner.Text()

		if lineText == ".exit" {
			return
		}

		if lineText == ".help" {
			fmt.Printf("Help is on the way!\n")
			continue
		}

		if lineText == "clear" {
			clearConsole(output)
			continue
		}

		l := lexer.New(lineText)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(output, p.Errors())
			continue
		}

		io.WriteString(output, program.String())
		io.WriteString(output, "\n")
	}
}
