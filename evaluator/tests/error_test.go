package tests

import "testing"

func TestMismatchOperations(t *testing.T) {
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
				5 > true;
			`,
			"type mismatch: INTEGER > BOOLEAN",
		},
		{
			`
				5 < true;
			`,
			"type mismatch: INTEGER < BOOLEAN",
		},
		{
			`
				"Hello" - "World";
			`,
			"unknown operator: STRING - STRING",
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
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testErrorObject(t, evaluated, val.expected)
	}
}

func TestIdentifierNotFound(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`
				de;
			`,
			"identifier not found: de",
		},
		{
			`
				const a = 5;
				const b = 6;
				const c = 7;

				a + b + c + d;
			`,
			"identifier not found: d",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testErrorObject(t, evaluated, val.expected)
	}
}

func TestConstantReassign(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
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
		testErrorObject(t, evaluated, val.expected)
	}
}

func TestDivisionByZero(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
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
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testErrorObject(t, evaluated, val.expected)
	}
}

func TestRangeDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
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
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testErrorObject(t, evaluated, val.expected)
	}
}

func TestShadowInitialize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
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
		testErrorObject(t, evaluated, val.expected)
	}
}
