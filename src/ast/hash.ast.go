package ast

import (
	"bytes"
	"strings"

	"github.com/Mostafa-DE/delang/token"
)

type Hash struct {
	Token token.Token // The '{' token
	Pairs map[Expression]Expression
}

func (hash *Hash) expressionNode() {}
func (hash *Hash) TokenLiteral() string {
	return hash.Token.Literal
}

func (hash *Hash) String() string {
	var out bytes.Buffer

	pairs := []string{}

	for key, value := range hash.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
