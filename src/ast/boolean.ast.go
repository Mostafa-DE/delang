package ast

import (
	"token"
)

type Boolean struct {
	Token token.Token
	Value bool
}

func (boolean *Boolean) String() string {
	return boolean.TokenLiteral()
}

func (b *Boolean) expressionNode() {}
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}
