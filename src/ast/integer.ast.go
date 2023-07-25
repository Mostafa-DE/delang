package ast

import (
	"token"
)

type IntegerLiteral struct {
	Token token.Token // token.INT
	Value int64
}

func (integerLiteral *IntegerLiteral) String() string {
	return integerLiteral.TokenLiteral()
}

func (integerLiteral *IntegerLiteral) expressionNode() {}
func (integerLiteral *IntegerLiteral) TokenLiteral() string {
	return integerLiteral.Token.Literal
}
