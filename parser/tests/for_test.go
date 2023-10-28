package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
)

func TestForStatement(t *testing.T) {
	inputs := []string{
		`
			let arr = [1, 2, 3, 4, 5];
			for idx, num in arr: {
				logs(num);
			}
		`,
		`
			for idx, num in [1, 2, 3, 4, 5]: {
				logs(num);
			}
		`,
	}

	for idx, input := range inputs {
		program := parseProgram(t, input)

		if program == nil {
			t.Fatalf("ParseProgram() returned nil")
		}

		if idx == 0 {
			if len(program.Statements) != 2 {
				t.Fatalf("program.Statements does not contain 2 statements. got=%d", len(program.Statements))
			}

			if !testLet(t, program.Statements[0], "arr") {
				return
			}

			if !testForStatement(t, program.Statements[1], "idx", "num", "arr") {
				return
			}

			forStmt, ok := program.Statements[1].(*ast.ForStatement)

			if !ok {
				t.Fatalf("stmt not *ast.ForStatement. got=%T", program.Statements[1])
			}

			if len(forStmt.Body.Statements) != 1 {
				t.Fatalf("forStmt.Body.Statements does not contain 1 statement. got=%d", len(forStmt.Body.Statements))
			}

			bodyStmt, ok := forStmt.Body.Statements[0].(*ast.ExpressionStatement)

			if !ok {
				t.Fatalf("stmt not *ast.ExpressionStatement. got=%T", forStmt.Body.Statements[0])
			}

			if bodyStmt.Expression.String() != "logs(num)" {
				t.Fatalf("bodyStmt.Expression.String() not 'logs(num)', got=%q", bodyStmt.Expression.String())
			}

		} else {
			if len(program.Statements) != 1 {
				t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
			}

			if !testForStatement(t, program.Statements[0], "idx", "num", "[1, 2, 3, 4, 5]") {
				return
			}

			forStmt, ok := program.Statements[0].(*ast.ForStatement)

			if !ok {
				t.Fatalf("stmt not *ast.ForStatement. got=%T", program.Statements[0])
			}

			if len(forStmt.Body.Statements) != 1 {
				t.Fatalf("forStmt.Body.Statements does not contain 1 statement. got=%d", len(forStmt.Body.Statements))
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
}

func testLet(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true
}

func testForStatement(t *testing.T, s ast.Statement, idxIdent string, varIdent string, expression string) bool {
	forStmt, ok := s.(*ast.ForStatement)

	if !ok {
		t.Errorf("s not *ast.ForStatement. got=%T", s)
		return false
	}

	if forStmt.IdxIdent.Value != idxIdent {
		t.Errorf("forStmt.IdxIdent.Value not '%s'. got=%s", idxIdent, forStmt.IdxIdent.Value)
		return false
	}

	if forStmt.VarIdent.Value != varIdent {
		t.Errorf("forStmt.VarIdent.Value not '%s'. got=%s", varIdent, forStmt.VarIdent.Value)
		return false
	}

	if forStmt.Expression.String() != expression {
		t.Errorf("forStmt.Expression.String() not '%s'. got=%s", expression, forStmt.Expression.String())
		return false
	}

	return true
}
