package ast

import "github.com/Mostafa-DE/delang/token"

type Integer struct {
	Token token.Token // token.INT
	Value int64
}

type Float struct {
	Token token.Token // token.Float
	Value float64
}

func (integer *Integer) String() string {
	return integer.TokenLiteral()
}

func (integer *Integer) expressionNode() {}
func (integer *Integer) TokenLiteral() string {
	return integer.Token.Literal
}

func (decimal *Float) String() string {
	return decimal.TokenLiteral()
}

func (decimal *Float) expressionNode() {}
func (decimal *Float) TokenLiteral() string {
	return decimal.Token.Literal
}
