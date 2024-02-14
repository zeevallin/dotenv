package lexer

import (
	"strings"

	"github.com/zeevallin/dotenv/token"
)

// Lexer represents a current lexing session of a dotenv file
type Lexer struct {
	input        []rune
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           rune // current char under examination
	line         int  // current line number
	column       int  // current column number
	tok          token.Token
}

// New spawns a new lexer using the given input
func New(input []byte) *Lexer {
	l := &Lexer{
		input: []rune(string(input)),
		tok:   token.Token{Type: token.EMPTY},
	}
	l.readChar()
	return l
}

// NextToken will attempt to extract the next token based on the lexer's position
func (l *Lexer) NextToken() token.Token {
	switch l.ch {
	case ' ', '\t', '\r':
		tok := l.newToken(token.WHITESPACE, string(l.ch))
		l.readChar()
		return tok
	case '\n':
		l.tok = l.newToken(token.NEWLINE, string(l.ch))
		l.readChar()
		return l.tok
	case '#':
		l.tok = l.newToken(token.COMMENT, "")
		l.tok.Literal = l.readComment()
		return l.tok
	case 0:
		l.tok = newToken(token.EOF, l.ch)
		l.readChar()
		return l.tok
	default:
		switch l.tok.Type {
		case token.EMPTY, token.NEWLINE:
			l.tok = l.newToken(token.KEY, "")
			l.tok.Literal = l.readKey()
			return l.tok
		case token.KEY:
			if isAssign(l.ch) {
				l.tok = l.newToken(token.ASSIGN, string(l.ch))
				l.readChar()
				return l.tok
			}
		case token.ASSIGN:
			l.tok = l.newToken(token.PLAIN_VALUE, "")
			l.tok.Literal = l.readValue()
			if strings.HasPrefix(l.tok.Literal, `"`) && strings.HasSuffix(l.tok.Literal, `"`) {
				l.tok.Type = token.QUOTED_VALUE
				return l.tok
			}
			return l.tok
		}

		l.tok = newToken(token.EOF, l.ch)
		l.readChar()
		return l.tok
	}
}

func (l *Lexer) newToken(tokenType token.Type, literal string) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: literal,
		Line:    l.line,
		Column:  l.column,
	}
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	if l.ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readKey() string {
	position := l.position
	for !(isWhiteSpace(l.ch) || isNewLine(l.ch) || isEOF(l.ch) || isAssign(l.ch) || isComment(l.ch)) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readValue() string {
	position := l.position
	if l.input[position] == '"' {
		for !(isNewLine(l.ch) || isEOF(l.ch) || isComment(l.ch)) {
			if l.peekChar() == '"' && l.ch != '\\' {
				l.readChar()
				l.readChar()
				break
			}
			l.readChar()
		}
	} else {
		for !(isNewLine(l.ch) || isEOF(l.ch) || isComment(l.ch)) {
			l.readChar()
		}
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readComment() string {
	position := l.position
	for !isNewLine(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func isAssign(ch rune) bool {
	return ch == '='
}

func isComment(ch rune) bool {
	return ch == '#'
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\r'
}

func isNewLine(ch rune) bool {
	return ch == '\n' || ch == ';'
}

func isEOF(ch rune) bool {
	return ch == '\x00'
}

func newToken(tokenType token.Type, ch rune) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
