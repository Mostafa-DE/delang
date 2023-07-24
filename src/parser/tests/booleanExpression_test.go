package tests

import (
	"ast"
	"lexer"
	"parser"
	"testing"
)

func TestBooleanExpression(t *testing.T) {
	booleanTests := []struct {
		input    string
		expected bool
	}{
		{"true;", true},
		{"false;", false},
	}

	for _, val := range booleanTests {
		l := lexer.New(val.input)
		p := parser.New(l)

		program := p.ParserProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
		}

		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		testBooleanLiteral(t, statement.Expression, val.expected)
	}
}
