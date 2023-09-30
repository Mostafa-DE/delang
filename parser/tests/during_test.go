package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/parser"
)

func TestDuringStatements(t *testing.T) {
	input := `
		during x < y: { logs("DE!!"); };
	`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	expression, ok := statement.Expression.(*ast.DuringExpression)

	if !ok {
		t.Fatalf("statement.Expression is not ast.DuringExpression. got=%T", statement.Expression)
	}

	if !testInfixExpression(t, expression.Condition, "x", "<", "y") {
		return
	}

	if len(expression.Body.Statements) != 1 {
		t.Errorf("Body is not 1 statements. got=%d\n", len(expression.Body.Statements))
	}
}
