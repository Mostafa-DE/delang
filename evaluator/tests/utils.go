package tests

import (
	"fmt"
	"testing"

	"github.com/Mostafa-DE/delang/evaluator"
	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/object"
	"github.com/Mostafa-DE/delang/parser"
	"github.com/shopspring/decimal"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		throwError(p.Errors()[0])
		for _, err := range p.Errors() {
			fmt.Println(err)
		}
		return &object.Error{Msg: p.Errors()[0]}
	}

	env := object.NewEnvironment()

	return evaluator.Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("Object is not an integer. Got %T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("Object has wrong value. Got %d, expected %d", result.Value, expected)
		return false
	}

	return true
}

func testFloatObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Float)
	if !ok {
		t.Errorf("Object is not a float. Got %T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("Object has wrong value. Got %f, expected %f", result.Value, expected)
		return false
	}

	return true
}

func testDecimalObject(t *testing.T, obj object.Object, expected decimal.Decimal) bool {
	result, ok := obj.(*object.Decimal)
	if !ok {
		t.Errorf("Object is not a decimal. Got %T (%+v)", obj, obj)
		return false
	}

	expectedStr := expected.String()
	resultStr := result.Value.String()

	if resultStr != expectedStr {
		t.Errorf("Object has wrong value. Got %s, expected %s", result.Value, expected)
		return false
	}

	return true
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj.Inspect() != "null" {
		t.Errorf("Object is not NULL. Got %T (%+v)", obj, obj)
		return false
	}

	if obj != evaluator.NULL {
		t.Errorf("Object is not NULL. Got %T (%+v)", obj, obj)
		return false
	}

	return true
}

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("Object is not a string. Got %T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("Object has wrong value. Got %s, expected %s", result.Value, expected)
		return false
	}

	return true
}

func testErrorObject(t *testing.T, obj object.Object, expected string) bool {
	err, ok := obj.(*object.Error)
	if !ok {
		t.Errorf("Object is not an error. Got %T (%+v)", obj, obj)
		return false
	}

	if err.Msg != expected {
		t.Errorf("Wrong error message. Got %s, expected %s", err.Msg, expected)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("Object is not a boolean. Got %T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("Object has wrong value. Got %t, expected %t", result.Value, expected)
		return false
	}

	return true
}

func throwError(format string, a ...interface{}) *object.Error {
	return &object.Error{Msg: fmt.Sprintf(format, a...)}
}
