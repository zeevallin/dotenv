package dotenv

import (
	"bytes"
	"io"

	"github.com/zeevallin/dotenv/ast"
	"github.com/zeevallin/dotenv/lexer"
	"github.com/zeevallin/dotenv/token"
)

func Unmarshal(raw []byte) (*ast.File, error) {
	return Read(bytes.NewReader(raw))
}

func Read(r io.Reader) (*ast.File, error) {
	raw, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	lexer := lexer.New(raw)

	tokens := []token.Token{}
	for {
		tok := lexer.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == token.EOF {
			break
		}
	}

	return ast.Compile(tokens), nil
}
