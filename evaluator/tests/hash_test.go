package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/evaluator"
	"github.com/Mostafa-DE/delang/object"
)

func TestHashLiterals(t *testing.T) {
	input := `
		let two = "two";
		
		{
			"one": 10 - 9,
			two: 1 + 1,
			"thr" + "ee": 6 / 2,
			4: 4,
			true: 5,
			false: 6
		}
		
	`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Hash)

	if !ok {
		t.Fatalf("Eval didn't return Hash. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey():      4,
		evaluator.TRUE.HashKey():                   5,
		evaluator.FALSE.HashKey():                  6,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]

		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}

		testIntegerObject(t, pair.Value, expectedValue)
	}
}

func TestHashIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`{"name": "DELANG"}["name"]`,
			"DELANG",
		},
		{
			`{"name": "DELANG"}["age"]`,
			&object.Null{},
		},
		{
			`let key = "name"; {"name": "DELANG"}[key]`,
			"DELANG",
		},
		{
			`{}["name"]`,
			&object.Null{},
		},
		{
			`{5: 5}[5]`,
			5,
		},
		{
			`{true: 5}[true]`,
			5,
		},
		{
			`{false: 5}[false]`,
			5,
		},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)

		switch expected := val.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			testStringObject(t, evaluated, expected)
		case *object.Null:
			testNullObject(t, evaluated)

		default:
			t.Errorf("Unknown type %T", expected)
		}
	}
}
