package parser

import (
	"token"
	"ast"
)


func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: p.currentToken}

	// if !p.expectPeekType(token.RETURN) {
		// return nil
	// }

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.currentTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}