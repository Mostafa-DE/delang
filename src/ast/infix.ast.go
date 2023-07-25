package ast

import (
	"bytes"
	"token"
)

type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (infixExpression *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(infixExpression.Left.String())
	out.WriteString(" " + infixExpression.Operator + " ")
	out.WriteString(infixExpression.Right.String())
	out.WriteString(")") // e.g => (1 + 2)

	return out.String()
}

func (infixExpression *InfixExpression) expressionNode() {}
func (infixExpression *InfixExpression) TokenLiteral() string {
	return infixExpression.Token.Literal
}
