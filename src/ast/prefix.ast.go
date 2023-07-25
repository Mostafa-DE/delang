package ast

import (
	"bytes"
	"token"
)

type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

func (prefixExpression *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(prefixExpression.Operator)
	out.WriteString(prefixExpression.Right.String())
	out.WriteString(")")

	return out.String()
}

func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}
