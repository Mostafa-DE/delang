package parser

import (
	"ast"
	"fmt"
	"lexer"
	"strconv"
	"token"
)

func (p *Parser) parseGroupedExpression() ast.Expression {
	defer untrace(trace("parseGroupedExpression"))
	p.nextToken()

	expression := p.parseExpression(LOWEST)

	if !p.expectPeekType(token.RIGHTPAR) {
		return nil
	}

	return expression
}

func (p *Parser) parseBoolean() ast.Expression {
	defer untrace(trace("parseBoolean"))
	return &ast.Boolean{Token: p.currentToken, Value: p.currentTokenTypeIs(token.TRUE)}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	defer untrace(trace("parseIntegerLiteral"))
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

func (p *Parser) parsePrefixExpression() ast.Expression {
	defer untrace(trace("parsePrefixExpression"))
	expression := &ast.PrefixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
	}

	p.nextToken() // Move to the next token

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	defer untrace(trace("parseInfixExpression"))
	expression := &ast.InfixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
		Left:     left,
	}

	precedence := p.currentPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	defer untrace(trace("parseExpression"))
	prefixParseFunc := p.prefixParseFuns[p.currentToken.Type]
	if prefixParseFunc == nil {
		p.noPrefixParseFnError(p.currentToken.Type)
		return nil
	}

	leftExp := prefixParseFunc()

	for !p.peekTokenTypeIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infixParseFunc := p.infixParseFuns[p.peekToken.Type]
		if infixParseFunc == nil {
			return leftExp
		}

		p.nextToken() // Move to the next token (operator) because we already have the left expression

		leftExp = infixParseFunc(leftExp)
	}

	return leftExp
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	defer untrace(trace("parseExpressionStatement"))
	statement := &ast.ExpressionStatement{Token: p.currentToken}
	statement.Expression = p.parseExpression(LOWEST)

	if p.peekTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexerInstance.NextToken()
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFunc) {
	p.infixParseFuns[tokenType] = fn
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFunc) {
	p.prefixParseFuns[tokenType] = fn
}

func initRegisterPrefix(p *Parser) {
	data := []struct {
		tokenType token.TokenType
		fn        prefixParseFunc
	}{
		{token.IDENT, p.parseIdentifier},
		{token.INT, p.parseIntegerLiteral},
		{token.EXCLAMATION, p.parsePrefixExpression},
		{token.MINUS, p.parsePrefixExpression},
		{token.TRUE, p.parseBoolean},
		{token.FALSE, p.parseBoolean},
		{token.LEFTPAR, p.parseGroupedExpression},
		{token.RIGHTPAR, p.parseGroupedExpression},
	}

	for _, val := range data {
		p.registerPrefix(val.tokenType, val.fn)
	}
}

func initRegisterInfix(p *Parser) {
	data := []struct {
		tokenType token.TokenType
		fn        infixParseFunc
	}{
		{token.PLUS, p.parseInfixExpression},
		{token.MINUS, p.parseInfixExpression},
		{token.SLASH, p.parseInfixExpression},
		{token.ASTERISK, p.parseInfixExpression},
		{token.EQUAL, p.parseInfixExpression},
		{token.NOTEQUAL, p.parseInfixExpression},
		{token.LESSTHAN, p.parseInfixExpression},
		{token.GREATERTHAN, p.parseInfixExpression},
	}

	for _, val := range data {
		p.registerInfix(val.tokenType, val.fn)
	}
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexerInstance: l, errors: []string{}}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	p.prefixParseFuns = make(map[token.TokenType]prefixParseFunc)
	initRegisterPrefix(p)

	p.infixParseFuns = make(map[token.TokenType]infixParseFunc)
	initRegisterInfix(p)

	return p
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

func (p *Parser) ParserProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.currentTokenTypeIs(token.EOFILE) {
		statement := p.parseStatement()

		program.Statements = append(program.Statements, statement)

		p.nextToken()
	}

	return program
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("No prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
