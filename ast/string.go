package ast

import "github.com/Mostafa-DE/delang/token"

type StringLiteral struct {
	Token token.Token // token.STRING
	Value string
}

func (stringLiteral *StringLiteral) String() string {
	return stringLiteral.TokenLiteral()
}

func (stringLiteral *StringLiteral) expressionNode() {}
func (stringLiteral *StringLiteral) TokenLiteral() string {
	return stringLiteral.Token.Literal
}
