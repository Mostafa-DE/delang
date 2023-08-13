package tests

import "testing"

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if true: { 10 }", 10},
		{"if false: { 10 }", nil},
		{"if 1: { 10 }", 10},
		{"if 1 < 2: { 10 }", 10},
		{"if 1 > 2: { 10 }", nil},
		{"if 1 > 2: { 10 } else { 20 }", 20},
		{"if 1 < 2: { 10 } else { 20 }", 10},
		{"if 0: {10} else {5}", 5},
		{"if null: {10}", "null"},
		{"if 0:{10}", "null"},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		integer, ok := val.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}
