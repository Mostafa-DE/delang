package ast

import (
	"bytes"

	"github.com/Mostafa-DE/delang/token"
)

type Node interface {
	TokenLiteral() string
	String() string // This will allow us to print AST nodes for debugging and to compare them with other AST nodes.
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct { // Root Node
	Statements []Statement
}

type VariableStatement struct {
	Token token.Token // token.LET or token.CONST
	Type  string      // "let" or "const"
	Name  *Identifier
	Value Expression
}

func (p *Program) TokenLiteral() string { // used only for debugging and testing
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// This will allow us to print AST nodes for debugging and to compare them with other AST nodes.
func (p *Program) String() string { // used only for debugging and testing
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
