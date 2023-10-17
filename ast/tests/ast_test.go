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

func TestLetStatement(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "name"},
					Value: "name",
				},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "'DE!'"},
					Value: "'DE!'",
				},
			},
		},
	}

	testString(t, program, "let name = 'DE!';")
}

func TestConstStatement(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ConstStatement{
				Token: token.Token{Type: token.CONST, Literal: "const"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "PI"},
					Value: "PI",
				},
				Value: &ast.StringLiteral{
					Token: token.Token{Type: token.IDENT, Literal: "'3.14'"},
					Value: "'3.14'",
				},
			},
		},
	}

	testString(t, program, "const PI = '3.14';")
}

func TestReturnStatement(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Token: token.Token{Type: token.RETURN, Literal: "return"},
				ReturnValue: &ast.Integer{
					Token: token.Token{Type: token.IDENT, Literal: "100"},
					Value: 100,
				},
			},
		},
	}

	testString(t, program, "return 100;")
}

func TestArray(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Array{
					Token: token.Token{Type: token.LEFTBRAC, Literal: "["},
					Elements: []ast.Expression{
						&ast.Integer{
							Token: token.Token{Type: token.IDENT, Literal: "1"},
							Value: 1,
						},
						&ast.Integer{
							Token: token.Token{Type: token.IDENT, Literal: "2"},
							Value: 2,
						},
						&ast.Integer{
							Token: token.Token{Type: token.IDENT, Literal: "3"},
							Value: 3,
						},
					},
				},
			},
		},
	}

	testString(t, program, "[1, 2, 3]")
}

func TestArrayIndex(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.IndexExpression{
					Token: token.Token{Type: token.LEFTBRAC, Literal: "["},
					Ident: &ast.Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "arr"},
						Value: "arr",
					},
					Index: &ast.Integer{
						Token: token.Token{Type: token.IDENT, Literal: "0"},
						Value: 0,
					},
				},
			},
		},
	}

	testString(t, program, "(arr[0])")
}

func TestHashIndex(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.IndexExpression{
					Token: token.Token{Type: token.LEFTBRAC, Literal: "["},
					Ident: &ast.Hash{
						Token: token.Token{Type: token.LEFTBRAC, Literal: "{"},
						Pairs: map[ast.Expression]ast.Expression{
							&ast.Identifier{
								Token: token.Token{Type: token.IDENT, Literal: "name"},
								Value: "name",
							}: &ast.Identifier{
								Token: token.Token{Type: token.IDENT, Literal: "DE!"},
								Value: "DE!",
							},
							&ast.Identifier{
								Token: token.Token{Type: token.IDENT, Literal: "age"},
								Value: "age",
							}: &ast.Identifier{
								Token: token.Token{Type: token.IDENT, Literal: "10"},
								Value: "10",
							},
						},
					},
					Index: &ast.StringLiteral{
						Token: token.Token{Type: token.STRING, Literal: "'name'"},
						Value: "'name'",
					},
				},
			},
		},
	}

	testString(t, program, "({name: DE!, age: 10}['name'])")
}

func TestHash(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Hash{
					Token: token.Token{Type: token.LEFTBRAC, Literal: "{"},
					Pairs: map[ast.Expression]ast.Expression{
						&ast.Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "name"},
							Value: "name",
						}: &ast.Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "DE!"},
							Value: "DE!",
						},
						&ast.Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "age"},
							Value: "age",
						}: &ast.Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "10"},
							Value: "10",
						},
					},
				},
			},
		},
	}

	testString(t, program, "{name: DE!, age: 10}")
}

func TestFunction(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Function{
					Token: token.Token{Type: token.FUNCTION, Literal: "fun"},
					Parameters: []*ast.Identifier{
						{
							Token: token.Token{Type: token.IDENT, Literal: "param1"},
							Value: "param1",
						},
						{
							Token: token.Token{Type: token.IDENT, Literal: "param2"},
							Value: "param2",
						},
					},
					Body: &ast.BlockStatement{
						Token: token.Token{Type: token.LEFTBRAC, Literal: "{"},
						Statements: []ast.Statement{
							&ast.ExpressionStatement{
								Expression: &ast.Identifier{
									Token: token.Token{Type: token.IDENT, Literal: "logs(param1, param2)"},
									Value: "logs(param1, param2)",
								},
							},
						},
					},
				},
			},
		},
	}

	testString(t, program, "fun(param1, param2) {logs(param1, param2);}")
}

func TestFuncCall(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.CallFunction{
					Token: token.Token{Type: token.LEFTPAR, Literal: "("},
					Function: &ast.Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "func"},
						Value: "func",
					},
					Arguments: []ast.Expression{
						&ast.Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "arg1"},
							Value: "arg1",
						},
						&ast.Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "arg2"},
							Value: "arg2",
						},
					},
				},
			},
		},
	}

	testString(t, program, "func(arg1, arg2)")
}

