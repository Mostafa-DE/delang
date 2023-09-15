package parser

import (
	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/token"

	"github.com/Mostafa-DE/delang/lexer"
)

const (
	_ int = iota // iota is a special constant that can be used to increment values
	LOWEST
	EQUAL       // ==
	LESSGREATER // > or <
	SUMSUB      // + or -
	MULDIV      //. *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
	INDEX       // array[index]
)

var precedences = map[token.TokenType]int{
	token.EQUAL:       EQUAL,
	token.NOTEQUAL:    EQUAL,
	token.LESSTHAN:    LESSGREATER,
	token.GREATERTHAN: LESSGREATER,
	token.PLUS:        SUMSUB,
	token.MINUS:       SUMSUB,
	token.SLASH:       MULDIV,
	token.ASTERISK:    MULDIV,
	token.LEFTPAR:     CALL,
	token.LEFTSQPRAC:  INDEX, // array indexing has the highest precedence
}

type (
	prefixParseFunc func() ast.Expression
	infixParseFunc  func(ast.Expression) ast.Expression
)

type Parser struct {
	lexerInstance   *lexer.Lexer
	currentToken    token.Token
	peekToken       token.Token
	errors          []string
	prefixParseFuns map[token.TokenType]prefixParseFunc
	infixParseFuns  map[token.TokenType]infixParseFunc
}
