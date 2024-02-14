package lexer_test

import (
	"testing"

	"github.com/zeevallin/dotenv/lexer"
	"github.com/zeevallin/dotenv/token"
)

func TestNextToken(t *testing.T) {
	input := ` # Comment 1 
FOO_ONE=BAR_ONE # Comment 2
FOO_TWO = BAR_TWO
FOO_THREE= BAR_THREE
FOO_FOUR =BAR_FOUR

# Comment 3
FOO_FIVE=BAR_FIVE
FOO_SIX="BAR_SIX"
FOO_SEVEN="BAR_SEVEN" # Comment 4
FOO_EIGHT="BAR_\"EIGHT\""`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.WHITESPACE, " "},
		{token.COMMENT, "# Comment 1 "},
		{token.NEWLINE, "\n"},
		{token.KEY, "FOO_ONE"},
		{token.ASSIGN, "="},
		{token.PLAIN_VALUE, "BAR_ONE "},
		{token.COMMENT, "# Comment 2"},
		{token.NEWLINE, "\n"},
		{token.KEY, "FOO_TWO"},
		{token.WHITESPACE, " "},
		{token.ASSIGN, "="},
		{token.WHITESPACE, " "},
		{token.PLAIN_VALUE, "BAR_TWO"},
		{token.NEWLINE, "\n"},
		{token.KEY, "FOO_THREE"},
		{token.ASSIGN, "="},
		{token.WHITESPACE, " "},
		{token.PLAIN_VALUE, "BAR_THREE"},
		{token.NEWLINE, "\n"},
		{token.KEY, "FOO_FOUR"},
		{token.WHITESPACE, " "},
		{token.ASSIGN, "="},
		{token.PLAIN_VALUE, "BAR_FOUR"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.COMMENT, "# Comment 3"},
		{token.NEWLINE, "\n"},
		{token.KEY, "FOO_FIVE"},
		{token.ASSIGN, "="},
		{token.PLAIN_VALUE, "BAR_FIVE"},
		{token.NEWLINE, "\n"},
		{token.KEY, "FOO_SIX"},
		{token.ASSIGN, "="},
		{token.QUOTED_VALUE, "\"BAR_SIX\""},
		{token.NEWLINE, "\n"},
		{token.KEY, "FOO_SEVEN"},
		{token.ASSIGN, "="},
		{token.QUOTED_VALUE, "\"BAR_SEVEN\""},
		{token.WHITESPACE, " "},
		{token.COMMENT, "# Comment 4"},
		{token.NEWLINE, "\n"},
		{token.KEY, "FOO_EIGHT"},
		{token.ASSIGN, "="},
		{token.QUOTED_VALUE, `"BAR_\"EIGHT\""`},
		{token.EOF, ""},
	}
	l := lexer.New([]byte(input))

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q (%q)",
				i, tt.expectedType, tok.Type, tok.Literal)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
