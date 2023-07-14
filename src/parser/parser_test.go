package parser

import (
	"testing"
	"ast"
	"lexer"
)


func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("Parser error: %q", msg)
	}

	t.FailNow()
}


func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let num = 1234;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParserProgram()

	checkParserErrors(t, p)

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


func TestReturnStatements(t *testing.T) {
	input := `
		return 5;
		return 10;
		return 1234;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParserProgram()

	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParserProgram() returned nil :( ")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Return statement doesn't contain 3 statements. got=%d", len(program.Statements))
	}

	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement) // type assertion to make sure we have a return statement

		if !ok {
			t.Errorf("statement not *ast.ReturnStatement. got=%T", statement)
			continue
		}

		if returnStatement.TokenLiteral() != "return" {
			t.Errorf("returnStatement.TokenLiteral not 'return'. got=%q", returnStatement.TokenLiteral())
		}
	}
}