package ast

import "github.com/Mostafa-DE/delang/token"

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
