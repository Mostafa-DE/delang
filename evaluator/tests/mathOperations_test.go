package tests

import (
	"testing"
)

func TestIntegerOperationsExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"55", 55},
		{"100", 100},
		{"-50", -50},
		{"-100", -100},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
		{"5 % 2", 1},
		{"5 % 3", 2},
		{"2 % 2", 0},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testIntegerObject(t, evaluated, val.expected)
	}
}

func TestFloatOperationsExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"55.5", 55.5},
		{"100.5", 100.5},
		{"1.111111", 1.111111},
		{"-50.5", -50.5},
		{"-100.5", -100.5},
		{"5.5 + 5.5 + 5.5 + 5.5 - 10.5", 11.5},
		{"2.5 * 2.5 * 2.5 * 2.5 * 2.5", 97.65625},
		{"-50.5 + 100.5 + -50", 0},
		{"5.5 * 2.5 + 11", 24.75},
		{"5.5 + 2.5 * 10.5", 31.75},
		{"20.5 + 2.5 * -10.5", -5.75},
		{"50.5 / 2.5 * 2.5 + 10.5", 61},
		{"2.5 * (5.5 + 10.5)", 40},
		{"3.5 * 3.5 * 3.5 + 10.5", 53.375},
		{"(5.5 + 10.5 * 2.5 + 15.5 / 3) * 2.5", 92.29166666666666},
		{"5.5 % 2.5", 0.5},
		{"5.5 % 3.5", 2},
		{"2.5 % 2.5", 0},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testFloatObject(t, evaluated, val.expected)
	}
}

func TestAddStringToNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`"1" + 5`, "15"},
		{`5 + "1"`, "51"},
		{`"1" + 5.5`, "15.5"},
		{`5.5 + "1"`, "5.51"},
		{`"5.5" + "5.5"`, "5.55.5"},
		{`"Number " + 5`, "Number 5"},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testStringObject(t, evaluated, val.expected)
	}
}

func TestAnd_OrLogical(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"true and true", true},
		{"true and false", false},
		{"false and true", false},
		{"false and false", false},
		{"true or true", true},
		{"true or false", true},
		{"false or true", true},
		{"false or false", false},
		{"1 and 1", 1},
		{"1 and 0", 0},
		{"0 and 1", 1},
		{"0 and 0", 0},
		{"1 or 1", 1},
		{"1 or 0", 1},
		{"0 or 1", 0},
		{"0 or 0", 0},
		{"1 and true", true},
		{"1 and false", false},
		{"0 and true", false},
		{"0 and false", false},
		{"1 or true", true},
		{"1 or false", true},
		{"0 or true", true},
		{"0 or false", false},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch val.expected.(type) {
		case bool:
			testBooleanObject(t, evaluated, val.expected.(bool))
		case int64:
			testIntegerObject(t, evaluated, val.expected.(int64))
		}
	}
}
