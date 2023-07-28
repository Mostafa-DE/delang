package tests

import (
	"ast"
	"lexer"
	"parser"
	"testing"
)

func TestCallFunction(t *testing.T) {
	input := `
		add(1, 2 * 3, 4 + 5);
	`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d\n", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T\n", program.Statements[0])
	}

	callFunction, ok := statement.Expression.(*ast.CallFunction)
	if !ok {
		t.Fatalf("statement.Expression is not ast.CallFunction. got=%T\n", statement.Expression)
	}

	if !testIdentifier(t, callFunction.Function, "add") {
		return
	}

	if len(callFunction.Arguments) != 3 {
		t.Fatalf("wrong length of arguments. got=%d\n", len(callFunction.Arguments))
	}

	testLiteralExpression(t, callFunction.Arguments[0], 1)
	testInfixExpression(t, callFunction.Arguments[1], 2, "*", 3)
	testInfixExpression(t, callFunction.Arguments[2], 4, "+", 5)
}

func TestCallFunctionParameters(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams interface{}
	}{
		{input: "add();", expectedParams: []string{}},
		{input: "add(x, y);", expectedParams: []string{"x", "y"}},
		{input: "add(1, 2);", expectedParams: []int{1, 2}},
		{input: "add(1, 2, 3);", expectedParams: []int{1, 2, 3}},
	}

	for _, val := range tests {
		l := lexer.New(val.input)
		p := parser.New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		statement := program.Statements[0].(*ast.ExpressionStatement)
		callFunction := statement.Expression.(*ast.CallFunction)

		switch val.expectedParams.(type) {
		case []string:
			if len(callFunction.Arguments) != len(val.expectedParams.([]string)) {
				t.Errorf(
					"length parameters wrong. want %d, got=%d\n",
					len(val.expectedParams.([]string)), len(callFunction.Arguments),
				)
			}

		case []int:
			if len(callFunction.Arguments) != len(val.expectedParams.([]int)) {
				t.Errorf(
					"length parameters wrong. want %d, got=%d\n",
					len(val.expectedParams.([]int)), len(callFunction.Arguments),
				)
			}

		}
	}
}
