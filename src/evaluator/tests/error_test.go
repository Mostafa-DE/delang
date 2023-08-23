package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/object"
)

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"5 + true;", "type mismatch: INTEGER + BOOLEAN"},
		{"5 + true; 5;", "type mismatch: INTEGER + BOOLEAN"},
		{"-true", "unknown operator: -BOOLEAN"},
		{"true + false;", "unknown operator: BOOLEAN + BOOLEAN"},
		{"5; true + false; 5", "unknown operator: BOOLEAN + BOOLEAN"},
		{"if 10 > 1: { true + false; }", "unknown operator: BOOLEAN + BOOLEAN"},
		{"if 10 > 1: { if 10 > 1: { return true + false; } return 1; }", "unknown operator: BOOLEAN + BOOLEAN"},
		{"de", "identifier not found: de"},
		{`"Hello" - "World"`, "unknown operator: STRING - STRING"},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		errObj, ok := evaluated.(*object.Error)

		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Msg != val.expected {
			t.Errorf("wrong error message. expected=%q, got=%q", val.expected, errObj.Msg)
		}
	}

}
