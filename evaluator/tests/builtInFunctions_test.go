package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/evaluator"
	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/object"
	"github.com/Mostafa-DE/delang/parser"
)

func TestLenFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{`len(1)`, "argument to `len` not supported, got INTEGER"},
		{`len("one", "two")`, "wrong number of arguments. got=2, want=1"},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))

		case string:
			errObj, ok := evaluated.(*object.Error)

			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if errObj.Msg != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Msg)
			}

		default:
			t.Errorf("Unknown type. got=%T", expected)

		}

	}
}

func TestLogsFunction(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    interface{}
	}{
		{
			"It should log a string in the main environment",
			`
				logs("hello world");
			`,
			[]string{"hello world"},
		},
		{
			"It should log multiple strings in the main environment",
			`
				logs("hello", "world");
			`,
			[]string{"hello", "world"},
		},
		{
			"It should log a string in a local environment",
			`
				let x = fun() {
					logs("hello world!");
				}

				x();
			`,
			[]string{"hello world!"},
		},
		{
			"It should log multiple strings in multiple local environments",
			`
				let x = fun(func) {
					func();
					logs("Inside the main function");
				}

				let y = fun() {
					logs("Inside the callback function");
				}

				x(y);
			`,
			[]string{"Inside the callback function", "Inside the main function"},
		},
		{
			"It should log a string in IIFE local environment",
			`
				fun() {
					logs("hello world!");
				}();
			`,
			[]string{"hello world!"},
		},
		{
			"It should log multiple strings in multiple cases",
			`
				let x = fun(n) {
					logs("Inside the main function");
				}

				x(fun() {
					logs("Inside the callback function");
				}());

				logs("Outside the main function");
			`,
			[]string{"Inside the callback function", "Inside the main function", "Outside the main function"},
		},
		{
			"It should log a null value if the argument is null",
			`
				let x;
				logs(x);
			`,
			[]string{"null"},
		},
		{
			"It should log a value inside self invoking function",
			`
				fun() {
					let x = 1;
					logs(x);
				}();

			`,
			[]string{"1"},
		},
	}

	for _, val := range tests {
		l := lexer.New(val.input)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) > 0 {
			throwError(p.Errors()[0])
			return
		}

		env := object.NewEnvironment()

		evaluated := evaluator.Eval(program, env)

		_, ok := evaluated.(*object.Null)

		if !ok {
			t.Errorf("object is not Null. got=%T (%+v)", evaluated, evaluated)
			continue
		}

		logs, _ := env.Get("bufferLogs")

		_, ok = logs.(*object.Buffer)

		if !ok {
			t.Errorf("The buffer is empty. got=%T", logs)
			continue
		}

		for idx, log := range logs.(*object.Buffer).Value {
			expected := val.expected.([]string)[idx]
			logVal := log.String()
			if logVal != expected {
				t.Errorf("wrong log message. expected=%q, got=%q", expected, logVal)
			}
		}

	}
}

func TestFirstFunction(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    interface{}
	}{
		{
			"It should return the first element of an array",
			`
				first([1, 2, 3]);
			`,
			1,
		},
		{
			"It should return null if the array is empty",
			`
				first([]);
			`,
			"null",
		},
		{
			"It should return error if the argument is not an array",
			`
				first(1);
			`,
			"argument to `first` must be ARRAY, got INTEGER",
		},
		{
			"It should return error if the number of arguments is not 1",
			`
				first([1], [2]);
			`,
			"wrong number of arguments passed to first(). got=2, want=1",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))

		case string:
			errObj, ok := evaluated.(*object.Error)

			if expected == "null" {
				if _, ok := evaluated.(*object.Null); !ok {
					t.Errorf("The return value is not null. got=%T (%+v)", evaluated, evaluated)
					continue
				}
			}

			if expected != "null" {
				if !ok {
					t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
					continue
				}

				if errObj.Msg != expected {
					t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Msg)
				}
			}

		default:
			t.Errorf("Unknown type. got=%T", expected)

		}
	}
}

