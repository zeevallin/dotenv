package token

import "strings"

type Type string

const (
	EMPTY        Type = "EMPTY"
	COMMENT      Type = "COMMENT"
	WHITESPACE   Type = "WHITESPACE"
	ASSIGN       Type = "ASSIGN"
	KEY          Type = "KEY"
	PLAIN_VALUE  Type = "PLAIN_VALUE"
	QUOTED_VALUE Type = "QUOTED_VALUE"
	NEWLINE      Type = "NEWLINE"
	ILLEGAL      Type = "ILLEGAL"
	EOF          Type = "EOF"
)

type Token struct {
	Type    Type
	Literal string
	Line    int
	Column  int
}

func IsQuotedValue(value string) bool {
	return strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`)
}

func New(line, col int, t Type, literal string) Token {
	return Token{
		Type:    t,
		Literal: literal,
		Line:    line,
		Column:  col,
	}
}
