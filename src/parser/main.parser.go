package parser

import (
	"ast"
	"lexer"
    "token"
	"fmt"
	"strconv"
)


func (p *Parser) parseIntegerLiteral() ast.Expression {
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
	expression := &ast.PrefixExpression{
		Token: p.currentToken,
		Operator: p.currentToken.Literal,
	}

	p.nextToken() // Move to the next token

	expression.Right = p.parseExpression(PREFIX)

	return expression
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
		fn prefixParseFunc
	}{
		{token.IDENT, p.parseIdentifier},
		{token.INT, p.parseIntegerLiteral},
		{token.EXCLAMATION, p.parsePrefixExpression},
		{token.MINUS, p.parsePrefixExpression},
	}

	for _, d := range data {
		p.registerPrefix(d.tokenType, d.fn)
	}
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexerInstance: l, errors: []string{}}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	p.prefixParseFuns = make(map[token.TokenType]prefixParseFunc)
	initRegisterPrefix(p)

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