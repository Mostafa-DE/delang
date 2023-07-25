package ast

import (
	"bytes"
	"token"
)

type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (letStatement *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(letStatement.TokenLiteral() + " ")
	out.WriteString(letStatement.Name.Value)
	out.WriteString(" = ")

	if letStatement.Value != nil {
		out.WriteString(letStatement.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
