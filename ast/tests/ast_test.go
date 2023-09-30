package tests

import (
	"testing"

	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/token"
)

func testString(t *testing.T, program ast.Node, expected string) {
	if program.String() != expected {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func TestLetString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "de1"},
					Value: "de1",
				},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "de2"},
					Value: "de2",
				},
			},
		},
	}

	testString(t, program, "let de1 = de2;")
}

func TestConstString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.CONST, Literal: "const"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "de1"},
					Value: "de1",
				},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "de2"},
					Value: "de2",
				},
			},
		},
	}

	testString(t, program, "const de1 = de2;")
}

func TestReturnString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Token: token.Token{Type: token.RETURN, Literal: "return"},
				ReturnValue: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "de1"},
					Value: "de1",
				},
			},
		},
	}

	testString(t, program, "return de1;")
}

func TestArrayString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Array{
					Token: token.Token{Type: token.LEFTBRAC, Literal: "["},
					Elements: []ast.Expression{
						&ast.Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "de1"},
							Value: "de1",
						},
						&ast.Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "de2"},
							Value: "de2",
						},
					},
				},
			},
		},
	}

	testString(t, program, "[de1, de2]")
}

func TestIndexExpressionString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.IndexExpression{
					Token: token.Token{Type: token.LEFTBRAC, Literal: "["},
					Ident: &ast.Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "de1"},
						Value: "de1",
					},
					Index: &ast.Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "de2"},
						Value: "de2",
					},
				},
			},
		},
	}

	testString(t, program, "(de1[de2])")
}

func TestHashString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Hash{
					Token: token.Token{Type: token.LEFTBRAC, Literal: "{"},
					Pairs: map[ast.Expression]ast.Expression{
						&ast.Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "de1"},
							Value: "de1",
						}: &ast.Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "de2"},
							Value: "de2",
						},
						&ast.Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "de3"},
							Value: "de3",
						}: &ast.Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "de4"},
							Value: "de4",
						},
					},
				},
			},
		},
	}

	testString(t, program, "{de1: de2, de3: de4}")
}

func TestFunctionString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Function{
					Token: token.Token{Type: token.FUNCTION, Literal: "fun"},
					Parameters: []*ast.Identifier{
						{
							Token: token.Token{Type: token.IDENT, Literal: "de1"},
							Value: "de1",
						},
						{
							Token: token.Token{Type: token.IDENT, Literal: "de2"},
							Value: "de2",
						},
					},
					Body: &ast.BlockStatement{
						Token: token.Token{Type: token.LEFTBRAC, Literal: "{"},
						Statements: []ast.Statement{
							&ast.ExpressionStatement{
								Expression: &ast.Identifier{
									Token: token.Token{Type: token.IDENT, Literal: "de3"},
									Value: "de3",
								},
							},
						},
					},
				},
			},
		},
	}

	testString(t, program, "fun(de1, de2) de3;")
}

func TestCallExpressionString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.CallFunction{
					Token: token.Token{Type: token.LEFTPAR, Literal: "("},
					Function: &ast.Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "de1"},
						Value: "de1",
					},
					Arguments: []ast.Expression{
						&ast.Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "de2"},
							Value: "de2",
						},
						&ast.Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "de3"},
							Value: "de3",
						},
					},
				},
			},
		},
	}

	testString(t, program, "de1(de2, de3)")
}

func TestIfStatement(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.IfExpression{
					Token: token.Token{Type: token.IF, Literal: "if"},
					Condition: &ast.InfixExpression{
						Token: token.Token{Type: token.INT, Literal: "1"},
						Left: &ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "1"},
							Value: 1,
						},
						Operator: ">",
						Right: &ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "2"},
							Value: 2,
						},
					},
					Consequence: &ast.BlockStatement{
						Token: token.Token{Type: token.LEFTBRAC, Literal: "{"},
						Statements: []ast.Statement{
							&ast.ExpressionStatement{
								Expression: &ast.Identifier{
									Token: token.Token{Type: token.IDENT, Literal: "de2"},
									Value: "de2",
								},
							},
						},
					},
					Alternative: &ast.BlockStatement{
						Token: token.Token{Type: token.LEFTBRAC, Literal: "{"},
						Statements: []ast.Statement{
							&ast.ExpressionStatement{
								Expression: &ast.Identifier{
									Token: token.Token{Type: token.IDENT, Literal: "de3"},
									Value: "de3",
								},
							},
						},
					},
				},
			},
		},
	}

	testString(t, program, "if (1 > 2): de2; else de3;")
}

