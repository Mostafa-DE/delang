package ast

import (
	"bytes"

	"github.com/Mostafa-DE/delang/token"
)

type AssignExpression struct {
	Token token.Token // The token.ASSIGN token
	Ident *Identifier
	Value Expression
}

func (assignExpression *AssignExpression) expressionNode() {}
func (assignExpression *AssignExpression) TokenLiteral() string {
	return assignExpression.Token.Literal
}

func (assignExpression *AssignExpression) statementNode() {}

func (assignExpression *AssignExpression) String() string {
	var out bytes.Buffer

	out.WriteString(assignExpression.Ident.String())
	out.WriteString(" = ")
	if assignExpression.Value != nil {
		out.WriteString(assignExpression.Value.String())
		out.WriteString(";")
	}

	return out.String()
}
