package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
)

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, val := range prefixTests {
		program := parseProgram(t, val.input)

		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
		}

		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		testPrefixExpression(t, statement.Expression, val.operator, val.integerValue)
	}
}

func testPrefixExpression(
	t *testing.T,
	expression ast.Expression,
	operator string,
	right interface{},
) bool {
	prefixExpression, ok := expression.(*ast.PrefixExpression)
	if !ok {
		t.Errorf("expression not *ast.PrefixExpression. got=%T", expression)
		return false
	}

	if prefixExpression.Operator != operator {
		t.Errorf("prefixExpression.Operator not '%s'. got='%s'", operator, prefixExpression.Operator)
		return false
	}

	if !testLiteralExpression(t, prefixExpression.Right, right) {
		return false
	}

	return true
}