func TestIfStatement(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.IfExpression{
					Token: token.Token{Type: token.IF, Literal: "if"},
					Condition: &ast.InfixExpression{
						Token: token.Token{Type: token.INT, Literal: "1"},
						Left: &ast.Integer{
							Token: token.Token{Type: token.INT, Literal: "1"},
							Value: 1,
						},
						Operator: ">",
						Right: &ast.Integer{
							Token: token.Token{Type: token.INT, Literal: "2"},
							Value: 2,
						},
					},
					Consequence: &ast.BlockStatement{
						Token: token.Token{Type: token.LEFTBRAC, Literal: "{"},
						Statements: []ast.Statement{
							&ast.ExpressionStatement{
								Expression: &ast.Identifier{
									Token: token.Token{Type: token.IDENT, Literal: "logs('Awesome!')"},
									Value: "logs('Awesome!')",
								},
							},
						},
					},
					Alternative: &ast.BlockStatement{
						Token: token.Token{Type: token.LEFTBRAC, Literal: "{"},
						Statements: []ast.Statement{
							&ast.ExpressionStatement{
								Expression: &ast.Identifier{
									Token: token.Token{Type: token.IDENT, Literal: "logs('Not Awesome!')"},
									Value: "logs('Not Awesome!')",
								},
							},
						},
					},
				},
			},
		},
	}

	testString(t, program, "if (1 > 2): {logs('Awesome!');} else {logs('Not Awesome!');}")
}

func TestAssignStatement(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.AssignExpression{
					Token: token.Token{Type: token.ASSIGN, Literal: "="},
					Ident: &ast.Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "var1"},
						Value: "var1",
					},
					Value: &ast.Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "var2"},
						Value: "var2",
					},
				},
			},
		},
	}

	testString(t, program, "var1 = var2;")
}

func TestPrefixExpression(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.PrefixExpression{
					Token:    token.Token{Type: token.MINUS, Literal: "-"},
					Operator: "-",
					Right: &ast.Integer{
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
					Left: &ast.Integer{
						Token: token.Token{Type: token.INT, Literal: "1"},
						Value: 1,
					},
					Operator: ">",
					Right: &ast.Integer{
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

func TestIntegerExpression(t *testing.T) {
	// The expression is anything that returns a value
	// 1 as an expression returns 1, also true, false, etc.

	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Integer{
					Token: token.Token{Type: token.INT, Literal: "1"},
					Value: 1,
				},
			},
		},
	}

	testString(t, program, "1")
}

func TestFloatExpression(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Float{
					Token: token.Token{Type: token.FLOAT, Literal: "1.1"},
					Value: 1.1,
				},
			},
		},
	}

	testString(t, program, "1.1")
}

func TestStringLiteralExpression(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.StringLiteral{
					Token: token.Token{Type: token.STRING, Literal: "This is a string"},
					Value: "This is a string",
				},
			},
		},
	}

	testString(t, program, "This is a string")
}

func TestIdentifierExpression(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "var"},
					Value: "var",
				},
			},
		},
	}

	testString(t, program, "var")
}

func TestDuringStatement(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.DuringExpression{
					Token: token.Token{Type: token.DURING, Literal: "during"},
					Condition: &ast.InfixExpression{
						Token: token.Token{Type: token.INT, Literal: "1"},
						Left: &ast.Integer{
							Token: token.Token{Type: token.INT, Literal: "1"},
							Value: 1,
						},
						Operator: ">",
						Right: &ast.Integer{
							Token: token.Token{Type: token.INT, Literal: "2"},
							Value: 2,
						},
					},
					Body: &ast.BlockStatement{
						Token: token.Token{Type: token.LEFTBRAC, Literal: "{"},
						Statements: []ast.Statement{
							&ast.ExpressionStatement{
								Expression: &ast.Identifier{
									Token: token.Token{Type: token.IDENT, Literal: "logs('Awesome!')"},
									Value: "logs('Awesome!')",
								},
							},
						},
					},
				},
			},
		},
	}

	testString(t, program, "during (1 > 2): {logs('Awesome!');}")
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

func TestForStatement(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ForStatement{
				Token: token.Token{Type: token.FOR, Literal: "for"},
				IdxIdent: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "idx"},
					Value: "idx",
				},
				VarIdent: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "num"},
					Value: "num",
				},
				Expression: &ast.Array{
					Token: token.Token{Type: token.LEFTBRAC, Literal: "["},
					Elements: []ast.Expression{
						&ast.Integer{
							Token: token.Token{Type: token.INT, Literal: "1"},
							Value: 1,
						},
						&ast.Integer{
							Token: token.Token{Type: token.INT, Literal: "2"},
							Value: 2,
						},
						&ast.Integer{
							Token: token.Token{Type: token.INT, Literal: "3"},
							Value: 3,
						},
					},
				},
				Body: &ast.BlockStatement{
					Token: token.Token{Type: token.LEFTBRAC, Literal: "{"},
					Statements: []ast.Statement{
						&ast.ExpressionStatement{
							Expression: &ast.Identifier{
								Token: token.Token{Type: token.IDENT, Literal: "logs(num)"},
								Value: "logs(num)",
							},
						},
					},
				},
			},
		},
	}

	testString(t, program, "for idx, num in [1, 2, 3]: {logs(num);}")
}
