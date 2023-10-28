package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
)

func TestParseSimpleArrayWithInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected []int64
	}{
		{"[1, 2, 3];", []int64{1, 2, 3}},
		{"[1, 2, 3, 4, 5];", []int64{1, 2, 3, 4, 5}},
		{"[1];", []int64{1}},
		{"[];", []int64{}},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := testExpressionStatement(t, program.Statements[0])
		array, ok := statement.Expression.(*ast.Array)

		if !ok {
			t.Fatalf("exp not *ast.Array. got=%T", statement.Expression)
		}

		if len(array.Elements) != len(val.expected) {
			t.Fatalf("len(array.Elements) not %d. got=%d", len(val.expected), len(array.Elements))
		}

		for idx, el := range array.Elements {
			testInteger(t, el, val.expected[idx])
		}
	}
}

func TestParseSimpleArrayWithStrings(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{`["Hello", "World"];`, []string{"Hello", "World"}},
		{`["Hello", "World", "Delang"];`, []string{"Hello", "World", "Delang"}},
		{`["Hello"];`, []string{"Hello"}},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := testExpressionStatement(t, program.Statements[0])
		array, ok := statement.Expression.(*ast.Array)
		if !ok {
			t.Fatalf("exp not *ast.Array. got=%T", statement.Expression)
		}

		if len(array.Elements) != len(val.expected) {
			t.Fatalf("len(array.Elements) not %d. got=%d", len(val.expected), len(array.Elements))
		}

		for idx, el := range array.Elements {
			testString(t, el, val.expected[idx])
		}
	}
}

func TestParseSimpleArrayWithBooleans(t *testing.T) {
	tests := []struct {
		input    string
		expected []bool
	}{
		{"[true, false];", []bool{true, false}},
		{"[false, true];", []bool{false, true}},
		{"[true, false, false, true];", []bool{true, false, false, true}},
		{"[false];", []bool{false}},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := testExpressionStatement(t, program.Statements[0])
		array, ok := statement.Expression.(*ast.Array)
		if !ok {
			t.Fatalf("exp not *ast.Array. got=%T", statement.Expression)
		}

		if len(array.Elements) != len(val.expected) {
			t.Fatalf("len(array.Elements) not %d. got=%d", len(val.expected), len(array.Elements))
		}

		for idx, el := range array.Elements {
			testBoolean(t, el, val.expected[idx])
		}
	}
}

func TestParseLengthNestedArray(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"[[1, 2, 3], [4, 5, 6]];", 2},
		{"[[1, 2, 3], [4, 5, 6], [7, 8, 9]];", 3},
		{"[[1, 2, 3]];", 1},
		{"[[1]];", 1},
		{"[[]];", 1},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := testExpressionStatement(t, program.Statements[0])
		array, ok := statement.Expression.(*ast.Array)

		if !ok {
			t.Fatalf("exp not *ast.Array. got=%T", statement.Expression)
		}

		if len(array.Elements) != val.expected {
			t.Fatalf("len(array.Elements) not %d. got=%d", val.expected, len(array.Elements))
		}
	}
}

func TestParseNestedArrayElements(t *testing.T) {
	tests := []struct {
		input    string
		expected [][]int64
	}{
		{"[[1, 2, 3], [4, 5, 6]];", [][]int64{{1, 2, 3}, {4, 5, 6}}},
		{"[[1, 2, 3], [4, 5, 6], [7, 8, 9]];", [][]int64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		{"[[1, 2, 3]];", [][]int64{{1, 2, 3}}},
		{"[[1]];", [][]int64{{1}}},
		{"[[]];", [][]int64{{}}},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := testExpressionStatement(t, program.Statements[0])
		array, ok := statement.Expression.(*ast.Array)

		if !ok {
			t.Fatalf("exp not *ast.Array. got=%T", statement.Expression)
		}

		if len(array.Elements) != len(val.expected) {
			t.Fatalf("len(array.Elements) not %d. got=%d", len(val.expected), len(array.Elements))
		}

		for idx, el := range array.Elements {
			nestedArray, ok := el.(*ast.Array)
			if !ok {
				t.Fatalf("nested exp not *ast.Array. got=%T", el)
			}

			if len(nestedArray.Elements) != len(val.expected[idx]) {
				t.Fatalf("len(nestedArray.Elements) not %d. got=%d", len(val.expected[idx]), len(nestedArray.Elements))
			}

			for jdx, nel := range nestedArray.Elements {
				testInteger(t, nel, val.expected[idx][jdx])
			}
		}
	}
}

func TestParseArrayWithInfixExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected []struct {
			left     int64
			operator string
			right    int64
		}
	}{
		{"[1 + 2, 3 - 4, 5 * 6];", []struct {
			left     int64
			operator string
			right    int64
		}{
			{1, "+", 2},
			{3, "-", 4},
			{5, "*", 6},
		}},
		{"[1 + 2];", []struct {
			left     int64
			operator string
			right    int64
		}{
			{1, "+", 2},
		}},
		{"[1 / 2];", []struct {
			left     int64
			operator string
			right    int64
		}{
			{1, "/", 2},
		}},
		{"[1 and 2];", []struct {
			left     int64
			operator string
			right    int64
		}{
			{1, "and", 2},
		}},
		{"[1 or 2];", []struct {
			left     int64
			operator string
			right    int64
		}{
			{1, "or", 2},
		}},
		{"[1 == 2];", []struct {
			left     int64
			operator string
			right    int64
		}{
			{1, "==", 2},
		}},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := testExpressionStatement(t, program.Statements[0])
		array, ok := statement.Expression.(*ast.Array)

		if !ok {
			t.Fatalf("exp not *ast.Array. got=%T", statement.Expression)
		}

		if len(array.Elements) != len(val.expected) {
			t.Fatalf("len(array.Elements) not %d. got=%d", len(val.expected), len(array.Elements))
		}

		for idx, el := range array.Elements {
			left := val.expected[idx].left
			operator := val.expected[idx].operator
			right := val.expected[idx].right
			testInfixExpression(t, el, left, operator, right)
		}
	}
}

func TestParseArrayWithPrefixExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected []struct {
			operator string
			right    int64
		}
	}{
		{"[!1, -2, !3, -4];", []struct {
			operator string
			right    int64
		}{
			{"!", 1},
			{"-", 2},
			{"!", 3},
			{"-", 4},
		}},
		{"[!1];", []struct {
			operator string
			right    int64
		}{
			{"!", 1},
		}},
		{"[-1];", []struct {
			operator string
			right    int64
		}{
			{"-", 1},
		}},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := testExpressionStatement(t, program.Statements[0])
		array, ok := statement.Expression.(*ast.Array)

		if !ok {
			t.Fatalf("exp not *ast.Array. got=%T", statement.Expression)
		}

		if len(array.Elements) != len(val.expected) {
			t.Fatalf("len(array.Elements) not %d. got=%d", len(val.expected), len(array.Elements))
		}

		for idx, el := range array.Elements {
			operator := val.expected[idx].operator
			right := val.expected[idx].right
			testPrefixExpression(t, el, operator, right)
		}
	}
}

func TestParseArrayWithIdentifier(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"[a, b, c];", []string{"a", "b", "c"}},
		{"[a];", []string{"a"}},
		{"[a, b];", []string{"a", "b"}},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := testExpressionStatement(t, program.Statements[0])
		array, ok := statement.Expression.(*ast.Array)

		if !ok {
			t.Fatalf("exp not *ast.Array. got=%T", statement.Expression)
		}

		if len(array.Elements) != len(val.expected) {
			t.Fatalf("len(array.Elements) not %d. got=%d", len(val.expected), len(array.Elements))
		}

		for idx, el := range array.Elements {
			testIdentifier(t, el, val.expected[idx])
		}
	}
}

func TestParseArrayWithFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected []struct {
			parameters []string
			body       []string
		}
	}{
		{`[fun(x, y) { return x + y }];`, []struct {
			parameters []string
			body       []string
		}{
			{[]string{"x", "y"}, []string{"return", "x", "+", "y"}},
		}},
		{`[fun(x, y) { return x + y }, fun(x, y) { return x - y }];`, []struct {
			parameters []string
			body       []string
		}{
			{[]string{"x", "y"}, []string{"return", "x", "+", "y"}},
			{[]string{"x", "y"}, []string{"return", "x", "-", "y"}},
		}},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		statement := testExpressionStatement(t, program.Statements[0])
		array, ok := statement.Expression.(*ast.Array)

		if !ok {
			t.Fatalf("exp not *ast.Array. got=%T", statement.Expression)
		}

		if len(array.Elements) != len(val.expected) {
			t.Fatalf("len(array.Elements) not %d. got=%d", len(val.expected), len(array.Elements))
		}

		for idx, el := range array.Elements {
			function, ok := el.(*ast.Function)
			params := val.expected[idx].parameters

			if !ok {
				t.Fatalf("exp not *ast.Function. got=%T", el)
			}

			if len(function.Parameters) != len(params) {
				t.Fatalf("len(function.Parameters) not %d. got=%d", len(params), len(function.Parameters))
			}

			for jdx, param := range function.Parameters {
				testIdentifier(t, param, params[jdx])
			}

			for _, stmt := range function.Body.Statements {
				returnStatement, ok := stmt.(*ast.ReturnStatement)
				if !ok {
					t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
					return
				}

				returnVal := returnStatement.ReturnValue.(*ast.InfixExpression)

				if returnVal == nil {
					t.Errorf("returnStatement.ReturnValue expected to be not nil")
					return
				}

				left := returnVal.Left.(*ast.Identifier)
				right := returnVal.Right.(*ast.Identifier)

				if left == nil || right == nil {
					t.Errorf("returnStatement.ReturnValue expected to be not nil")
					return
				}

				testInfixExpression(t, returnStatement.ReturnValue, left.Value, returnVal.Operator, right.Value)
			}
		}
	}
}

func TestParseSimpleIndexExpression(t *testing.T) {
	input := "arr[1 + 1];"

	program := parseProgram(t, input)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	indexExp, ok := stmt.Expression.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("exp not *ast.IndexExpression. got=%T", stmt.Expression)
	}

	testIdentifier(t, indexExp.Ident, "arr")
	testInfixExpression(t, indexExp.Index, 1, "+", 1)
}

func TestParseNestedIndexExpression(t *testing.T) {
	input := "arr[1 + 1][2 + 2];"

	program := parseProgram(t, input)
	stmt := testExpressionStatement(t, program.Statements[0])
	ident, ok := stmt.Expression.(*ast.IndexExpression).Ident.(*ast.IndexExpression)

	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}

	testIdentifier(t, ident.Ident, "arr")
	testInfixExpression(t, ident.Index, 1, "+", 1)

	indexExp, ok := stmt.Expression.(*ast.IndexExpression)

	if !ok {
		t.Fatalf("exp not *ast.IndexExpression. got=%T", stmt.Expression)
	}

	testInfixExpression(t, indexExp.Index, 2, "+", 2)
}
