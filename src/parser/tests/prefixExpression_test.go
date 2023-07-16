package tests

import (
	"testing"
	"lexer"
	"parser"
	"ast"
	"fmt"
)


func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct{
		input string
		operator string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, val := range prefixTests {
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

		prefixExpression, ok := statement.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("expression not *ast.PrefixExpression. got=%T", statement.Expression)
		}

		if prefixExpression.Operator != val.operator {
			t.Fatalf("prefixExpression.Operator not %s. got=%s", val.operator, prefixExpression.Operator)
		}

		if !testIntegerLiteral(t, prefixExpression.Right, val.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, expression ast.Expression, value int64) bool {
	integerLiteral, ok := expression.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("expression not *ast.IntegerLiteral. got=%T", expression)
		return false
	}

	if integerLiteral.Value != value {
		t.Errorf("integerLiteral.Value not %d. got=%d", value, integerLiteral.Value)
		return false
	}

	if integerLiteral.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integerLiteral.TokenLiteral not %d. got=%s", value, integerLiteral.TokenLiteral())
		return false
	}

	return true
}