package lexer2

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var eof = rune(0)

type stateFn func(*Lexer) stateFn

func New(name, input string) *Lexer {
	l := &Lexer{
		name:   name,
		input:  input,
		start:  0,
		pos:    0,
		width:  0,
		char:   0,
		line:   1,
		col:    0,
		state:  lex,
		tokens: make(chan Token, 2),
	}
	return l
}

type Lexer struct {
	name  string
	input string

	start int // start position of this token
	pos   int // current position in the input
	width int // width of last rune read from input

	char rune // last rune read from input

	col  int
	line int

	state stateFn

	tokens chan Token
}

// NextToken returns the next token from the input
func (l *Lexer) NextToken() Token {
	for {
		select {
		case token := <-l.tokens:
			return token
		default:
			l.state = l.state(l)
		}
	}
}

// emit passes an item back to the client.
func (l *Lexer) emit(t Type) {
	l.tokens <- Token{
		Kind:    t,
		Literal: l.input[l.start:l.pos],
		Line:    l.line,
		Column:  l.col,
	}
	l.start = l.pos
}

// next returns the next rune in the input.
func (l *Lexer) next() (rune rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	l.char, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width

	if l.char == '\n' {
		l.line++
		l.col = 0
	} else {
		l.col++
	}

	return l.char
}

// ignore skips over the pending input before this point.
func (l *Lexer) ignore() {
	l.start = l.pos
}

// backup steps back one rune.
// can be called only once per call of next.
func (l *Lexer) backup() {
	l.pos -= l.width
	l.col--
}

// peek returns but does not consume the next rune in the input.
func (l *Lexer) peek() rune {
	rune := l.next()
	l.backup()
	return rune
}

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, l.next()) {
	}
	l.backup()
}

func (l *Lexer) acceptRunRunes(runes ...rune) {
	for containsRune(runes, l.next()) {
	}
	l.backup()
}

func containsRune(runes []rune, r rune) bool {
	for _, rr := range runes {
		if rr == r {
			return true
		}
	}
	return false
}

func (l *Lexer) errorf(format string, args ...any) {
	l.tokens <- Token{
		Kind:    ERROR,
		Literal: fmt.Sprintf(format, args...),
		Line:    l.line,
		Column:  l.col,
	}
}
