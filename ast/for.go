package ast

import (
	"bytes"

	"github.com/Mostafa-DE/delang/token"
)

type ForStatement struct {
	Token      token.Token     // The token.FOR
	IdxIdent   *Identifier     // The index identifier of the loop
	VarIdent   *Identifier     // The variable identifier of the loop
	Expression Expression      // The expression to iterate over
	Body       *BlockStatement // The body of the loop
}

func (fs *ForStatement) expressionNode() {}

func (fs *ForStatement) TokenLiteral() string {
	return fs.Token.Literal
}

func (fs *ForStatement) String() string {
	var out bytes.Buffer

	out.WriteString("for ")
	out.WriteString(fs.IdxIdent.String())
	out.WriteString(", ")
	out.WriteString(fs.VarIdent.String())
	out.WriteString(" in ")
	out.WriteString(fs.Expression.String())
	out.WriteString(":")
	out.WriteString(" {")
	out.WriteString(fs.Body.String() + ";")
	out.WriteString("}")

	return out.String()

}

func (fs *ForStatement) statementNode() {}
