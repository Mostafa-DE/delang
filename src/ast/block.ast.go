package ast

import (
	"bytes"

	"github.com/Mostafa-DE/delang/token"
)

type BlockStatement struct {
	Token      token.Token // token.LEFTBRAC and token.RIGHTBRAC
	Statements []Statement
}

func (blockStatement *BlockStatement) statementNode() {}
func (blockStatement *BlockStatement) TokenLiteral() string {
	return blockStatement.Token.Literal
}

func (blockStatement *BlockStatement) String() string {
	var out bytes.Buffer

	for _, statement := range blockStatement.Statements {
		out.WriteString(statement.String())
	}

	return out.String()
}
