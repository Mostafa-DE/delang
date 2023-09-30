package ast

import "github.com/Mostafa-DE/delang/token"

type BreakStatement struct {
	Token token.Token // token.BREAK
}

func (breakStatement *BreakStatement) statementNode()       {}
func (breakStatement *BreakStatement) TokenLiteral() string { return breakStatement.Token.Literal }

func (breakStatement *BreakStatement) String() string {
	return breakStatement.Token.Literal
}
