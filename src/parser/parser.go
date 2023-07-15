package parser

import (
	"ast"
	"lexer"
    "token"
	"fmt"
	"strconv"
)

type Parser struct {// TODO: This is should be in ast_types.go file
	lexerInstance *lexer.Lexer
	currentToken token.Token
	peekToken token.Token
	errors []string
	prefixParseFuns map[token.TokenType]prefixParseFunc
	infixParseFuns map[token.TokenType]infixParseFunc
}

type (// TODO: This is should be in ast_types.go file
	prefixParseFunc func() ast.Expression
	infixParseFunc func(ast.Expression) ast.Expression
)

const (
	_ int = iota // iota is a special constant that can be used to increment values
	LOWEST
	EQUALS // ==
	LESSGREATER // > or <
	SUM // +
	PRODUCT //. *
	PREFIX // -X or !X
	CALL // myFunction(X)
)


func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFunc) {
	p.prefixParseFuns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFunc) {
	p.infixParseFuns[tokenType] = fn
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	// This function to convert the current token to an integer literal
	literal := &ast.IntegerLiteral{Token: p.currentToken}

	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("Could not parse %q as integer", p.currentToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	literal.Value = value
	return literal
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
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

	p.prefixParseFuns = make(map[token.TokenType]prefixParseFunc)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)

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
			return p.parseExpressionStatement()
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseLetStatement() *ast.LetStatement {// TODO: This is should be in sperate file
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

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {// TODO: This is should be in sperate file
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

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {// TODO: This is should be in sperate file
	statement := &ast.ExpressionStatement{Token: p.currentToken}
	statement.Expression = p.parseExpression(LOWEST)

	if p.peekTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseExpression(precedence int) ast.Expression {// TODO: This is should be in sperate file
	prefix := p.prefixParseFuns[p.currentToken.Type]
	if prefix == nil {
		return nil
	}

	leftExp := prefix()

	return leftExp
}




func (p *Parser) currentTokenTypeIs(t token.TokenType) bool {// TODO: This is should be in utils file
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenTypeIs(t token.TokenType) bool {// TODO: This is should be in utils file
	return p.peekToken.Type == t
}

func (p *Parser) expectPeekType(t token.TokenType) bool {// TODO: This is should be in utils file
	if p.peekTokenTypeIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(tokType token.TokenType) {// TODO: This is should be in utils file
	msg := fmt.Sprintf("Expected next token to be '%s', got '%s' instead", tokType, p.peekToken.Type)

	p.errors = append(p.errors, msg)
}
