package parser

import (
	"ast"
	"lexer"
    "token"
	"fmt"
)

type Parser struct {
	lexerInstance *lexer.Lexer
	currentToken token.Token
	peekToken token.Token
	errors []string
}


func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexerInstance.NextToken()
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexerInstance: l, errors: []string{}}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParserProgram() *ast.Program{
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.currentTokenTypeIs(token.EOFILE) {
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
		case token.RETURN:
			return p.parseReturnStatement()
		default:
			return nil
	}
}


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

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: p.currentToken}


	if !p.expectPeekType(token.RETURN) {
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.currentTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) currentTokenTypeIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenTypeIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeekType(t token.TokenType) bool {
	if p.peekTokenTypeIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(tokType token.TokenType) {
	msg := fmt.Sprintf("Expected next token to be '%s', got '%s' instead", tokType, p.peekToken.Type)

	p.errors = append(p.errors, msg)
}