package tests

import "testing"

func TestConstStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"const a = 5; a;", 5},
		{
			`
				const a = 5 * 5;
				a;
			`,
			25,
		},
		{
			`
				const a = 5;
				const b = a; 
				const c = a + b + 5; 
				c;
			`,
			15,
		},
		{
			`
				const f = fun(x, y) {
					return x + y;
				}
				f(5, 5);
			`,
			10,
		},
	}
	for _, val := range tests {
		testIntegerObject(t, testEval(val.input), val.expected)
	}
}
