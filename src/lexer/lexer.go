package lexer

import "token"

type Lexer struct {
	input string
	currentPosition int // current position in input, points to the character in the input that corresponds to the ch byte.
	readNextPosition int // current reading position in input, points to the “next” character in the input.
	currentChar byte // current char under examination.
}

func New(input string) *Lexer {
	l := &Lexer{input: input} // position, readPosition = 0 by default because the type (int)
	l.readChar()
	return l
}


func (l *Lexer) readChar() {
	if l.readNextPosition >= len(l.input) {
		l.currentChar = 0 // 0 in ASCII means NUL wich indicate that we reach the end of file
	} else {
		l.currentChar = l.input[l.readNextPosition]
	}

	l.currentPosition = l.readNextPosition
	l.readNextPosition += 1 // Always point to Next index
}


func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.currentChar {
		case '=':
			// You may consider extract this to function
			if l.peekChar() == '=' {
				prevChar := l.currentChar
				l.readChar() // to increment the position
				tok = token.Token{Type: token.EQUAL, Literal: string(prevChar) + string(l.currentChar)}
			} else {
				tok = newToken(token.ASSIGN, l.currentChar)
			}
		case '+':
			tok = newToken(token.PLUS, l.currentChar)
		case '-':
			tok = newToken(token.MINUS, l.currentChar)
		case '{':
			tok = newToken(token.LEFTBRAC, l.currentChar)
		case '}':
			tok = newToken(token.RIGHTBRAC, l.currentChar)
		case '(':
			tok = newToken(token.LEFTPAR, l.currentChar)
		case ')':
			tok = newToken(token.RIGHTPAR, l.currentChar)
		case '!':
			// You may consider extract this to function
			if l.peekChar() == '=' {
				prevChar := l.currentChar // "!"
				l.readChar()
				tok = token.Token{Type: token.NOTEQUAL, Literal: string(prevChar) + string(l.currentChar)}
			} else {
				tok = newToken(token.EXCLAMATION, l.currentChar)
			}
		case '/':
			tok = newToken(token.SLASH, l.currentChar)
		case '*':
			tok = newToken(token.ASTERISK, l.currentChar)
		case '<':
			tok = newToken(token.LESSTHAN, l.currentChar)
		case '>':
			tok = newToken(token.GREATERTHAN, l.currentChar)
		case ';':
			tok = newToken(token.SEMICOLON, l.currentChar)
		case ',':
			tok = newToken(token.COMMA, l.currentChar)
		case 0: // End of the line
			tok.Literal = ""
			tok.Type = token.EOFILE
		default:
			if isLetter(l.currentChar) {
				tok.Literal = l.readIdentifier()
				tok.Type = token.LookupIdent(tok.Literal)
				return tok // This return is important because we don't want to call readChar() again
			} else if isNumber(l.currentChar) {
				tok.Literal = l.readNumber()
				tok.Type = token.INT
				return tok // This return is important because we don't want to call readChar() again
			} else {
				tok = newToken(token.ILLEGAL, l.currentChar)	
			}
	}

	l.readChar()
	return tok
}

func isLetter(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.currentPosition
	for isLetter(l.currentChar) {
		l.readChar()
	}

	return l.input[position:l.currentPosition]
}


func newToken(tokenType token.TokenType, currentChar byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(currentChar)}
}

func (l *Lexer) skipWhiteSpace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.readChar()
	}
}

func isNumber(char byte) bool { 
	// This is just for INT, consider adding (float, hex, octal) etc...
	return char >= '0' && char <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.currentPosition
	for isNumber(l.currentChar) {
		l.readChar()
	}

	return l.input[position: l.currentPosition]
}

func (l *Lexer) peekChar() byte { 
	// This function similar to readChar but we don't increment the position or set currentChar, for something like == and !=
	// The idea of this function just to look to the next char that come after = or ! to decide is it from type
	// ASSIGN('=') or EQUAL('==') or BANG('!') or NOTEQUAL('!=')
	if l.readNextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readNextPosition]
	}

}
