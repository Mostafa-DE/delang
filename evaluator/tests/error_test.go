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
		{
			`
				5 + true;
			`,
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			`
				5 + true; 5;
			`,
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			`
				-true;
			`,
			"unknown operator: -BOOLEAN",
		},
		{
			`
				true + false;
			`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
				5;
				true + false;
				5;
			`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
				if 10 > 1: {
					if 10 > 1: {
						return true + false;
					}
				}
			`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
				de;
			`,
			"identifier not found: de",
		},
		{
			`
				"Hello" - "World";
			`,
			"unknown operator: STRING - STRING",
		},
		{
			`
				const a = 5;
				const a = 6;
			`,
			"Cannot redeclare constant 'a'",
		},
		{
			`
				const f = fun(x, y) {
					return x + y;
				}
				const f = 10;
			`,
			"Cannot redeclare constant 'f'",
		},
		{
			`
				const PI = 3;
				PI = 4;
			`,
			"Cannot reassign constant 'PI'",
		},
		{
			`
				const PI = 3;
				let PI = 4;
			`,
			"Cannot reassign constant 'PI'",
		},
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
