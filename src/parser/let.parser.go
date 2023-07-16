package parser

import (
	"token"
	"ast"
)

func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.currentToken}

	if !p.expectPeekType(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	
	if !p.expectPeekType(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.currentTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}
	
	return statement
}