func TestLastFunction(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    interface{}
	}{
		{
			"It should return the last element of an array",
			`
				last([1, 2, 3]);
			`,
			3,
		},
		{
			"It should return the last element of an array",
			`
				last([1]);
			`,
			1,
		},
		{
			"It should return null if the array is empty",
			`
				last([]);
			`,
			"null",
		},
		{
			"It should return error if the argument is not an array",
			`
				last(1);
			`,
			"argument to `last` must be ARRAY, got INTEGER",
		},
		{
			"It should return error if the number of arguments is not 1",
			`
				last([1], [2]);
			`,
			"wrong number of arguments passed to last(). got=2, want=1",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))

		case string:
			errObj, ok := evaluated.(*object.Error)

			if expected == "null" {
				if _, ok := evaluated.(*object.Null); !ok {
					t.Errorf("The return value is not null. got=%T (%+v)", evaluated, evaluated)
					continue
				}
			}

			if expected != "null" {
				if !ok {
					t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
					continue
				}

				if errObj.Msg != expected {
					t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Msg)
				}
			}

		default:
			t.Errorf("Unknown type. got=%T", expected)

		}
	}
}

func TestSkipFirstFunction(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    interface{}
	}{
		{
			"It should skip the first element of an array and return the rest",
			`
				skipFirst([1, 2, 3]);
			`,
			[]int{2, 3},
		},
		{
			"It should return an empty array if the array has only one element",
			`
				skipFirst([1]);
			`,
			[]int{},
		},
		{
			"It should return null if the array is empty",
			`
				skipFirst([]);
			`,
			"null",
		},
		{
			"It should return error if the argument is not an array",
			`
				skipFirst(1);
			`,
			"argument to `skipFirst` must be ARRAY, got INTEGER",
		},
		{
			"It should return error if the number of arguments is not 1",
			`
				skipFirst([1], [2]);
			`,
			"wrong number of arguments passed to skipFirst(). got=2, want=1",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case []int:
			array, ok := evaluated.(*object.Array)

			if !ok {
				t.Errorf("object is not Array. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if len(array.Elements) != len(expected) {
				t.Errorf("Mismatch number of elements in the array. want=%d, got=%d", len(expected), len(array.Elements))
				continue
			}

			for idx, element := range array.Elements {
				testIntegerObject(t, element, int64(expected[idx]))
			}

		case string:
			errObj, ok := evaluated.(*object.Error)

			if expected == "null" {
				if _, ok := evaluated.(*object.Null); !ok {
					t.Errorf("The return value is not null. got=%T (%+v)", evaluated, evaluated)
					continue
				}
			}

			if expected != "null" {
				if !ok {
					t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
					continue
				}

				if errObj.Msg != expected {
					t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Msg)
				}
			}

		default:
			t.Errorf("Unknown type. got=%T", expected)
		}
	}

}

func TestSkipLastFunction(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    interface{}
	}{
		{
			"It should skip the last element of an array and return the rest",
			`
				skipLast([1, 2, 3]);
			`,
			[]int{1, 2},
		},
		{
			"It should return an empty array if the array has only one element",
			`
				skipLast([1]);
			`,
			[]int{},
		},
		{
			"It should return null if the array is empty",
			`
				skipLast([]);
			`,
			"null",
		},
		{
			"It should return error if the argument is not an array",
			`
				skipLast(1);
			`,
			"argument to `skipLast` must be ARRAY, got INTEGER",
		},
		{
			"It should return error if the number of arguments is not 1",
			`
				skipLast([1], [2]);
			`,
			"wrong number of arguments passed to skipLast(). got=2, want=1",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case []int:
			array, ok := evaluated.(*object.Array)

			if !ok {
				t.Errorf("object is not Array. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if len(array.Elements) != len(expected) {
				t.Errorf("Mismatch number of elements in the array. want=%d, got=%d", len(expected), len(array.Elements))
				continue
			}

			for idx, element := range array.Elements {
				testIntegerObject(t, element, int64(expected[idx]))
			}

		case string:
			errObj, ok := evaluated.(*object.Error)

			if expected == "null" {
				if _, ok := evaluated.(*object.Null); !ok {
					t.Errorf("The return value is not null. got=%T (%+v)", evaluated, evaluated)
					continue
				}
			}

			if expected != "null" {
				if !ok {
					t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
					continue
				}

				if errObj.Msg != expected {
					t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Msg)
				}
			}

		default:
			t.Errorf("Unknown type. got=%T", expected)

		}
	}
}

