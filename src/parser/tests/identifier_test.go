package tests

import (
	"testing"
	"ast"
	"lexer"
	"parser"
)

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
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

	identifier, ok := statement.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression not *ast.Identifier. got=%T", statement.Expression)
	}

	if identifier.Value != "foobar" {
		t.Errorf("identifier.Value not %s. got=%s", "foobar", identifier.Value)
	}

	if identifier.TokenLiteral() != "foobar" {
		t.Errorf("identifier.TokenLiteral not %s. got=%s", "foobar", identifier.TokenLiteral())
	}
	
}