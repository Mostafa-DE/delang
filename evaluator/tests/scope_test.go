package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/object"
)

func TestInnerScopeIsolation(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`
				let outer = 1;
				fun() {
					let outer = 2;
				}

				return outer;
			`,
			1,
		},
		{
			`
				for num in range(4): {
					let inner = 2;
				}

				return inner;
			`,
			&object.Error{Msg: "identifier not found: inner"},
		},
		{
			`
				let count = 0;
				during count <= 10: {
					count = count + 1;
					let inner = 1;
				}

				return inner;
			`,
			&object.Error{Msg: "identifier not found: inner"},
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case *object.Error:
			testErrorObject(t, evaluated, expected.Msg)

		default:
			t.Errorf("Unknown type %T", expected)
		}
	}
}

func TestNestedScopes(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`
				let outer = 1;
				fun() {
					let inner = 2;
					return outer + inner;
				}();
			`,
			3,
		},
		{
			`
				let outer = 1;
				fun() {
					let inner = 2;
					fun() {
						return outer + inner;
					}();
				}();
			`,
			3,
		},
		{
			`
				let outer = 1;
				fun() {
					let inner = 2;
					fun() {
						let innerMost = 3;
						return outer + inner + innerMost;
					}();
				}();
			`,
			6,
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case *object.Error:
			testErrorObject(t, evaluated, expected.Msg)

		default:
			t.Errorf("Unknown type %T", expected)
		}
	}
}
