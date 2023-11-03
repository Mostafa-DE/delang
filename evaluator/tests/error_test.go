package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/object"
)

// TODO: This should be multiple tests instead of one big test, to cover more cases
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
		{
			`
				for _, num in [1, 2, 3]: {
					logs(_);
				}
			`,
			"identifier not found: _",
		},
		{
			`1 / 0;`,
			"division by zero",
		},
		{
			`1 % 0;`,
			"division by zero",
		},
		{
			`1.1 / 0;`,
			"division by zero",
		},
		{
			`1.1 % 0;`,
			"division by zero",
		},
		{
			`decimal(1) / 0;`,
			"division by zero",
		},
		{
			`
				_getDecimalData["divPrec"] = -1;

				decimal(300) / decimal(1.2121);
			`,
			"Valid range for divPrec is [0 to 28]",
		},
		{
			`
				_getDecimalData["prec"] = -1;

				decimal(300) * decimal(1.2121);
			`,
			"Valid range for prec is [0 to 8]",
		},
		{
			`
				_getDecimalData["prec"] = -1;

				decimal(300) + decimal(1.2121);
			`,
			"Valid range for prec is [0 to 8]",
		},
		{
			`
				_getDecimalData["prec"] = -1;

				decimal(300) - decimal(1.2121);
			`,
			"Valid range for prec is [0 to 8]",
		},
		{
			`
				_getDecimalData["divPrec"] = -1;

				decimal(300) % decimal(1.2121);
			`,
			"Valid range for divPrec is [0 to 28]",
		},
		{
			`
				const len = fun(arr) {
					return arr;
				}
			`,
			"Shadowing of 'len' is not allowed",
		},
		{
			`
				const first = fun(arr) {
					return arr[0];
				}
			`,
			"Shadowing of 'first' is not allowed",
		},
		{
			`
				const last = fun(arr) {
					return arr[len(arr) - 1];
				}
			`,
			"Shadowing of 'last' is not allowed",
		},
		{
			`
				const skipFirst = fun(arr) {
					return skipFirst(arr);
				}
			`,
			"Shadowing of 'skipFirst' is not allowed",
		},
		{
			`
				const skipLast = fun(arr) {
					return skipLast(arr);
				}
			`,
			"Shadowing of 'skipLast' is not allowed",
		},
		{
			`
				const push = "Hello";
			`,
			"Shadowing of 'push' is not allowed",
		},
		{
			`
				const pop = "Hello";
			`,
			"Shadowing of 'pop' is not allowed",
		},
		{
			`
				const logs = "Hello";
			`,
			"Shadowing of 'logs' is not allowed",
		},
		{
			`
				const range = [1, 2, 3];
			`,
			"Shadowing of 'range' is not allowed",
		},
		{
			`
				const decimal = 10;
			`,
			"Shadowing of 'decimal' is not allowed",
		},
		{
			`
				const typeof = "integer";
			`,
			"Shadowing of 'typeof' is not allowed",
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
