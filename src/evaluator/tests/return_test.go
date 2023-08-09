package tests

import "testing"

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{
			`	
				9; 
				return 2 * 5;
				9;
				
			`, 10,
		},
		{
			`
				if (10 > 1) {
					if (10 > 1) {
						return 10;
					}
					
					return 1;
				}

			`, 10,
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testIntegerObject(t, evaluated, val.expected)
	}
}
