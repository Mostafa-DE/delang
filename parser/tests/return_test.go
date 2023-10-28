package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
)

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"return 5;", 5},
		{"return 10;", 10},
		{"return 1234;", 1234},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)

		if program == nil {
			t.Fatalf("ParseProgram() returned nil :( ")
		}

		statement, ok := program.Statements[0].(*ast.ReturnStatement)
		if !ok { // type assertion to make sure we have a return statement
			t.Errorf("statement not *ast.ReturnStatement. got=%T", statement)
			continue
		}

		if statement.TokenLiteral() != "return" {
			t.Errorf("returnStatement.TokenLiteral not 'return'. got=%q", statement.TokenLiteral())
		}

		if !testLiteralExpression(t, statement.ReturnValue, val.expectedValue) {
			return
		}

	}
}
