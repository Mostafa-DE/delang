package tests

import (
	"evaluator"
	"lexer"
	"object"
	"parser"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return evaluator.Eval(program)
}
