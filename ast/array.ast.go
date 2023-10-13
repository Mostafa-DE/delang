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
