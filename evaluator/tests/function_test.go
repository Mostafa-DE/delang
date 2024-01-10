package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/object"
)

func TestFunctionObject(t *testing.T) {
	input := "fun(x) { x + 2; };"
	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)

	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v",
			fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"

	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{
			`
				let logsTen = fun(x) {
					return x;
				};
					
				logsTen(10);
			
			`, 10,
		},
		{
			`
				let logsTwo = fun(x) {
					return x;
				};
			
				logsTwo(5);
		
			`, 5,
		},
		{
			`
				let mulByTwo = fun(x) {
				 	return x * 2;
				};
			
				mulByTwo(5);
			
			`, 10,
		},
		{
			`
				let add = fun(x, y) {
				 	return x + y;
				};
			
				add(2, 2);
		
			`, 4,
		},
		{
			`
				let add = fun(x, y) {
					 return x + y;
				};
				
				add(2 + 2, add(2, 2));
				
			`, 8,
		},
		{
			`
				fun(x) {
				 	return x;
				}(2)
			
			`, 2,
		},
	}
	for _, val := range tests {
		testIntegerObject(t, testEval(val.input), val.expected)
	}
}

func TestClosures(t *testing.T) {
	input :=
		`
			let test = fun(x) {
				return fun(y) { 
					return x + y;
				};
			};
		
			let apply = test(1);
		
			apply(2);
		`

	testIntegerObject(t, testEval(input), 3)

}

func TestHOF(t *testing.T) {
	input :=
		`

			let add = fun(x, y) {
				return x + y;
			}

			let HOF = fun(f, x, y) {
				return f(x, y);
			}

			HOF(add, 1, 2);
		`

	testIntegerObject(t, testEval(input), 3)
}

func TestCallingWithoutArgs(t *testing.T) {
	tests := []struct {
		explanation string
		input       string
		expected    object.Object
	}{
		{
			explanation: "Calling without args, should return error",
			input: `
				const add = fun(x, y) { return x + y; };
				// const add = fun() { return 1 + 1; };
				add();
			`,
			expected: &object.Error{Msg: "wrong number of arguments: want=2, got=0"},
		},
		{
			explanation: "Calling without args, should return error",
			input: `
				const _logs = fun(x) { return x; };
				_logs();
			`,
			expected: &object.Error{Msg: "wrong number of arguments: want=1, got=0"},
		},
	}
	for _, val := range tests {
		evaluated := testEval(val.input)

		if evaluated.Inspect() != val.expected.Inspect() {
			t.Errorf("Expected %v, got %v", val.expected, evaluated)
		}
	}
}

func TestAccessDataFromOuterScope(t *testing.T) {
	tests := []struct {
		explanation string
		input       string
		expected    object.Object
	}{
		{
			explanation: "Accessing data from outer scope, calling inner function directly",
			input: `
				const outer = fun(x) {
					const inner = fun() {
						return x;
					};

					return inner();
				};

				outer(10);
			`,
			expected: &object.Integer{Value: 10},
		},
		{
			explanation: "Accessing data from outer scope, assigning inner function to a variable and calling it",
			input: `
				const outer = fun() {
					const num = 10;

					return fun() {
						return num;
					};
				};
				
				const f = outer();

				f();
			`,
			expected: &object.Integer{Value: 10},
		},
		{
			explanation: "Accessing data from outer scope, using IIFE",
			input: `
				const outer = fun() {
					const msg = "Hello";
					return fun() {
						return msg;
					};
				};

				outer()();
			`,
			expected: &object.String{Value: "Hello"},
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		if evaluated.Inspect() != val.expected.Inspect() {
			t.Errorf("Expected %v, got %v", val.expected, evaluated)
		}
	}
}

func TestRecursiveFunction(t *testing.T) {
	tests := []struct {
		explanation string
		input       string
		expected    object.Object
	}{
		{
			explanation: "Recursive function",
			input: `
				const factorial = fun(x) {
					if x == 0: {
						return 1;
					}
					return x * factorial(x - 1);
				};

				factorial(5);
			`,
			expected: &object.Integer{Value: 120},
		},
		{
			explanation: "Recursive function",
			input: `
				const factorial = fun(x) {
					if x == 0: {
						return 1;
					}
					return x * factorial(x - 1);
				};

				factorial(10);
			`,
			expected: &object.Integer{Value: 3628800},
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		if evaluated.Inspect() != val.expected.Inspect() {
			t.Errorf("Expected %v, got %v", val.expected, evaluated)
		}
	}
}

func TestFunctionWithNoEndOrStartCurlyBracket(t *testing.T) {
	tests := []struct {
		explanation string
		input       string
		expected    object.Object
	}{
		{
			explanation: "It should return an error if no end curly bracket",
			input: `
				fun(x, y) {
					return x + y;
				
			`,
			expected: &object.Error{Msg: "Function is not closed with '}'"},
		},
		{
			explanation: "It should return an error if no start curly bracket",
			input: `
				fun(x, y)
					return x + y;
				}
		`,
			expected: &object.Error{Msg: "Expected next token to be '{', got 'RETURN' instead"},
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		if evaluated.Inspect() != val.expected.Inspect() {
			t.Errorf("Expected %v, got %v", val.expected, evaluated)
		}

	}
}

// Ignore this test for now
func TestFunctionWithDefaultParameters(t *testing.T) {
	t.Skip()
	tests := []struct {
		explanation string
		input       string
		expected    object.Object
	}{
		{
			explanation: "Function with default parameters",
			input: `
				const add = fun(x, y = 10) {
					return x + y;
				};

				add(5);
			`,
			expected: &object.Integer{Value: 15},
		},
		{
			explanation: "Function with default parameters",
			input: `
				const add = fun(x, y = 10) {
					return x + y;
				};

				add(5, 5);
			`,
			expected: &object.Integer{Value: 10},
		},
		{
			explanation: "Function with default parameters",
			input: `
				const add = fun(x, y = 10) {
					return x + y;
				};

				add(5, 5, 5);
			`,
			expected: &object.Error{Msg: "wrong number of arguments: want=2, got=3"},
		},
		{
			explanation: "Function with default parameters",
			input: `
				const add = fun(x, y = 10) {
					return x + y;
				};

				add(5, 5, 5, 5);
			`,
			expected: &object.Error{Msg: "wrong number of arguments: want=2, got=4"},
		},
		{
			explanation: "Function with default parameters",
			input: `
				const add = fun(x, y = 10) {
					return x + y;
				};

				add();
			`,
			expected: &object.Error{Msg: "wrong number of arguments: want=2, got=0"},
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		if evaluated.Inspect() != val.expected.Inspect() {
			t.Errorf("Expected %v, got %v", val.expected, evaluated)
		}
	}
}