func TestPushFunction(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    interface{}
	}{
		{
			"It should push an element to the end of an array",
			`
				push([1, 2, 3], 4);
			`,
			[]int{1, 2, 3, 4},
		},
		{
			"It should push an element to an empty array",
			`
				push([], 1);
			`,
			[]int{1},
		},
		{
			"It should push the element to the array with modifying the original array",
			`
				let x = [1, 2, 3];
				push(x, 4);
				return x;
			`,
			[]int{1, 2, 3, 4},
		},
		// TODO: Move this to a separate test
		{
			"It should return error if the first argument is not an array",
			`
				push(1, 2);
			`,
			"argument to `push` must be ARRAY, got INTEGER",
		},
		// TODO: Move this to a separate test
		{
			"It should return error if the number of arguments is not 2",
			`
				push([1], 2, 3);
			`,
			"wrong number of arguments passed to push(). got=3, want=2",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case []int:
			array, ok := evaluated.(*object.Array)

			if !ok {
				t.Errorf("object is not Array. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if len(array.Elements) != len(expected) {
				t.Errorf("Mismatch number of elements in the array. want=%d, got=%d", len(expected), len(array.Elements))
				continue
			}

			for idx, element := range array.Elements {
				testIntegerObject(t, element, int64(expected[idx]))
			}

		case string:
			errObj, ok := evaluated.(*object.Error)

			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if errObj.Msg != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Msg)
			}
		}
	}
}

func TestPopFunction(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    interface{}
	}{
		{
			"It should pop the last element of an array",
			`
				pop([1, 2, 3]);
			`,
			[]int{1, 2},
		},
		{
			"It should return empty array if the array has only one element",
			`
				pop([1]);
			`,
			[]int{},
		},
		{
			"It should return null if the array is empty",
			`
				pop([]);
			`,
			"null",
		},
		{
			"It should return error if the argument is not an array",
			`
				pop(1);
			`,
			"argument to `pop` must be ARRAY, got INTEGER",
		},
		{
			"It should return error if the number of arguments is not 1",
			`
				pop([1], [2]);
			`,
			"wrong number of arguments passed to pop(). got=2, want=1",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case []int:
			array, ok := evaluated.(*object.Array)

			if !ok {
				t.Errorf("object is not Array. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if len(array.Elements) != len(expected) {
				t.Errorf("Mismatch number of elements in the array. want=%d, got=%d", len(expected), len(array.Elements))
				continue
			}

			for idx, element := range array.Elements {
				testIntegerObject(t, element, int64(expected[idx]))
			}

		case string:
			errObj, ok := evaluated.(*object.Error)

			if expected == "null" {
				if _, ok := evaluated.(*object.Null); !ok {
					t.Errorf("The return value is not null. got=%T (%+v)", evaluated, evaluated)
					continue
				}
			}

			if expected != "null" {
				if !ok {
					t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
					continue
				}

				if errObj.Msg != expected {
					t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Msg)
				}
			}

		default:
			t.Errorf("Unknown type. got=%T", expected)

		}
	}
}

func TestRangeFunction(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    interface{}
	}{
		{
			"It should return an array of numbers from 0 to the given number",
			`
				range(5);
			`,
			[]int{0, 1, 2, 3, 4, 5},
		},
		{
			"It should return an empty array if the given number is 0",
			`
				range(0);
			`,
			[]int{0},
		},
		{
			"It should return an empty array if the given number is negative",
			`
				range(-1);
			`,
			[]int{},
		},

		// TODO: Move this to a separate test
		{
			"It should return error if the argument is not an integer",
			`
				range("1");
			`,
			"argument to `range` must be INTEGER",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case []int:
			array, ok := evaluated.(*object.Array)

			if !ok {
				t.Errorf("object is not Array. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if len(array.Elements) != len(expected) {
				t.Errorf("Mismatch number of elements in the array. want=%d, got=%d", len(expected), len(array.Elements))
				continue
			}

			for idx, element := range array.Elements {
				testIntegerObject(t, element, int64(expected[idx]))
			}

		case string:
			errObj, ok := evaluated.(*object.Error)

			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if errObj.Msg != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Msg)
			}
		}
	}
}

