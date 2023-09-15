package ast

import (
	"bytes"

	"github.com/Mostafa-DE/delang/token"
)

type ConstStatement struct {
	Token token.Token // token.CONST
	Name  *Identifier
	Value Expression
}

func (constStatement *ConstStatement) String() string {
	var out bytes.Buffer

	out.WriteString(constStatement.TokenLiteral() + " ")
	out.WriteString(constStatement.Name.Value)
	out.WriteString(" = ")

	if constStatement.Value != nil {
		out.WriteString(constStatement.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (cs *ConstStatement) statementNode() {}
func (cs *ConstStatement) TokenLiteral() string {
	return cs.Token.Literal
}
