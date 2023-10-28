package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
)

func TestBooleanExpression1(t *testing.T) {
	booleanTests := []struct {
		input    string
		expected bool
	}{
		{"true;", true},
		{"false;", false},
	}

	for _, val := range booleanTests {
		program := parseProgram(t, val.input)

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
