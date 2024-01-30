package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/object"
)

func TestArray(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d",
			len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
}

func TestArrayIndexExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[1, 2, 3][0];",
			1,
		},
		{
			"[1, 2, 3][1];",
			2,
		},
		{
			"[1, 2, 3][2];",
			3,
		},
		{
			`
				let i = 0;
				return [1][i];
			`,
			1,
		},
		{
			"[1, 2, 3][1 + 1];",
			3,
		},
		{
			`
				let myArray = [1, 2, 3];
				return myArray[2];
			`,
			3,
		},
		{
			`
				let myArray = [1, 2, 3];
			 	return myArray[0] + myArray[1] + myArray[2];
			`,
			6,
		},
		{
			`
				let arr = [1, 2, 3];
				let i = arr[0];

				arr[i];
			`,
			2,
		},
		{
			`
				let arr = [1, 2, [1, 2], 3];
				arr[2][1];
			`,
			2,
		},
		{
			"[1, 2, 3][3];",
			&object.Error{Msg: "Index out of range"},
		},
		{
			"[1, 2, 3][-1];",
			&object.Error{Msg: "Index out of range"},
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
