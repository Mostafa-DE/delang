package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/object"
	"github.com/Mostafa-DE/delang/parser"
)

func TestHashStringKeys(t *testing.T) {
	input := `
		{
			"name": "DELANG!",
			"hobby": "Being a programming language",
			"age": 1,
		}
	`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	hash, ok := statement.Expression.(*ast.Hash)

	if !ok {
		t.Fatalf("statement.Expression is not ast.Hash , got=%T", statement.Expression)
	}

	if len(hash.Pairs) != 3 {
		t.Errorf("hash.Pairs has wrong length. got=%d", len(hash.Pairs))
	}

	expected := map[string]interface{}{
		"name":  "DELANG!",
		"hobby": "Being a programming language",
		"age":   1,
	}

	for key, value := range hash.Pairs {
		stringKey, ok := key.(*ast.StringLiteral)

		if !ok {
			t.Errorf("key is not ast.StringLiteral. got=%T", key)
		}

		expectedValue := expected[stringKey.String()]

		if intVal, ok := expectedValue.(int); ok {
			testIntegerLiteral(t, value, int64(intVal))
		} else {
			if value.String() != expectedValue {
				t.Errorf("value is not %s. got=%s", expectedValue, value.String())
			}
		}
	}
}

func TestEmptyHash(t *testing.T) {
	input := "{}"

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	hash, ok := statement.Expression.(*ast.Hash)

	if !ok {
		t.Fatalf("statement.Expression is not ast.Hash , got=%T", statement.Expression)
	}

	if len(hash.Pairs) != 0 {
		t.Errorf("hash.Pairs has wrong length. got=%d", len(hash.Pairs))
	}
}

func TestHashWithExpressions(t *testing.T) {
	input := `{"age1": 1 * 1, "age2": 1 / 1, "age3": 2 - 1}`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	hash, ok := statement.Expression.(*ast.Hash)

	if !ok {
		t.Fatalf("statement.Expression is not ast.Hash , got=%T", statement.Expression)
	}

	if len(hash.Pairs) != 3 {
		t.Errorf("hash.Pairs has wrong length. got=%d", len(hash.Pairs))
	}

	tests := map[string]func(ast.Expression){
		"age1": func(exp ast.Expression) {
			testInfixExpression(t, exp, 1, "*", 1)
		},
		"age2": func(exp ast.Expression) {
			testInfixExpression(t, exp, 1, "/", 1)
		},
		"age3": func(exp ast.Expression) {
			testInfixExpression(t, exp, 2, "-", 1)
		},
	}

	for key, value := range hash.Pairs {
		literal, ok := key.(*ast.StringLiteral)

		if !ok {
			t.Errorf("key is not ast.StringLiteral. got=%T", key)

			continue
		}

		testFunc, ok := tests[literal.String()]

		if !ok {
			t.Errorf("No test function for key %q found", literal.String())
			continue
		}

		testFunc(value)
	}
}

func TestHashKey(t *testing.T) {
	hello1 := &object.String{Value: "Hello World"}
	hello2 := &object.String{Value: "Hello World"}
	diff1 := &object.String{Value: "My name is DE"}
	diff2 := &object.String{Value: "My name is DE"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
