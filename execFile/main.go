package execFile

import (
	"fmt"
	"os"

	"github.com/Mostafa-DE/delang/evaluator"
	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/object"
	"github.com/Mostafa-DE/delang/parser"
)

func getFileContent() []byte {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file to run")
		os.Exit(1)
	}

	filename := os.Args[1]
	content, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return content
}

func Run() {
	fileContent := getFileContent()

	l := lexer.New(string(fileContent))
	p := parser.New(l)

	program := p.ParseProgram()

	env := object.NewEnvironment()

	if parserErrors(p) {
		return
	}

	eval := evaluator.Eval(program, env)

	if eval != nil {
		fmt.Println(eval.Inspect())
	}
}

func parserErrors(p *parser.Parser) bool {
	if len(p.Errors()) != 0 {
		fmt.Println("Error parsing program:")
		fmt.Println(p.Errors()[0])

		return true
	}

	return false
}
