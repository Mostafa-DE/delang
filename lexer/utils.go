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
	return char >= '0' && char <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.currentPosition
	for isLetter(l.currentChar) || isNumber(l.currentChar) {
		l.readChar()
	}

	return l.input[position:l.currentPosition]
}

func (l *Lexer) readNumber() string {
	// In DE we only support `int` and `float` types
	// For alternative float you can use `decimal` type
	// Refer to the documentation for more information
	// https://delang.mostafade.com/play/dataTypes/decimal1

	position := l.currentPosition
	for isNumber(l.currentChar) {
		l.readChar()
	}

	for l.currentChar == '.' || isNumber(l.currentChar) {
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

func (l *Lexer) peekNChar(NumCharToLook int64) byte {
	if l.currentPosition+int(NumCharToLook) >= len(l.input) {
		return 0
	} else {
		return l.input[l.currentPosition+int(NumCharToLook)]
	}
}

func (l *Lexer) readString() string {
	// TODO: add support for character escaping
	position := l.currentPosition + 1
	for {
		l.readChar()
		if l.currentChar == '\n' || l.currentChar == '\r' || l.currentChar == 0 || l.currentChar == '"' {
			break
		}
	}

	return l.input[position:l.currentPosition]
}

func (l *Lexer) skipComment() {
	for l.currentChar != '\n' && l.currentChar != 0 {
		l.readChar()
	}
}
