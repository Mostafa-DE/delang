package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/object"
)

func TestStringLiteral(t *testing.T) {
	input := `"DELANG!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)

	if !ok {
		t.Fatalf("Object is not a string. Got %T (%+v)", evaluated, evaluated)
	}

	if str.Value != "DELANG!" {
		t.Errorf("String has wrong value. Got %q", str.Value)
	}
}

func TestAddStringToNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`"1" + 5`, "15"},
		{`5 + "1"`, "51"},
		{`"1" + 5.5`, "15.5"},
		{`5.5 + "1"`, "5.51"},
		{`"5.5" + "5.5"`, "5.55.5"},
		{`"Number " + 5`, "Number 5"},
	}

	for _, val := range tests {
		evaluated := testEval(val.input)
		testStringObject(t, evaluated, val.expected)
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"DE" + " " + "Lang!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)

	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "DE Lang!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}
