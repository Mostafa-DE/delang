package lexer

import (
	"strings"

	"github.com/Mostafa-DE/delang/token"
)

type Lexer struct {
	input            string
	currentPosition  int  // current position in input, points to the character in the input that corresponds to the ch byte.
	readNextPosition int  // current reading position in input, points to the “next” character in the input.
	currentChar      byte // current char under examination.
}

func New(input string) *Lexer {
	// Replace the single quotes with double quotes
	// This is because single quotes in Go used to represent runes (characters) not strings
	// But in our language we want to allow single double quotes to represent strings
	input = strings.ReplaceAll(input, `'`, `"`)

	l := &Lexer{input: input}
	l.readChar()
	return l
}

func newToken(tokenType token.TokenType, currentChar byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(currentChar)}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.currentChar {
	case '=':
		if l.peekChar() == '=' {
			prevChar := l.currentChar // "="
			l.readChar()
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

	case ':':
		tok = newToken(token.COLON, l.currentChar)

	case '!':
		if l.peekChar() == '=' {
			prevChar := l.currentChar // "!"
			l.readChar()
			tok = token.Token{Type: token.NOTEQUAL, Literal: string(prevChar) + string(l.currentChar)}
		} else {
			tok = newToken(token.EXCLAMATION, l.currentChar)
		}

	case '/':
		if l.peekChar() == '/' {
			l.skipComment()
			return l.NextToken()
		}
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

	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()

	case '[':
		tok = newToken(token.LEFTSQPRAC, l.currentChar)

	case ']':
		tok = newToken(token.RIGHTSQPRAC, l.currentChar)

	case '%':
		tok = newToken(token.MOD, l.currentChar)

	case '_':
		tok = newToken(token.UNDERSCORE, l.currentChar)

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
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.currentChar)
		}
	}

	l.readChar()
	return tok
}
