package parser

import (
	"ast"
	"lexer"
	"token"
)

type Parser struct {
	lexerInstance *lexer.Lexer
	currentToken token.Token
	peekToken token.Token
}


func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexerInstance.NextToken()
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexerInstance: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParserProgram() *ast.Program{
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != token.EOFILE {
		statement := p.parseStatement()
		
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		p.nextToken()
	}

	return program
}


func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
		case token.LET:
			return p.parseLetStatement()
		default:
			return nil
	}
}


func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.currentToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	
	return statement
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
	
		return false
	}
}
