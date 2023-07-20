package tests

import (
	"ast"
	"lexer"
	"parser"
	"testing"
)

func TestInfixExpression(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, val := range infixTests {
		l := lexer.New(val.input)
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

		infixExpression, ok := statement.Expression.(*ast.InfixExpression)

		if !ok {
			t.Fatalf("expression not *ast.InfixExpression. got=%T", statement.Expression)
		}

		if !testIntegerLiteral(t, infixExpression.Left, val.leftValue) {
			return
		}

		if infixExpression.Operator != val.operator {
			t.Fatalf("infixExpression.Operator not '%s'. got='%s'", infixExpression.Operator, val.operator)
		}

		if !testIntegerLiteral(t, infixExpression.Right, val.rightValue) {
			return
		}

	}
}
