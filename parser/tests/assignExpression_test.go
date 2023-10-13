package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/parser"
)

func TestAssignExpression(t *testing.T) {
	input := `x = 5;`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 2, len(program.Statements))
	}

	if program.Statements[0].String() != "x = 5;" {
		t.Fatalf("program.Statements[0] is not 'x = 5;', got=%q", program.Statements[0].String())
	}

	if program.Statements[0].TokenLiteral() != "x" {
		t.Fatalf("program.Statements[0].TokenLiteral() is not 'x', got=%q", program.Statements[0].TokenLiteral())
	}
}

func TestAssignDictKeyExpression(t *testing.T) {
	input := `x["key"] = 5;`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 2, len(program.Statements))
	}

	if program.Statements[0].String() != `(x[key])` {
		t.Fatalf(`program.Statements[0] is not '(x[key])', got=%q`, program.Statements[0].String())
	}

	if program.Statements[0].TokenLiteral() != "x" {
		t.Fatalf("program.Statements[0].TokenLiteral() is not 'x', got=%q", program.Statements[0].TokenLiteral())
	}

	exp, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	indExp, ok := exp.Expression.(*ast.IndexExpression)

	if !ok {
		t.Fatalf("exp not *ast.IndexExpression. got=%T", exp.Expression)
	}

	if indExp.Ident.String() != "x" {
		t.Fatalf("indExp.Ident.String() is not 'x', got=%q", indExp.Ident.String())
	}

	if indExp.Index.String() != "key" {
		t.Fatalf("indExp.Index.String() is not 'key', got=%q", indExp.Index.String())
	}

	if indExp.Value.String() != "5" {
		t.Fatalf("indExp.Value.String() is not '5', got=%q", indExp.Value.String())
	}

}

func TestAssignArrayIndexExpression(t *testing.T) {
	input := `x[1] = 5;`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 2, len(program.Statements))
	}

	if program.Statements[0].String() != `(x[1])` {
		t.Fatalf(`program.Statements[0] is not '(x[1])', got=%q`, program.Statements[0].String())
	}

	if program.Statements[0].TokenLiteral() != "x" {
		t.Fatalf("program.Statements[0].TokenLiteral() is not 'x', got=%q", program.Statements[0].TokenLiteral())
	}

	exp, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	indExp, ok := exp.Expression.(*ast.IndexExpression)

	if !ok {
		t.Fatalf("exp not *ast.IndexExpression. got=%T", exp.Expression)
	}

	if indExp.Ident.String() != "x" {
		t.Fatalf("indExp.Ident.String() is not 'x', got=%q", indExp.Ident.String())
	}

	if indExp.Index.String() != "1" {
		t.Fatalf("indExp.Index.String() is not '1', got=%q", indExp.Index.String())
	}

	if indExp.Value.String() != "5" {
		t.Fatalf("indExp.Value.String() is not '5', got=%q", indExp.Value.String())
	}

}
