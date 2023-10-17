package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/parser"

	"github.com/Mostafa-DE/delang/ast"
)

func TestIntegerExpression(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(p.Errors()) != 0 {
		t.Fatalf("Parser has %d errors", len(p.Errors()))
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	testInteger(t, statement.Expression, 5)
}
