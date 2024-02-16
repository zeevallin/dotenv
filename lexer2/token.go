package lexer2

type Type string

const (
	TEXT       Type = "TEXT"
	COMMENT    Type = "COMMENT"
	IDENTIFIER Type = "IDENTIFIER"
	STRING     Type = "STRING"
	SPACE      Type = "SPACE"
	NEWLINE    Type = "NEWLINE"
	EQUALS     Type = "EQUALS"
	ERROR      Type = "ERROR"
	ILLEGAL    Type = "ILLEGAL"
	EOF        Type = "EOF"
)

type Token struct {
	Kind    Type
	Literal string
	Line    int
	Column  int
}
