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
