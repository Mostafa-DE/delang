package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/parser"
)

func TestParseArray(t *testing.T) {
	input := "[1, 2, 2 * 2];"

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	array, ok := stmt.Expression.(*ast.Array)
	if !ok {
		t.Fatalf("exp not *ast.Array. got=%T", stmt.Expression)
	}

	if len(array.Elements) != 3 {
		t.Fatalf("len(array.Elements) not 3. got=%d", len(array.Elements))
	}

	testIntegerLiteral(t, array.Elements[0], 1)
	testIntegerLiteral(t, array.Elements[1], 2)
	testInfixExpression(t, array.Elements[2], 2, "*", 2)
}

func TestParseIndexExpression(t *testing.T) {
	input := "deArray[1 + 1];"

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	indexExp, ok := stmt.Expression.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("exp not *ast.IndexExpression. got=%T", stmt.Expression)
	}

	testIdentifier(t, indexExp.Ident, "deArray")
	testInfixExpression(t, indexExp.Index, 1, "+", 1)
}
