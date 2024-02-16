package lexer2_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/zeevallin/dotenv/lexer2"
)

func TestLexer_NextToken(t *testing.T) {
	l := lexer2.New("r1.env", "  hello\nworld ")

	var tokens []lexer2.Token
	for token := l.NextToken(); token.Kind != lexer2.EOF; token = l.NextToken() {
		tokens = append(tokens, token)
	}

	spew.Dump(tokens)

	t.Fail()
}
