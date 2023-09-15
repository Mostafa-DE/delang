package ast

import "github.com/Mostafa-DE/delang/token"

type ExpressionStatement struct {
	// We need this because sometimes we have an expression act like a statement
	// e.g let x = 5;
	// x + 10; // this is an expression but it acts like a statement
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (expressionStatement *ExpressionStatement) String() string {
	if expressionStatement.Expression != nil {
		return expressionStatement.Expression.String()
	}

	return ""
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
