package parser

import (
	"token"
	"lexer"
	"ast"
)


const (
	_ int = iota // iota is a special constant that can be used to increment values
	LOWEST
	EQUALS // ==
	LESSGREATER // > or <
	SUM // +
	PRODUCT //. *
	PREFIX // -X or !X
	CALL // myFunction(X)
)

type (
	prefixParseFunc func() ast.Expression
	infixParseFunc func(ast.Expression) ast.Expression
)

type Parser struct {// TODO: This is should be in ast_types.go file
	lexerInstance *lexer.Lexer
	currentToken token.Token
	peekToken token.Token
	errors []string
	prefixParseFuns map[token.TokenType]prefixParseFunc
	infixParseFuns map[token.TokenType]infixParseFunc
}