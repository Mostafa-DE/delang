package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/parser"
)

func TestOperatorPrecedence(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"-1 * 2 + 3", "(((-1) * 2) + 3)"},
		{"-a * b", "((-a) * b)"},
		{"!-a", "(!(-a))"},
		{"a + b + c", "((a + b) + c)"},
		{"a + b - c", "((a + b) - c)"},
		{"a * b * c", "((a * b) * c)"},
		{"a * b / c", "((a * b) / c)"},
		{"a + b / c", "(a + (b / c))"},
		{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)"},
		{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)"},
		{"5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))"},
		{"5 < 4 != 3 > 4", "((5 < 4) != (3 > 4))"},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
		{"true", "true"},
		{"false", "false"},
		{"3 > 5 == false", "((3 > 5) == false)"},
		{"3 < 5 == true", "((3 < 5) == true)"},
		{"1 + (2 + 3) + 4", "((1 + (2 + 3)) + 4)"},
		{"(5 + 5) * 2", "((5 + 5) * 2)"},
		{"2 / (5 + 5)", "(2 / (5 + 5))"},
		{"-(5 + 5)", "(-(5 + 5))"},
		{"!(true == true)", "(!(true == true))"},
		{
			"a + add(b * c) + d",
			"((a + add((b * c))) + d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
		{
			"a * [1, 2, 3, 4][b * c] * d",
			"((a * ([1, 2, 3, 4][(b * c)])) * d)",
		},
		{
			"add(a * b[2], b[1], 2 * [1, 2][1])",
			"add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))",
		},
		{
			"1 % 2",
			"(1 % 2)",
		},
		{
			"true and false",
			"(true and false)",
		},
		{
			"1 and 2",
			"(1 and 2)",
		},
		{
			"1 == 1 and 2 == 2",
			"((1 == 1) and (2 == 2))",
		},
		{
			"1 < 2 and 2 < 3",
			"((1 < 2) and (2 < 3))",
		},
		{
			"1 < 2 and 2 < 3 and 3 < 4",
			"(((1 < 2) and (2 < 3)) and (3 < 4))",
		},
		{
			"true or false",
			"(true or false)",
		},
		{
			"1 or 2",
			"(1 or 2)",
		},
		{
			"1 == 1 or 2 == 2",
			"((1 == 1) or (2 == 2))",
		},
		{
			"1 < 2 or 2 < 3",
			"((1 < 2) or (2 < 3))",
		},
		{
			"1 < 2 or 2 < 3 or 3 < 4",
			"(((1 < 2) or (2 < 3)) or (3 < 4))",
		},
		{
			"1 and 2 or 3",
			"((1 and 2) or 3)",
		},
	}

	for _, val := range tests {
		l := lexer.New(val.input)
		p := parser.New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()

		if actual != val.expected {
			t.Errorf("expected=%q, got=%q", val.expected, actual)
		}
	}
}
