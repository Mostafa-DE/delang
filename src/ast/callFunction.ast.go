package ast

import (
	"bytes"
	"strings"
	"token"
)

type CallFunction struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (callFunction *CallFunction) expressionNode()      {}
func (callFunction *CallFunction) TokenLiteral() string { return callFunction.Token.Literal }

func (callFunction *CallFunction) String() string {
	var out bytes.Buffer

	args := []string{}

	for _, arg := range callFunction.Arguments {
		args = append(args, arg.String())
	}

	out.WriteString(callFunction.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
