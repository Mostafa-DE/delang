package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/parser"
)

func TestStringLiteralExpression(t *testing.T) {
	input := `"DELANG";`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. Got=%d, Want=%d", len(program.Statements), 1)
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. Got=%T", program.Statements[0])
	}

	stringLiteral, ok := stmt.Expression.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("exp not *ast.StringLiteral. Got=%T", stmt.Expression)
	}

	if stringLiteral.Value != "DELANG" {
		t.Errorf("stringLiteral.Value not %s. Got=%s", "DELANG", stringLiteral.Value)
	}

	if stringLiteral.TokenLiteral() != "DELANG" {
		t.Errorf("stringLiteral.TokenLiteral not %s. Got=%s", "DELANG", stringLiteral.TokenLiteral())
	}
}
