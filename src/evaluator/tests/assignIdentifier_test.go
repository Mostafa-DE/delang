package tests

import (
	"testing"
)

func TestAssignExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"let y = true; y = false; y", false},
		{"let foobar = 838383; foobar = 123; foobar;", int64(123)},
		{"let foobar = 838383; foobar = 123; foobar = 321; foobar;", int64(321)},
	}

	for _, val := range tests {
		if intVal, ok := val.expected.(int64); ok {
			testVal := testEval(val.input)

			testIntegerObject(t, testVal, intVal)
		}

		if boolVal, ok := val.expected.(bool); ok {
			testBooleanObject(t, testEval(val.input), boolVal)
		}
	}

}
