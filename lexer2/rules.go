package lexer2

func lex(l *Lexer) stateFn {
	for {
		if l.next() == eof {
			break
		}
		switch {
		case isSpace(l.char):
			return lexSpace
		case isNewLine(l.char):
			return lexNewLine
		case l.char == '"' || l.char == '\'':
			return lexString
		case l.char == '=':
			l.emit(EQUALS)
			return lex
		case l.char == '#':
			return lexComment
		default:
			return lexText
		}
	}
	l.emit(EOF)
	return nil
}

func lexSpace(l *Lexer) stateFn {
	l.acceptRunRunes(' ', '\t')
	l.emit(SPACE)
	return lex
}

func lexNewLine(l *Lexer) stateFn {
	l.emit(NEWLINE)
	return lex
}

func lexIdentifier(l *Lexer) stateFn {
	for {
		switch l.next() {
		case ' ', '\t', '\n', '\r', '=', eof:
			l.backup()
			l.emit(IDENTIFIER)
			return lex
		}
	}
}

func lexText(l *Lexer) stateFn {
	switch {
	case isDigit(l.char) || l.char == '=':
	}
	for {
		ch := l.next()
		switch {
		case !isChar(ch):
		case isNewLine(ch) || ch == '=' || ch == '#' || ch == '"' || ch == '\'' || ch == eof:
			l.backup()
			l.emit(TEXT)
			return lex
		}
	}
}

func lexComment(l *Lexer) stateFn {
	for {
		switch l.next() {
		case '\n', '\r', eof:
			l.backup()
			l.emit(COMMENT)
			return lex
		}
	}
}

func lexString(l *Lexer) stateFn {
	quote := l.char
	line, col := l.line, l.col
	for {
		switch l.next() {
		case quote:
			l.emit(STRING)
			return lex
		case '\n', '\r', eof:
			l.backup()
			l.errorf("unclosed string at %s:%d:%d", l.name, line, col)
			l.emit(ILLEGAL)
			return lex
		}
	}
}

func isNewLine(char rune) bool {
	return char == '\n' || char == '\r'
}

func isSpace(char rune) bool {
	return char == ' ' || char == '\t'
}

func isChar(char rune) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z'
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
