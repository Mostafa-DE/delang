package parser

import (
	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/token"

	"github.com/Mostafa-DE/delang/lexer"
)

const (
	_            int = iota
	LOWEST           // Lowest precedence
	AND_OR           // and or or
	EQUAL            // ==
	LESS_GREATER     // > or <
	SUM_SUB          // + or -
	MUL_DIV_MOD      // * or / or %
	PREFIX           // -X or !X
	CALL             // myFunction(X)
	INDEX            // array[index]
)

var precedences = map[token.TokenType]int{
	token.EQUAL:       EQUAL,
	token.NOTEQUAL:    EQUAL,
	token.LESSTHAN:    LESS_GREATER,
	token.GREATERTHAN: LESS_GREATER,
	token.AND:         AND_OR,
	token.OR:          AND_OR,
	token.PLUS:        SUM_SUB,
	token.MINUS:       SUM_SUB,
	token.SLASH:       MUL_DIV_MOD,
	token.ASTERISK:    MUL_DIV_MOD,
	token.MOD:         MUL_DIV_MOD,
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
