package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
)

func TestConstStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"const x = 5;", "x", 5},
		{"const y = 10;", "y", 10},
		{"const num = 1234;", "num", 1234},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := program.Statements[0]

		if !testConstStatement(t, statement, val.expectedIdentifier) {
			return
		}

		value := statement.(*ast.ConstStatement).Value

		if !testLiteralExpression(t, value, val.expectedValue) {
			return
		}

		identifier := statement.(*ast.ConstStatement).Name

		if !testLiteralExpression(t, identifier, val.expectedIdentifier) {
			return
		}
	}
}

func testConstStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "const" {
		t.Errorf("s.TokenLiteral not 'const'. got=%q", s.TokenLiteral())
		return false
	}

	constStatement, ok := s.(*ast.ConstStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if constStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value not '%s'. got=%s", constStatement.Name.Value, name)
		return false
	}

	if constStatement.Name.TokenLiteral() != name {
		t.Errorf("letStatement.Name not '%s'. got=%s", constStatement.Name, name)
		return false
	}

	return true
}
