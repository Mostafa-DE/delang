package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/token"
)

// TODO: Make this test more generic
func TestLetString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "name"},
					Value: "name",
				},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherName"},
					Value: "anotherName",
				},
			},
		},
	}

	if program.String() != "let name = anotherName;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func TestConstString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.CONST, Literal: "const"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "name"},
					Value: "name",
				},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherName"},
					Value: "anotherName",
				},
			},
		},
	}

	if program.String() != "const name = anotherName;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
