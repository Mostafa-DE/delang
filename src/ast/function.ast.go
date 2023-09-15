package ast

import (
	"bytes"
	"strings"

	"github.com/Mostafa-DE/delang/token"
)

type Function struct {
	Token      token.Token // the 'fun' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (f *Function) expressionNode()      {}
func (f *Function) TokenLiteral() string { return f.Token.Literal }

func (f *Function) String() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(f.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(f.Body.String() + ";")

	return out.String()
}
