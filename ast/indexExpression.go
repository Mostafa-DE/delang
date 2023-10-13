package ast

import (
	"bytes"

	"github.com/Mostafa-DE/delang/token"
)

type IndexExpression struct {
	Token token.Token // The [ token
	Ident Expression
	Index Expression
	Value Expression
}

func (idx *IndexExpression) expressionNode() {}
func (idx *IndexExpression) TokenLiteral() string {
	return idx.Token.Literal
}

func (idx *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(idx.Ident.String())
	out.WriteString("[")
	out.WriteString(idx.Index.String())
	out.WriteString("])")

	return out.String()
}
