package ast

import (
	"bytes"

	"github.com/Mostafa-DE/delang/token"
)

type DuringExpression struct {
	Token     token.Token // token.DURING
	Condition Expression
	Body      *BlockStatement // during body
}

func (duringExpression *DuringExpression) expressionNode() {}
func (duringExpression *DuringExpression) TokenLiteral() string {
	return duringExpression.Token.Literal
}

func (duringExpression *DuringExpression) String() string {
	var out bytes.Buffer

	out.WriteString("during")
	out.WriteString(" ")
	out.WriteString(duringExpression.Condition.String() + ":")
	out.WriteString(" ")
	out.WriteString(duringExpression.Body.String() + ";")

	return out.String()
}
