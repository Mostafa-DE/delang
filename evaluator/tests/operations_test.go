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

func TestExclamationOperator(T *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!!true", true},
		{"!!false", false},
		{"!10", false},
		{"!!123", true},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testBooleanObject(T, evaluated, val.expected)
	}
}

func TestGreater_Less_EQ_Operation(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`1 > 2`, false},
		{`1 < 2`, true},
		{`1 > 1`, false},
		{`1 < 1`, false},
		{`1 >= 1`, true},
		{`1 <= 1`, true},
		{`1 >= 2`, false},
		{`1 <= 2`, true},
		{
			`
				const a = 10;
				const b = 20;

				return a >= b;
			`,
			false,
		},
		{
			`
				const a = 10;
				const b = 20;
				const arr = [1, 2];

				return a >= b and arr[0] <= arr[1];
			`,
			false,
		},
		{
			`
				const arr = [1, 2];
				push(arr, 3);

				if arr[2] >= arr[1]: {
					return true;
				} else {
					return false;
				}
			`,
			true,
		},
		{
			`
				const arr = [];

				for num in [1, 2, 3]: {
					push(arr, num);
				}

				return arr[0] <= arr[1] and arr[1] < arr[2];
			`,
			true,
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testBooleanObject(t, evaluated, val.expected)
	}
}
