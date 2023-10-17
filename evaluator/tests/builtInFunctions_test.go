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
				let x = fun() {
					logs("Inside the main function");
				}

				x(fun() {
					logs("Inside the callback function");
				}());

				logs("Outside the main function");
			`,
			[]string{"Inside the callback function", "Inside the main function", "Outside the main function"},
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
			"It should return error if the first argument is not an array",
			`
				push(1, 2);
			`,
			"argument to `push` must be ARRAY, got INTEGER",
		},
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
			[]int{0, 1, 2, 3, 4},
		},
		{
			"It should return an empty array if the given number is 0",
			`
				range(0);
			`,
			[]int{},
		},
		{
			"It should return an empty array if the given number is negative",
			`
				range(-1);
			`,
			[]int{},
		},
		{
			"It should return error if the argument is not an integer",
			`
				range("1");
			`,
			"argument to `range` must be INTEGER, got STRING",
		},
		{
			"It should return error if the number of arguments is not 1",
			`
				range(1, 2);
			`,
			"wrong number of arguments passed to range(). got=2, want=1",
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
