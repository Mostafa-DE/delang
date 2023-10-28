package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
)

func TestIntegerExpression(t *testing.T) {
	input := "5;"

	program := parseProgram(t, input)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	testInteger(t, statement.Expression, 5)
}
