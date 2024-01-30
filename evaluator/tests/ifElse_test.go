package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/object"
)

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if true: { 10 }", 10},
		{"if false: { 10 }", &object.Null{}},
		{"if 1: { 10 }", 10},
		{"if 1 < 2: { 10 }", 10},
		{"if 1 > 2: { 10 }", &object.Null{}},
		{"if 1 > 2: { 10 } else { 20 }", 20},
		{"if 1 < 2: { 10 } else { 20 }", 10},
		{"if 0: {10} else {5}", 5},
		{"if null: {10}", &object.Error{Msg: "identifier not found: null"}},
		{"if 0:{10}", &object.Null{}},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expression := val.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expression))
		case *object.Null:
			testNullObject(t, evaluated)
		case *object.Error:
			testErrorObject(t, evaluated, expression.Msg)
		default:
			t.Errorf("Unknown type %T", expression)
		}
	}
}
