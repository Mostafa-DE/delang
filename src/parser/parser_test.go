package parser

import (
	"testing"
	"ast"
	"lexer"
)


func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let num = 1234;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParserProgram()

	if program == nil {
		t.Fatalf("ParserProgram() returned nil :( ")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Let statement doesn't contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"num"},
	}

	for idx, val := range tests {
		statement := program.Statements[idx]

		if !testLetStatement(t, statement, val.expectedIdentifier){
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
		t.Errorf("letStatement.Name.Value not '%s'. got=%s", name, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("letStatement.Name not '%s'. got=%s", name, letStatement.Name)
		return false
	}

	return true
}
