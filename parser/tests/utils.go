package tests

import (
	"fmt"
	"testing"

	"github.com/Mostafa-DE/delang/parser"

	"github.com/Mostafa-DE/delang/ast"
)

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

func testIntegerLiteral(t *testing.T, expression ast.Expression, value int64) bool {
	integerLiteral, ok := expression.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("expression not *ast.IntegerLiteral. got=%T", expression)
		return false
	}

	if integerLiteral.Value != value {
		t.Errorf("integerLiteral.Value not %d. got=%d", integerLiteral.Value, value)
		return false
	}

	if integerLiteral.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integerLiteral.TokenLiteral not %d. got=%s", value, integerLiteral.TokenLiteral())
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

func testLiteralExpression(
	t *testing.T,
	expression ast.Expression,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, expression, int64(v))
	case int64:
		return testIntegerLiteral(t, expression, v)
	case string:
		return testIdentifier(t, expression, v)
	case bool:
		return testBooleanLiteral(t, expression, v)
	}

	t.Errorf("type of expression not handled. got=%T", expression)
	return false
}