func TestShiftFunction(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    interface{}
	}{
		{
			"It should remove the first element of an array",
			`
				shift([1, 2, 3]);
			`,
			[]int{2, 3},
		},
		{
			"It should return an empty array if the array has only one element",
			`
				shift([1]);
			`,
			[]int{},
		},
		{
			"It should return null if the array is empty",
			`
				shift([]);
			`,
			"null",
		},
		{
			"It should return error if the argument is not an array",
			`
				shift(1);
			`,
			"argument to `shift` must be ARRAY, got INTEGER",
		},
		{
			"It should return error if the number of arguments is not 1",
			`
				shift([1], [2]);
			`,
			"wrong number of arguments passed to shift(). got=2, want=1",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case []int:
			array, ok := evaluated.(*object.Array)

			if !ok {
				t.Errorf("object is not Array. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if len(array.Elements) != len(expected) {
				t.Errorf("Mismatch number of elements in the array. want=%d, got=%d", len(expected), len(array.Elements))
				continue
			}

			for idx, element := range array.Elements {
				testIntegerObject(t, element, int64(expected[idx]))
			}

		case string:
			errObj, ok := evaluated.(*object.Error)

			if expected == "null" {
				if _, ok := evaluated.(*object.Null); !ok {
					t.Errorf("The return value is not null. got=%T (%+v)", evaluated, evaluated)
					continue
				}
			}

			if expected != "null" {
				if !ok {
					t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
					continue
				}

				if errObj.Msg != expected {
					t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Msg)
				}
			}

		default:
			t.Errorf("Unknown type. got=%T", expected)
		}
	}
}

func TestUnshiftFunction(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    interface{}
	}{
		{
			"It should add an element to the beginning of an array",
			`
				unshift([1, 2, 3], 0);
			`,
			[]int{0, 1, 2, 3},
		},
		{
			"It should add an element to an empty array",
			`
				unshift([], 1);
			`,
			[]int{1},
		},
		{
			"It should add the element to the array with modifying the original array",
			`
				let x = [1, 2, 3];
				unshift(x, 0);
				return x;
			`,
			[]int{0, 1, 2, 3},
		},
		{
			"It should return error if the first argument is not an array",
			`
				unshift(1, 2);
			`,
			"argument to `unshift` must be ARRAY, got INTEGER",
		},
		{
			"It should return error if the number of arguments is not 2",
			`
				unshift([1], 2, 3);
			`,
			"wrong number of arguments passed to unshift(). got=3, want=2",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case []int:
			array, ok := evaluated.(*object.Array)

			if !ok {
				t.Errorf("object is not Array. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if len(array.Elements) != len(expected) {
				t.Errorf("Mismatch number of elements in the array. want=%d, got=%d", len(expected), len(array.Elements))
				continue
			}

			for idx, element := range array.Elements {
				testIntegerObject(t, element, int64(expected[idx]))
			}

		case string:
			errObj, ok := evaluated.(*object.Error)

			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if errObj.Msg != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Msg)
			}
		}
	}
}

func TestDelFunction(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    interface{}
	}{
		{
			"It should delete the given key from the object",
			`
				let x = { "name": "Mostafa", "age": 25 };
				del(x, "name");
				return x;
			`,
			map[string]interface{}{"age": 25},
		},
		{
			"It should delete the given key from the object with modifying the original object",
			`
				let x = { "name": "Mostafa", "age": 25 };
				del(x, "name");
				return x;
			`,
			map[string]interface{}{"age": 25},
		},
		{
			"It should return error if the first argument is not an object",
			`
				del(1, "name");
			`,
			"first argument to `del` must be HASH, got INTEGER",
		},
		{
			"It should return error if the number of arguments is not 2",
			`
				del({ "name": "Mostafa" }, "name", "age");
			`,
			"wrong number of arguments passed to del(). got=3, want=2",
		},
		{
			"It should return error if the key is not hashable",
			`
				del({ "name": "Mostafa" }, [1, 2, 3]);
			`,
			"unusable as hash key: ARRAY",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case map[string]interface{}:
			object, ok := evaluated.(*object.Hash)

			if !ok {
				t.Errorf("object is not Hash. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if len(object.Pairs) != len(expected) {
				t.Errorf("Mismatch Hash pairs. expected=%d, got=%d", len(expected), len(object.Pairs))
				continue
			}

		case string:
			errObj, ok := evaluated.(*object.Error)

			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if errObj.Msg != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Msg)
			}
		}
	}
}

