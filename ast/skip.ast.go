package ast

import "github.com/Mostafa-DE/delang/token"

type SkipStatement struct {
	Token token.Token // the 'skip' token
}

func (ls *SkipStatement) statementNode() {}

func (ls *SkipStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *SkipStatement) String() string {
	return ls.Token.Literal
}
