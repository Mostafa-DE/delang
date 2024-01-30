package tests

import (
	"testing"
)

func TestSimpleVariableAssign(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			`
				let y = 1;
				y = 2;

				return y;
			`,
			2,
		},
		{
			`
				let foobar = 838383;
				foobar = 123;

				return foobar;
			`,
			123,
		},
		{
			`
				let foobar = 838383;
				foobar = 123;
				foobar = 321;

				return foobar;
			`,
			321,
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testIntegerObject(t, evaluated, int64(val.expected))
	}
}

func TestSimpleDictIndexAssign(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			`
				const dict = {"key": 5};
				dict["key"] = 10;

				return dict["key"];
			`,
			10,
		},
		{
			`
				const dict = {"key": 5};
				dict["key"] = 10;
				dict["key"] = 20;

				return dict["key"];
			`,
			20,
		},
		{
			`
				const dict = {"key": 5};
				dict["key"] = 10;
				dict["key"] = 20;
				dict["key"] = 30;

				return dict["key"];
			`,
			30,
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testIntegerObject(t, evaluated, int64(val.expected))
	}
}

func TestSimpleArrayIndexAssign(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			`
				const arr = [1, 2, 3];
				arr[0] = 10;

				return arr[0];
			`,
			10,
		},
		{
			`
				const arr = [1, 2, 3];
				arr[0] = 10;
				arr[0] = 20;

				return arr[0];
			`,
			20,
		},
		{
			`
				const arr = [1, 2, 3];
				arr[0] = 10;
				arr[0] = 20;
				arr[0] = 30;

				return arr[0];
			`,
			30,
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testIntegerObject(t, evaluated, int64(val.expected))
	}
}

func TestScopAssignWithFor(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			`
				let x = 1;
				for i in [1, 2, 3]: {
					x = i;
				}

				return x;
			`,
			3,
		},
		{
			`
				let x = 1;
				for i in [1, 2, 3]: {
					let x = i;
				}
				
				return x;
			`,
			1,
		},
		{
			`
				let x = 1;
				for i in [1, 2, 3]: {
					let x = i;
					x = i;
				}
				
				return x;
			`,
			1,
		},
		// This behavior is not expected, but it is the current behavior.
		// It should recognize that there is a new variable in the scope after accessing it.
		// It should throw an error saying `cannot access variable before it is declared`.
		{
			`
				let x = 1;
				for i in [1, 2, 3]: {
					x = i;
					let x = i;
				}

				return x;
			`,
			1,
		},
		// Same as above.
		{
			`
				let x = 1;
				for i in [1, 2, 3]: {
					x = i;
					let x = i;
					x = i;
				}

				return x;
			`,
			1,
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testIntegerObject(t, evaluated, int64(val.expected))
	}
}

func TestScopeAssignWithFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			`
				let x = 1;
				let test = fun() {
					x = 2;
				};

				test();
			`,
			2,
		},
		{
			`
				let x = 1;
				let test = fun() {
					let x = 2;
				};

				return x;
			`,
			1,
		},
		{
			`
				let x = 1;
				fun() {
					let x = 2;
					x = 3;
				}();

				return x;
			`,
			1,
		},
		{
			`
				let x = 1;
				fun() {
					x = 2;
					let x = 3;

					return x;
				}();
			`,
			3,
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testIntegerObject(t, evaluated, int64(val.expected))
	}
}

func TestScopeAssignWithDuring(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			`
				let x = 1;
				during x < 3: {
					x = x + 1;
				}

				return x;
			`,
			3,
		},
		{
			`
				let x = 1;
				during x < 10: {
					x = x + 2;
				}

				return x;
			`,
			11,
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testIntegerObject(t, evaluated, int64(val.expected))
	}
}
