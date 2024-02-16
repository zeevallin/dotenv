package dotenv

import (
	"io"
	"os"

	"github.com/zeevallin/dotenv/ast"
	"github.com/zeevallin/dotenv/lexer"
	"github.com/zeevallin/dotenv/token"
)

type Parser struct {
	AllowDuplicates   bool
	AllowQuotedValues bool
}

type File struct {
	Lines []Line
}

type Line struct {
	Key     string
	Value   string
	Comment string
}

func (p Parser) Parse(b []byte) (*ast.File, error) {
	lexer := lexer.New(b)
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

var DefaultParser = Parser{
	AllowDuplicates:   true,
	AllowQuotedValues: true,
}

func Read(r io.Reader) (*ast.File, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return DefaultParser.Parse(b)
}

func Parse(b []byte) (*ast.File, error) {
	return DefaultParser.Parse(b)
}

func Open(filename string) (*ast.File, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return DefaultParser.Parse(b)
}
