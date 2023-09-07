package parser

import (
	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/lexer"
	"github.com/Mostafa-DE/delang/token"
)

func (p *Parser) parseExpression(precedence int) ast.Expression {
	// defer untrace(trace("parseExpression"))
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
		{token.IF, p.parseIfExpression},
		{token.FUNCTION, p.parseFunction},
		{token.STRING, p.parseStringLiteral},
		{token.LEFTSQPRAC, p.parseArray},
		{token.LEFTBRAC, p.parseHash},
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
		{token.LEFTPAR, p.parseCallFunction},
		{token.LEFTSQPRAC, p.parseIndexExpression},
	}

	for _, val := range data {
		p.registerInfix(val.tokenType, val.fn)
	}
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexerInstance: l, errors: []string{}}

	// Read two tokens, so currentToken and peekToken are both set
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

	case token.CONST:
		return p.parseConstStatement()

	case token.RETURN:
		return p.parseReturnStatement()

	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.currentTokenTypeIs(token.EOFILE) {
		statement := p.parseStatement()

		program.Statements = append(program.Statements, statement)

		p.nextToken()
	}

	return program
}
