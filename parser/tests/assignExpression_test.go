package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
)

func TestAssignExpression(t *testing.T) {
	tests := []struct {
		input         string
		expectedIdent string
		expectedValue string
	}{
		{"x = 5", "x", "5"},
		{"y = 10", "y", "10"},
		{"PI = 3.14", "PI", "3.14"},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)

		exp, expOk := program.Statements[0].(*ast.ExpressionStatement)

		if !expOk {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		assignExp, assignOk := exp.Expression.(*ast.AssignExpression)

		if !assignOk {
			t.Fatalf("exp not *ast.AssignExpression. got=%T", exp.Expression)
		}

		ident := assignExp.Ident.Value

		if ident != val.expectedIdent {
			t.Fatalf("ident.String() is not %q, got=%q", val.expectedIdent, ident)
		}
	}
}

func TestAssignDictKeyExpression(t *testing.T) {
	tests := []struct {
		input          string
		expectedIdent  string
		expectedLookup string
		expectedValue  string
	}{
		{`x["key"] = 5`, "x", "key", "5"},
		{`y["key"] = 10`, "y", "key", "10"},
		{`PI["key"] = 3.14`, "PI", "key", "3.14"},
		{`langs[name] = "Delang"`, "langs", "name", "Delang"},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)

		if program.Statements[0].TokenLiteral() != val.expectedIdent {
			t.Fatalf(
				"program.Statements[0].TokenLiteral() is not '%q', got=%q",
				val.expectedIdent,
				program.Statements[0].TokenLiteral(),
			)
		}

		exp, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		indExp, ok := exp.Expression.(*ast.IndexExpression)

		if !ok {
			t.Fatalf("exp not *ast.IndexExpression. got=%T", exp.Expression)
		}

		if indExp.Ident.String() != val.expectedIdent {
			t.Fatalf("indExp.Ident.String() is not '%q', got=%q", val.expectedIdent, indExp.Ident.String())
		}

		if indExp.Index.String() != val.expectedLookup {
			t.Fatalf("indExp.Index.String() is not '%q', got=%q", val.expectedLookup, indExp.Index.String())
		}

		if indExp.Value.String() != val.expectedValue {
			t.Fatalf("indExp.Value.String() is not '%s', got=%q", val.expectedValue, indExp.Value.String())
		}
	}

}

func TestAssignArrayIndexExpression(t *testing.T) {
	tests := []struct {
		input          string
		expectedIdent  string
		expectedLookup string
		expectedValue  string
	}{
		{`x[1] = 5`, "x", "1", "5"},
		{`y[1] = 10`, "y", "1", "10"},
		{`PI[1] = 3.14`, "PI", "1", "3.14"},
		{`langs[num] = "Delang"`, "langs", "num", "Delang"},
		{`x[1 + 1] = 5`, "x", "(1 + 1)", "5"},
		{`x[fun(){return 1;}] = 5`, "x", "fun() {return 1;}", "5"},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)

		if program.Statements[0].TokenLiteral() != val.expectedIdent {
			t.Fatalf(
				"program.Statements[0].TokenLiteral() is not '%q', got=%q",
				val.expectedIdent,
				program.Statements[0].TokenLiteral(),
			)
		}

		exp, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		indExp, ok := exp.Expression.(*ast.IndexExpression)

		if !ok {
			t.Fatalf("exp not *ast.IndexExpression. got=%T", exp.Expression)
		}

		if indExp.Ident.String() != val.expectedIdent {
			t.Fatalf("indExp.Ident.String() is not '%q', got=%q", val.expectedIdent, indExp.Ident.String())
		}

		if indExp.Index.String() != val.expectedLookup {
			t.Fatalf("indExp.Index.String() is not '%q', got=%q", val.expectedLookup, indExp.Index.String())
		}

		if indExp.Value.String() != val.expectedValue {
			t.Fatalf("indExp.Value.String() is not '%s', got=%q", val.expectedValue, indExp.Value.String())
		}
	}
}
