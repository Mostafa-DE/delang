package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/object"
)

func TestForEval(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{

		{
			`
				const arr = [1, 2, 3];
				const newArr = [];
				for idx, num in arr: {
					logs(num);
					push(newArr, num);
				}

				return newArr;
			`,
			[]int{1, 2, 3},
		},
		{
			`
				const arr = [1, 2, 3, 4, 5];
				const newArr = [];
				for idx, num in arr: {
					if num == 3: {
						break;
					}
					
					logs(num);
					push(newArr, num);
				}

				return newArr;
			`,
			[]int{1, 2},
		},
		{
			`
				const arr = [];

				for _, num in [1, 2, 3, 4, 5]: {
					push(arr, num);
				}

				return arr;
			`,
			[]int{1, 2, 3, 4, 5},
		},
		{
			`
				const arr = [];
				for idx, num in [1, 2, 3, 4, 5]: {
					push(arr, "Index: " + idx + " Number: " + num);
				}

				return arr;
			`,
			[]string{
				"Index: 0 Number: 1",
				"Index: 1 Number: 2",
				"Index: 2 Number: 3",
				"Index: 3 Number: 4",
				"Index: 4 Number: 5",
			},
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		if evaluated.Type() == object.ERROR_OBJ {
			t.Errorf("Expected no error, got %s", evaluated.Inspect())
			continue
		}

		if evaluated != nil {
			switch expected := val.expected.(type) {
			case []int:
				if evaluated.Type() != object.ARRAY_OBJ {
					t.Errorf("Expected iterable %T, got %T", expected, evaluated)
					continue
				}

				arr := evaluated.(*object.Array)
				if len(arr.Elements) != len(expected) {
					t.Errorf("Expected array length %d, got %d", len(expected), len(arr.Elements))
				}

				for i, v := range arr.Elements {
					testIntegerObject(t, v, int64(expected[i]))
				}
			case []string:
				if evaluated.Type() == object.NULL_OBJ {
					t.Errorf("Expected iterable %T, got %T", expected, evaluated)
					continue
				}

				arr := evaluated.(*object.Array)
				if len(arr.Elements) != len(expected) {
					t.Errorf("Expected array length %d, got %d", len(expected), len(arr.Elements))
				}

				for i, v := range arr.Elements {
					testStringObject(t, v, expected[i])
				}
			}
		}
	}
}

func TestForWithString(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`
				let _str = "";
				for char in "Hello World": {
					_str = _str + char;
				}

				return _str;
			`,
			"Hello World",
		},
		{
			`
				let _str = "";
				for char in "Hello World": {
					if char == "o": {
						break;
					}

					_str = _str + char;
				}

				return _str;
			`,
			"Hell",
		},
		{
			`
				const arr = [];
				for idx, num in "Hello World": {
					logs(num);
					push(arr, num);
				}

				return arr;
			`,
			[]string{"H", "e", "l", "l", "o", " ", "W", "o", "r", "l", "d"},
		},
		{
			`
				const arr = [];
				for idx, num in "Hello World": {
					if num == "o": {
						skip;
					}

					if num == "W": {
						skip;
					}

					logs(num);

					push(arr, num);
				}

				return arr;
			`,
			[]string{"H", "e", "l", "l", " ", "r", "l", "d"},
		},
		{
			`
				const arr = [];

				for idx, lang in ["DE", "Go", "Rust", "C++", "Python", "JavaScript"]: {
					if lang == "C++": {
						skip;
					}

					logs(lang);

					push(arr, lang);

				}

				return arr;
			`,
			[]string{"DE", "Go", "Rust", "Python", "JavaScript"},
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		if evaluated.Type() == object.ERROR_OBJ {
			t.Errorf("Expected no error, got %s", evaluated.Inspect())
			continue
		}

		if evaluated != nil {
			switch expected := val.expected.(type) {
			case []string:
				if evaluated.Type() == object.NULL_OBJ {
					t.Errorf("Expected iterable %T, got %T", expected, evaluated)
					continue
				}

				arr := evaluated.(*object.Array)
				if len(arr.Elements) != len(expected) {
					t.Errorf("Expected array length %d, got %d", len(expected), len(arr.Elements))
				}

				for idx, _value := range arr.Elements {
					testStringObject(t, _value, expected[idx])
				}
			case string:
				testStringObject(t, evaluated, expected)
			}
		}
	}
}

func TestForWithStringConcatenation(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`
				let _str = "";
				for idx, char in "Hello World": {
					_str = _str + idx + char;
				}

				return _str;
			`,
			"0H1e2l3l4o5 6W7o8r9l10d",
		},
		{
			`
				let _str = "";
				for idx, char in "Hello World": {
					if char == "o": {
						break;
					}

					_str = _str + idx + char;
				}

				return _str;
			`,
			"0H1e2l3l",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testStringObject(t, evaluated, val.expected)
	}
}

func TestForWithExplicitlyOmittedIndex(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`
				let _str = "";
				for _, char in "Hello World": {
					_str = _str + char;
				}

				return _str;
			`,
			"Hello World",
		},
		{
			`
				let _str = "";
				for _, char in "Hello World": {
					if char == "o": {
						break;
					}

					_str = _str + char;
				}

				return _str;
			`,
			"Hell",
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testStringObject(t, evaluated, val.expected)
	}
}

func TestForWithExplicitlyOmittedIndexAndValue(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Error
	}{
		{
			`
				let _str = "";
				for _, _ in "Hello World": {
					_str = _str + "x";
				}

				return _str;
			`,
			object.Error{Msg: "Cannot use two underscores in for statement"},
		},
		{
			`
				let _str = "";
				for _, _ in "Hello World": {
					if _ == "o": {
						break;
					}

					_str = _str + "x";
				}

				return _str;
			`,
			object.Error{Msg: "Cannot use two underscores in for statement"},
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testErrorObject(t, evaluated, val.expected.Msg)
	}
}

func TestForWithExplicitlyOmittedValue(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Error
	}{
		{
			`
				let _str = "";
				for idx, _ in "Hello World": {
					_str = _str + idx;
				}

				return _str;
			`,
			object.Error{Msg: "Cannot use underscore as a variable identifier in for statement"},
		},
		{
			`
				let _str = "";
				for idx, _ in "Hello World": {
					if _ == "o": {
						break;
					}

					_str = _str + idx;
				}

				return _str;
			`,
			object.Error{Msg: "Cannot use underscore as a variable identifier in for statement"},
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testErrorObject(t, evaluated, val.expected.Msg)
	}
}