func TestTypeof(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`
				typeof(1);
			`,
			"INTEGER",
		},
		{
			`
				typeof("hello");
			`,
			"STRING",
		},
		{
			`
				typeof(true);
			`,
			"BOOLEAN",
		},
		{
			`
				typeof([1, 2, 3]);
			`,
			"ARRAY",
		},
		{
			`
				typeof({ "name": "Mostafa" });
			`,
			"HASH",
		},
		{
			`
				typeof(fun() {});
			`,
			"FUNCTION",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testStringObject(t, evaluated, val.expected)
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`
				copy([1, 2, 3]);
			`,
			[]int{1, 2, 3},
		},
		{
			`
				const arr = [1, 2, 3];
				copy(arr);
			`,
			[]int{1, 2, 3},
		},
		{
			`
				const arr = [1, 2, 3];
				const arr2 = copy(arr);
				arr2[0] = 0;

				return arr;
			`,
			[]int{1, 2, 3},
		},
		{
			`
				const arr1 = [1, 2, 3];
				const arr2 = arr1;

				return arr1 == arr2;
			`,
			true,
		},
		{
			`
				const arr1 = [1, 2, 3];
				const arr2 = copy(arr1);

				return arr1 == arr2;
			`,
			false,
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case []int:
			array, _ := evaluated.(*object.Array)

			if len(array.Elements) != len(expected) {
				t.Errorf("Mismatch number of elements in the array. want=%d, got=%d", len(expected), len(array.Elements))
				continue
			}

			for idx, element := range array.Elements {
				testIntegerObject(t, element, int64(expected[idx]))
			}

		case bool:
			boolObj, _ := evaluated.(*object.Boolean)
			testBooleanObject(t, boolObj, expected)

		default:
			t.Errorf("Unknown type. got=%T", expected)
		}
	}
}

func TestInt(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`
				int("1");
			`,
			1,
		},
		{
			`
				int(1);
			`,
			1,
		},
		{
			`
				int(1.5);
			`,
			1,
		},
		{
			`
				int(true);
			`,
			1,
		},
		{
			`
				int(false);
			`,
			0,
		},
		{
			`
				const num = decimal("1.5"); 
				int(num);
			`,
			1,
		},
		{
			`
				int("hello");
			`,
			"string argument to `int` not supported, got `hello`",
		},
		{
			`
				int([1, 2, 3]);
			`,
			"argument to `int` not supported, got `[1, 2, 3]`",
		},
		{
			`
				int(1, 1);
			`,
			"wrong number of arguments passed to int(). got=2, want=1",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case int:
			intObj, _ := evaluated.(*object.Integer)
			testIntegerObject(t, intObj, int64(expected))

		case string:
			errObj, _ := evaluated.(*object.Error)
			testErrorObject(t, errObj, expected)

		default:
			t.Errorf("Unknown type. got=%T", expected)
		}
	}
}

func TestFloat(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`
				float("1");
			`,
			float64(1),
		},
		{
			`
				float(1);
			`,
			float64(1),
		},
		{
			`
				typeof(float(1));
			`,
			"FLOAT",
		},
		{
			`
				float(1.5);
			`,
			float64(1.5),
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case float64:
			floatObj, _ := evaluated.(*object.Float)
			testFloatObject(t, floatObj, expected)

		case string:
			strObj, _ := evaluated.(*object.String)
			testStringObject(t, strObj, expected)

		default:
			t.Errorf("Unknown type. got=%T", expected)
		}
	}
}

func TestBool(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			`
				bool("1");
			`,
			true,
		},
		{
			`
				bool(1);
			`,
			true,
		},
		{
			`
				bool(1.5);
			`,
			true,
		},
		{
			`
				bool("hello");
			`,
			true,
		},
		{
			`
				bool([1, 2, 3]);
			`,
			true,
		},
		{
			`
				bool(decimal("1.5"));
			`,
			true,
		},
		{
			`
				bool(true);
			`,
			true,
		},
		{
			`
				bool("");
			`,
			false,
		},
		{
			`
				bool(0);
			`,
			false,
		},
		{
			`
				bool(0.0);
			`,
			false,
		},
		{
			`
				bool(false);
			`,
			false,
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testBooleanObject(t, evaluated, val.expected)
	}
}

func TestStr(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`
				str(1);
			`,
			"1",
		},
		{
			`
				str(1.5);
			`,
			"1.5",
		},
		{
			`
				str(true);
			`,
			"true",
		},
		{
			`
				str(false);
			`,
			"false",
		},
		{
			`
				str([1, 2, 3]);
			`,
			"[1, 2, 3]",
		},
		{
			`
				str({ "lang": "DE!" });
			`,
			"{'lang': 'DE!'}",
		},
		{
			`
				str(decimal("1.5"));
			`,
			"1.5",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testStringObject(t, evaluated, val.expected)
	}
}
