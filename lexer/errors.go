package lexer

import "fmt"

type ParseError struct {
	Line int
	Err  error
}

func (e ParseError) Error() string {
	return fmt.Sprintf("error on line %d: %s", e.Line, e.Err)
}

func (e ParseError) Unwrap() error {
	return e.Err
}
