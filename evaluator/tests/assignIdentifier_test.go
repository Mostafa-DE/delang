package tests

import (
	"testing"
)

func TestAssignExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`
				let y = true;
				y = false;
				return y;
			`,
			false,
		},
		{
			`
				let foobar = 838383;
				foobar = 123;
				return foobar;
			`,
			int64(123),
		},
		{
			`
				let foobar = 838383;
				foobar = 123;
				foobar = 321;
				return foobar;
			`,
			int64(321),
		},
		{
			`
				const dict = {"key": 5};
				dict["key"] = 10;
				return dict["key"];
			`,
			int64(10),
		},
		{
			`
				const dict = {"key": 5};
				dict["key"] = 10;
				dict["key"] = 20;
				return dict["key"];
			`,
			int64(20),
		},
		{
			`
				const arr = [1, 2, 3];
				arr[0] = 10;
				return arr[0];
			`,
			int64(10),
		},
		{
			`
				const arr = [1, 2, 3];
				arr[0] = 10;
				arr[0] = 20;
				return arr[0];
			`,
			int64(20),
		},
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
