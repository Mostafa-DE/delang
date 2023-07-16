package tests

import (
	"testing"
	"lexer"
	"parser"
	"ast"
)

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParserProgram()
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

	integerLiteral, ok := statement.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("expression not *ast.IntegerLiteral. got=%T", statement.Expression)
	}

	if integerLiteral.Value != 5 {
		t.Errorf("integerLiteral.Value not %d. got=%d", 5, integerLiteral.Value)
	}

	if integerLiteral.TokenLiteral() != "5" {
		t.Errorf("integerLiteral.TokenLiteral not %s. got=%s", "5", integerLiteral.TokenLiteral())
	}
}