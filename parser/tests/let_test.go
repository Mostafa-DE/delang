package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
)

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let y = 10;", "y", 10},
		{"let num = 1234;", "num", 1234},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := program.Statements[0]

		if !testLetStatement(t, statement, val.expectedIdentifier) {
			return
		}

		value := statement.(*ast.LetStatement).Value

		if !testLiteralExpression(t, value, val.expectedValue) {
			return
		}

		identifier := statement.(*ast.LetStatement).Name

		if !testLiteralExpression(t, identifier, val.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value not '%s'. got=%s", letStatement.Name.Value, name)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("letStatement.Name not '%s'. got=%s", letStatement.Name, name)
		return false
	}

	return true
}
