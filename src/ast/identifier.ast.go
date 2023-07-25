package ast

import (
	"token"
)

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (identifier *Identifier) String() string {
	return identifier.Value
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
