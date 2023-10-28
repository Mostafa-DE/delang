package tests

import (
	"fmt"
	"testing"

	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/parser"

	"github.com/Mostafa-DE/delang/ast"
)

func parseProgram(t *testing.T, input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)

	parseProgram := p.ParseProgram()
	checkParserErrors(t, p)

	if parseProgram == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	return parseProgram
}

func checkParserErrors(t *testing.T, p *parser.Parser) {
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

func testIdentifier(t *testing.T, expression ast.Expression, value string) bool {
	identifier, ok := expression.(*ast.Identifier)
	if !ok {
		t.Errorf("expression not *ast.Identifier. got=%T", expression)
		return false
	}

	if identifier.Value != value {
		t.Errorf("identifier.Value not %s. got=%s", identifier.Value, value)
		return false
	}

	if identifier.TokenLiteral() != value {
		t.Errorf("identifier.TokenLiteral not %s. got=%s", value, identifier.TokenLiteral())
		return false
	}

	return true
}

func testInteger(t *testing.T, expression ast.Expression, value int64) bool {
	integer, ok := expression.(*ast.Integer)
	if !ok {
		t.Errorf("expression not *ast.Integer. got=%T", expression)
		return false
	}

	if integer.Value != value {
		t.Errorf("integer.Value not %d. got=%d", integer.Value, value)
		return false
	}

	if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integer.TokenLiteral not %d. got=%s", value, integer.TokenLiteral())
		return false
	}

	return true
}

func testBooleanLiteral(t *testing.T, expression ast.Expression, value bool) bool {
	booleanLiteral, ok := expression.(*ast.Boolean)
	if !ok {
		t.Errorf("expression not *ast.Boolean. got=%T", expression)
		return false
	}

	if booleanLiteral.Value != value {
		t.Errorf("booleanLiteral.Value not %t. got=%t", booleanLiteral.Value, value)
		return false
	}

	if booleanLiteral.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("booleanLiteral.TokenLiteral not %t. got=%s", value, booleanLiteral.TokenLiteral())
		return false
	}

	return true
}

func testString(t *testing.T, expression ast.Expression, value string) bool {
	stringLiteral, ok := expression.(*ast.StringLiteral)
	if !ok {
		t.Errorf("expression not *ast.String. got=%T", expression)
		return false
	}

	if stringLiteral.Value != value {
		t.Errorf("stringLiteral.Value not %s. got=%s", stringLiteral.Value, value)
		return false
	}

	if stringLiteral.TokenLiteral() != value {
		t.Errorf("stringLiteral.TokenLiteral not %s. got=%s", value, stringLiteral.TokenLiteral())
		return false
	}

	return true
}

func testBoolean(t *testing.T, expression ast.Expression, value bool) bool {
	boolean, ok := expression.(*ast.Boolean)
	if !ok {
		t.Errorf("expression not *ast.Boolean. got=%T", expression)
		return false
	}

	if boolean.Value != value {
		t.Errorf("boolean.Value not %t. got=%t", boolean.Value, value)
		return false
	}

	if boolean.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("boolean.TokenLiteral not %t. got=%s", value, boolean.TokenLiteral())
		return false
	}

	return true
}

func testLiteralExpression(
	t *testing.T,
	expression ast.Expression,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case int:
		return testInteger(t, expression, int64(v))
	case int64:
		return testInteger(t, expression, v)
	case string:
		return testIdentifier(t, expression, v)
	case bool:
		return testBooleanLiteral(t, expression, v)
	}

	t.Errorf("type of expression not handled. got=%T", expression)
	return false
}

func testExpressionStatement(t *testing.T, stmt ast.Statement) *ast.ExpressionStatement {
	exp, ok := stmt.(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("stmt not *ast.ExpressionStatement. got=%T", stmt)
		return nil
	}

	return exp
}

func testInfixExpression(
	t *testing.T,
	expression ast.Expression,
	left interface{},
	operator string,
	right interface{},
) bool {
	infixExpression, ok := expression.(*ast.InfixExpression)
	if !ok {
		t.Errorf("expression not *ast.InfixExpression. got=%T", expression)
		return false
	}

	if !testLiteralExpression(t, infixExpression.Left, left) {
		return false
	}

	if infixExpression.Operator != operator {
		t.Errorf("infixExpression.Operator not '%s'. got='%s'", operator, infixExpression.Operator)
		return false
	}

	if !testLiteralExpression(t, infixExpression.Right, right) {
		return false
	}

	return true
}
