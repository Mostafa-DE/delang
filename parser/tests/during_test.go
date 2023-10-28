package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
)

func TestDuringStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			condition struct {
				left     string
				operator string
				right    string
			}
			body string
		}
	}{
		{
			`
				during x < y: { logs("DE!!"); };
			`,
			struct {
				condition struct {
					left     string
					operator string
					right    string
				}
				body string
			}{
				condition: struct {
					left     string
					operator string
					right    string
				}{
					left:     "x",
					operator: "<",
					right:    "y",
				},
				body: `logs(DE!!)`,
			},
		},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		expression, ok := statement.Expression.(*ast.DuringExpression)

		if !ok {
			t.Fatalf("statement.Expression is not ast.DuringExpression. got=%T", statement.Expression)
		}

		left := val.expected.condition.left
		operator := val.expected.condition.operator
		right := val.expected.condition.right

		if !testInfixExpression(t, expression.Condition, left, operator, right) {
			return
		}

		if len(expression.Body.Statements) != 1 {
			t.Errorf("Body is not 1 statements. got=%d\n", len(expression.Body.Statements))
		}

		body, ok := expression.Body.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("Body.Statements[0] is not ast.ExpressionStatement. got=%T", expression.Body.Statements[0])
		}

		if body.String() != val.expected.body {
			t.Errorf("body.String() is not %q. got=%q", val.expected.body, body.String())
		}

	}
}
