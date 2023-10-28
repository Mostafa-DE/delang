package ast

import (
	"bytes"

	"github.com/Mostafa-DE/delang/token"
)

type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression
}

func (returnStatement *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(returnStatement.TokenLiteral() + " ")

	if returnStatement.ReturnValue != nil {
		out.WriteString(returnStatement.ReturnValue.String())
	}

	return out.String()
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