func TestAssignStatement(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.AssignExpression{
					Token: token.Token{Type: token.ASSIGN, Literal: "="},
					Ident: &ast.Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "de1"},
						Value: "de1",
					},
					Value: &ast.Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "de2"},
						Value: "de2",
					},
				},
			},
		},
	}

	testString(t, program, "de1 = de2;")
}

func TestPrefixExpression(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.PrefixExpression{
					Token:    token.Token{Type: token.MINUS, Literal: "-"},
					Operator: "-",
					Right: &ast.IntegerLiteral{
						Token: token.Token{Type: token.INT, Literal: "1"},
						Value: 1,
					},
				},
			},
		},
	}

	testString(t, program, "(-1)")
}

func TestInfixExpression(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.InfixExpression{
					Token: token.Token{Type: token.INT, Literal: "1"},
					Left: &ast.IntegerLiteral{
						Token: token.Token{Type: token.INT, Literal: "1"},
						Value: 1,
					},
					Operator: ">",
					Right: &ast.IntegerLiteral{
						Token: token.Token{Type: token.INT, Literal: "2"},
						Value: 2,
					},
				},
			},
		},
	}

	testString(t, program, "(1 > 2)")
}

func TestBooleanExpression(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Boolean{
					Token: token.Token{Type: token.TRUE, Literal: "true"},
					Value: true,
				},
			},
		},
	}

	testString(t, program, "true")
}

func TestIntegerLiteralExpression(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.IntegerLiteral{
					Token: token.Token{Type: token.INT, Literal: "1"},
					Value: 1,
				},
			},
		},
	}

	testString(t, program, "1")
}

func TestStringLiteralExpression(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.StringLiteral{
					Token: token.Token{Type: token.STRING, Literal: ""},
					Value: "",
				},
			},
		},
	}

	testString(t, program, "")
}

func TestIdentifierExpression(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "de1"},
					Value: "de1",
				},
			},
		},
	}

	testString(t, program, "de1")
}

func TestExpressionStatement(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "de1"},
					Value: "de1",
				},
			},
		},
	}

	testString(t, program, "de1")
}

func TestDuringStatement(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.DuringExpression{
					Token: token.Token{Type: token.DURING, Literal: "during"},
					Condition: &ast.InfixExpression{
						Token: token.Token{Type: token.INT, Literal: "1"},
						Left: &ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "1"},
							Value: 1,
						},
						Operator: ">",
						Right: &ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "2"},
							Value: 2,
						},
					},
					Body: &ast.BlockStatement{
						Token: token.Token{Type: token.LEFTBRAC, Literal: "{"},
						Statements: []ast.Statement{
							&ast.ExpressionStatement{
								Expression: &ast.Identifier{
									Token: token.Token{Type: token.IDENT, Literal: "de2"},
									Value: "de2",
								},
							},
						},
					},
				},
			},
		},
	}

	testString(t, program, "during (1 > 2): de2;")
}

func TestBreakStatement(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.BreakStatement{
				Token: token.Token{Type: token.BREAK, Literal: "break"},
			},
		},
	}

	testString(t, program, "break")
}

func TestSkipStatement(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.SkipStatement{
				Token: token.Token{Type: token.SKIP, Literal: "skip"},
			},
		},
	}

	testString(t, program, "skip")
}

// func TestForStatement(t *testing.T) {
// 	program := &ast.Program{
// 		Statements: []ast.Statement{
// 			&ast.ExpressionStatement{
// 				Expression: &ast.ForExpression{
// 					Token: token.Token{Type: token.FOR, Literal: "for"},
// 					Condition: &ast.InfixExpression{
// 						Token: token.Token{Type: token.INT, Literal: "1"},
// 						Left: &ast.IntegerLiteral{
// 							Token: token.Token{Type: token.INT, Literal: "1"},
// 							Value: 1,
// 						},
// 						Operator: ">",
// 						Right: &ast.IntegerLiteral{
// 							Token: token.Token{Type: token.INT, Literal: "2"},
// 							Value: 2,
// 						},
// 					},
// 					Body: &ast.BlockStatement{
// 						Token: token.Token{Type: token.LEFTBRAC, Literal: "{"},
// 						Statements: []ast.Statement{
// 							&ast.ExpressionStatement{
// 								Expression: &ast.Identifier{
// 									Token: token.Token{Type: token.IDENT, Literal: "de2"},
// 									Value: "de2",
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	testString(t, program, "for (1 > 2): de2;")
// }
