package parser

import (
	"fmt"
	"strconv"

	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/token"
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

	p.nextToken()

	statement.Value = p.parseExpression(LOWEST)

	if p.peekTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	statement.ReturnValue = p.parseExpression(LOWEST)

	if p.peekTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	// defer untrace(trace("parseExpressionStatement"))
	statement := &ast.ExpressionStatement{Token: p.currentToken}
	statement.Expression = p.parseExpression(LOWEST)

	if p.peekTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	// defer untrace(trace("parsePrefixExpression"))
	expression := &ast.PrefixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
	}

	p.nextToken() // Move to the next token

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	// defer untrace(trace("parseInfixExpression"))
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

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	// defer untrace(trace("parseIntegerLiteral"))
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

func (p *Parser) parseBoolean() ast.Expression {
	// defer untrace(trace("parseBoolean"))
	return &ast.Boolean{Token: p.currentToken, Value: p.currentTokenTypeIs(token.TRUE)}
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	// defer untrace(trace("parseGroupedExpression"))
	p.nextToken()

	expression := p.parseExpression(LOWEST)

	if !p.expectPeekType(token.RIGHTPAR) {
		return nil
	}

	return expression
}

func (p *Parser) parseIfExpression() ast.Expression {
	// defer untrace(trace("parseIfExpression"))
	expression := &ast.IfExpression{Token: p.currentToken}

	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)

	if !p.expectPeekType(token.COLON) {
		p.errors = append(p.errors, "Expected ':' after if condition")
		return nil
	}

	if !p.expectPeekType(token.LEFTBRAC) {
		p.errors = append(p.errors, "Expected '{' after if condition")
		return nil
	}

	expression.Consequence = p.parseBlockStatement()

	if p.peekTokenTypeIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeekType(token.LEFTBRAC) {
			p.errors = append(p.errors, "Expected '{' after else")
			return nil
		}

		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	// defer untrace(trace("parseBlockStatement"))
	block := &ast.BlockStatement{Token: p.currentToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.currentTokenTypeIs(token.RIGHTBRAC) && !p.currentTokenTypeIs(token.EOFILE) {
		if p.currentTokenTypeIs(token.ELSE) {
			p.errors = append(p.errors, "Unexpected 'else' statement, if block is not closed with '}'")
			return nil
		}

		statement := p.parseStatement()
		block.Statements = append(block.Statements, statement)

		p.nextToken()
	}

	return block
}

func (p *Parser) parseFunction() ast.Expression {
	// defer untrace(trace("parseFunction"))
	function := &ast.Function{Token: p.currentToken}

	if !p.expectPeekType(token.LEFTPAR) {
		return nil
	}

	function.Parameters = p.parseFunctionParameters()

	if !p.expectPeekType(token.LEFTBRAC) {
		return nil
	}

	function.Body = p.parseBlockStatement()

	return function
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	// defer untrace(trace("parseFunctionParameters"))
	identifiers := []*ast.Identifier{}

	if p.peekTokenTypeIs(token.RIGHTPAR) {
		// No parameters
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	identifier := &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	identifiers = append(identifiers, identifier)

	for p.peekTokenTypeIs(token.COMMA) {
		// Skip the comma
		p.nextToken()
		p.nextToken()

		// Parse the identifier
		identifier := &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
		identifiers = append(identifiers, identifier)
	}

	if !p.expectPeekType(token.RIGHTPAR) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseCallFunction(function ast.Expression) ast.Expression {
	// defer untrace(trace("parseCallExpression"))
	expression := &ast.CallFunction{Token: p.currentToken, Function: function}
	expression.Arguments = p.parseCallArguments()

	return expression
}

func (p *Parser) parseCallArguments() []ast.Expression {
	// defer untrace(trace("parseCallArguments"))
	arguments := []ast.Expression{}

	if p.peekTokenTypeIs(token.RIGHTPAR) {
		// No arguments
		p.nextToken()
		return arguments
	}

	p.nextToken()
	arguments = append(arguments, p.parseExpression(LOWEST))

	for p.peekTokenTypeIs(token.COMMA) {
		// Skip the comma
		p.nextToken()
		p.nextToken()

		// Parse the argument
		arguments = append(arguments, p.parseExpression(LOWEST))
	}

	if !p.expectPeekType(token.RIGHTPAR) {
		return nil
	}

	return arguments
}

func (p *Parser) parseStringLiteral() ast.Expression {
	// defer untrace(trace("parseStringLiteral"))
	return &ast.StringLiteral{Token: p.currentToken, Value: p.currentToken.Literal}
}
