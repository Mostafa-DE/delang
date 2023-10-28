package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/parser"
)

func TestItShouldWorkWithLetStatement(t *testing.T) {
	tests := []struct {
		input      string
		idxIdent   string
		varIdent   string
		arrayIdent string
	}{
		{
			`
				let arr = [1, 2, 3, 4, 5];
				for idx, num in arr: {
					logs(num);
				}
			`,
			"idx",
			"num",
			"arr",
		},
		{
			`
				let nums = [1, 2];
				for _, num in nums: {
					logs(num);
				}
			`,
			"",
			"num",
			"nums",
		},
		{
			`
				let nums = [1, 2, 3];
				for num in nums: {
					logs(num);
				}
			`,
			"",
			"num",
			"nums",
		},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)

		if !testForLoopStatement(t, program.Statements[1], val.idxIdent, val.varIdent, val.arrayIdent) {
			return
		}

		forStmt, ok := program.Statements[1].(*ast.ForStatement)

		if !ok {
			t.Fatalf("stmt not *ast.ForStatement. got=%T", program.Statements[1])
		}

		bodyStmt, ok := forStmt.Body.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("stmt not *ast.ExpressionStatement. got=%T", forStmt.Body.Statements[0])
		}

		if bodyStmt.Expression.String() != "logs(num)" {
			t.Fatalf("bodyStmt.Expression.String() not 'logs(num)', got=%q", bodyStmt.Expression.String())
		}
	}
}

func TestItShouldWorkWithConsts(t *testing.T) {
	tests := []struct {
		input      string
		idxIdent   string
		varIdent   string
		arrayIdent string
	}{
		{
			`
				const arr = [1, 2, 3, 4, 5];
				for idx, num in arr: {
					logs(num);
				}
			`,
			"idx",
			"num",
			"arr",
		},
		{
			`
				const nums = [1, 2];
				for _, num in nums: {
					logs(num);
				}
			`,
			"",
			"num",
			"nums",
		},
		{
			`
				const nums = [1, 2, 3];
				for num in nums: {
					logs(num);
				}
			`,
			"",
			"num",
			"nums",
		},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)

		if !testForLoopStatement(t, program.Statements[1], val.idxIdent, val.varIdent, val.arrayIdent) {
			return
		}

		forStmt, ok := program.Statements[1].(*ast.ForStatement)

		if !ok {
			t.Fatalf("stmt not *ast.ForStatement. got=%T", program.Statements[1])
		}

		bodyStmt, ok := forStmt.Body.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("stmt not *ast.ExpressionStatement. got=%T", forStmt.Body.Statements[0])
		}

		if bodyStmt.Expression.String() != "logs(num)" {
			t.Fatalf("bodyStmt.Expression.String() not 'logs(num)', got=%q", bodyStmt.Expression.String())
		}
	}
}

func TestItShouldWorkWithExpression(t *testing.T) {
	tests := []struct {
		input      string
		idxIdent   string
		varIdent   string
		arrayIdent string
	}{
		{
			`
				for idx, num in [1, 2, 3, 4, 5]: {
					logs(num);
				}
			`,
			"idx",
			"num",
			"",
		},
		{
			`
				for _, num in [1, 2]: {
					logs(num);
				}
			`,
			"",
			"num",
			"",
		},
		{
			`
				for num in [1, 2, 3]: {
					logs(num);
				}
			`,
			"",
			"num",
			"",
		},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)

		if !testForLoopStatement(t, program.Statements[0], val.idxIdent, val.varIdent, val.arrayIdent) {
			return
		}

		forStmt, ok := program.Statements[0].(*ast.ForStatement)

		if !ok {
			t.Fatalf("stmt not *ast.ForStatement. got=%T", program.Statements[0])
		}

		bodyStmt, ok := forStmt.Body.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("stmt not *ast.ExpressionStatement. got=%T", forStmt.Body.Statements[0])
		}

		if bodyStmt.Expression.String() != "logs(num)" {
			t.Fatalf("bodyStmt.Expression.String() not 'logs(num)', got=%q", bodyStmt.Expression.String())
		}
	}
}

