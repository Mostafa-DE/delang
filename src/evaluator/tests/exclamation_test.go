package tests

import "testing"

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
