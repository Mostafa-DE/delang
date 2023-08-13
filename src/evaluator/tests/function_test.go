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
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
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