func TestItShouldWorkWithBreak(t *testing.T) {
	tests := []struct {
		input     string
		excpected string
	}{
		{
			`
				for idx, num in [1, 2, 3, 4, 5]: {
					if num == 3: {
						break;
					}
					logs(num);
				}
			`,
			"if (num == 3): {break;}",
		},
		{
			`
				for _, num in [1, 2]: {
					if num == 2: {
						break;
					}
					logs(num);
				}
			`,
			"if (num == 2): {break;}",
		},
		{
			`
				for num in [1, 2, 3]: {
					if num == 1: {
						break;
					}
					logs(num);
				}
			`,
			"if (num == 1): {break;}",
		},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		forStmt, ok := program.Statements[0].(*ast.ForStatement)

		if !ok {
			t.Fatalf("stmt not *ast.ForStatement. got=%T", program.Statements[0])
		}

		bodyStmt, ok := forStmt.Body.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("stmt not *ast.ExpressionStatement. got=%T", forStmt.Body.Statements[0])
		}

		if bodyStmt.Expression.String() != val.excpected {
			t.Fatalf(
				"bodyStmt.Expression.String() not '%s', got=%q",
				val.excpected,
				bodyStmt.Expression.String(),
			)
		}
	}
}

func TestItShouldWorkWithSkip(t *testing.T) {
	tests := []struct {
		input     string
		excpected string
	}{
		{
			`
				for idx, num in [1, 2, 3, 4, 5]: {
					if num == 3: {
						skip;
					}
					logs(num);
				}
			`,
			"if (num == 3): {skip;}",
		},
		{
			`
				for _, num in [1, 2]: {
					if num == 2: {
						skip;
					}
					logs(num);
				}
			`,
			"if (num == 2): {skip;}",
		},
		{
			`
				for num in [1, 2, 3]: {
					if num == 1: {
						skip;
					}
					logs(num);
				}
			`,
			"if (num == 1): {skip;}",
		},
	}

	for _, val := range tests {
		program := parseProgram(t, val.input)
		forStmt, ok := program.Statements[0].(*ast.ForStatement)

		if !ok {
			t.Fatalf("stmt not *ast.ForStatement. got=%T", program.Statements[0])
		}

		bodyStmt, ok := forStmt.Body.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("stmt not *ast.ExpressionStatement. got=%T", forStmt.Body.Statements[0])
		}

		if bodyStmt.Expression.String() != val.excpected {
			t.Fatalf(
				"bodyStmt.Expression.String() not '%s', got=%q",
				val.excpected,
				bodyStmt.Expression.String(),
			)
		}
	}
}

func TestItShouldFailWithIncorrectSyntax(t *testing.T) {
	tests := []struct {
		input         string
		expectedError string
	}{
		{
			`
				for idx, num in [1, 2] {
					logs(num);
				}
			`,
			"Expected next token to be ':', got '{' instead",
		},
		{
			`
				for idx, num in [1]:
					logs(num);
			`,
			"Expected next token to be '{', got 'IDENT' instead",
		},
		{
			`
				for idx, num [1, 2]: {
					logs(num);
				}

			`,
			"Expected next token to be 'IN', got '[' instead",
		},
		{
			`
				for _ in [1, 2]: {
					logs(num);
				}
			`,
			"Expected a comma after underscore",
		},
		{
			`
				for _, in [1, 2]: {
					logs(num);
				}
			`,
			"Expected an identifier after underscore",
		},
		{
			`
				for _, _ in [1, 2]: {
					logs(num);
				}
			`,
			"Cannot use two underscores in for statement",
		},
		{
			`
				for idx, _ in [1, 2]: {
					logs(num);
				}
			`,
			"Cannot use underscore as a variable identifier in for statement",
		},
	}

	for _, val := range tests {
		l := lexer.New(val.input)
		p := parser.New(l)

		p.ParseProgram()

		errors := p.Errors()

		if len(errors) == 0 {
			t.Errorf("Parser didn't return any errors")

			return
		}

		if errors[0] != val.expectedError {
			t.Errorf("Error message not '%s', got=%s", val.expectedError, errors[0])
		}

	}
}

func testForLoopStatement(t *testing.T, s ast.Statement, idxIdent string, varIdent string, expression string) bool {
	forStmt, ok := s.(*ast.ForStatement)

	if !ok {
		t.Errorf("s not *ast.ForStatement. got=%T", s)
		return false
	}

	if idxIdent != "" {
		if forStmt.IdxIdent.Value != idxIdent {
			t.Errorf("forStmt.IdxIdent.Value not '%s'. got=%s", idxIdent, forStmt.IdxIdent.Value)
			return false
		}
	}

	if forStmt.VarIdent.Value != varIdent {
		t.Errorf("forStmt.VarIdent.Value not '%s'. got=%s", varIdent, forStmt.VarIdent.Value)
		return false
	}

	if expression != "" {

		if forStmt.Expression.String() != expression {
			t.Errorf("forStmt.Expression.String() not '%s'. got=%s", expression, forStmt.Expression.String())
			return false
		}
	}

	return true
}
