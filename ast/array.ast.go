package ast

import (
	"bytes"
	"strings"

	"github.com/Mostafa-DE/delang/token"
)

type Array struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

type IndexExpression struct {
	Token token.Token // The [ token
	Ident Expression
	Index Expression
}

func (array *Array) expressionNode() {}
func (array *Array) TokenLiteral() string {
	return array.Token.Literal
}

func (array *Array) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range array.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
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
