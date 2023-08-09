package repl

import (
	"bufio"
	"evaluator"
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
			break
		}

		command := scanner.Text()

		if command == ".exit" {
			break
		}

		if command == ".help" {
			fmt.Printf("Help is on the way!\n")
			continue
		}

		if command == ".clear" {
			clearConsole(output)
			continue
		}

		l := lexer.New(command)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(output, p.Errors())
			continue
		}

		// fmt.Printf("Parsed: %s\n", program.String())

		evaluated := evaluator.Eval(program)

		if evaluated != nil {
			io.WriteString(output, evaluated.Inspect())
			io.WriteString(output, "\n")
		}

	}
}
