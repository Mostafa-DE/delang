package ast

import "token"


type Node interface {
	TokenLiteral() string
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

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

type LetStatement struct {
	Token token.Token // token.LET
	Name *Identifier
	Value Expression
}

func (p *Program) TokenLiteral() string { // used only for debugging and testing
	if len(p.Statements) > 0{
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (ls *LetStatement) statementNode(){}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
