package lexer

func isLetter(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_'
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

func (l *Lexer) readIdentifier() string {
	position := l.currentPosition
	for isLetter(l.currentChar) {
		l.readChar()
	}

	return l.input[position:l.currentPosition]
}

func (l *Lexer) readNumber() string {
	position := l.currentPosition
	for isNumber(l.currentChar) {
		l.readChar()
	}

	return l.input[position:l.currentPosition]
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

func (l *Lexer) peekChar() byte {
	/*
		This function similar to readChar but we don't increment the position or set currentChar,
		for something like == and != , The idea of this function just to look to the next char that
		come after = or ! to decide is it from type ASSIGN('=') or EQUAL('==') or BANG('!') or NOTEQUAL('!=')
	*/
	if l.readNextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readNextPosition]
	}

}
