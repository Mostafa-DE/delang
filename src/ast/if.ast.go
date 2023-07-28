package ast

import (
	"bytes"
	"token"
)

type IfExpression struct {
	token.Token // token.IF
	Condition   Expression
	Consequence *BlockStatement // If block
	Alternative *BlockStatement // (Optional) Else block
}

func (ifExpression *IfExpression) expressionNode() {}
func (ifExpression *IfExpression) TokenLiteral() string {
	return ifExpression.Token.Literal
}
func (ifExpression *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(" ")
	out.WriteString(ifExpression.Condition.String())
	out.WriteString(" ")
	out.WriteString(ifExpression.Consequence.String())

	if ifExpression.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(ifExpression.Alternative.String())
	}

	return out.String()
}
