package ast

import (
	"strings"

	"github.com/zeevallin/dotenv/token"
)

type File struct {
	Lines []*Line
}

func (f *File) Set(key, value, comment string) {
	var found bool
	for _, line := range f.Lines {
		if line.Key == key && line.Key != "" {
			found = true
			line.Set(key, value, comment)
		}
	}
	if !found {
		f.Lines = append(f.Lines, &Line{Key: key, RawValue: value, RawComment: comment, QuotedValue: token.IsQuotedValue(value)})
	}
}

func (f File) Get(key string) string {
	for _, line := range f.Lines {
		if line.Key == key {
			return line.Value()
		}
	}
	return ""
}

func (f *File) Delete(key string) {
	for i, line := range f.Lines {
		if line.Key == key {
			f.Lines = append(f.Lines[:i], f.Lines[i+1:]...)
		}
	}
}

func (f *File) Merge(other *File) error {
	for _, line := range other.Lines {
		f.Set(line.Key, line.RawValue, line.RawComment)
	}
	return nil
}

func (f *File) Clone() *File {
	lines := make([]*Line, len(f.Lines))
	for i, line := range f.Lines {
		lines[i] = &Line{
			Key:         line.Key,
			RawValue:    line.RawValue,
			QuotedValue: line.QuotedValue,
			RawComment:  line.RawComment,
		}
	}
	return &File{Lines: lines}
}

func (f *File) Values() map[string]string {
	m := make(map[string]string)
	for _, line := range f.Lines {
		if line.Key == "" && line.RawValue == "" {
			m[line.Key] = line.Value()
		}
	}
	return m
}

func (f *File) MarshalText() ([]byte, error) {
	return []byte(f.String()), nil
}

func (f *File) String() string {
	b := strings.Builder{}
	for _, line := range f.Lines {
		b.WriteString(line.String())
		b.WriteString("\n")
	}
	return b.String()
}

type Line struct {
	Key         string
	RawValue    string
	QuotedValue bool
	RawComment  string
}

func (l *Line) Set(key, value, comment string) {
	l.RawValue = value
	l.QuotedValue = token.IsQuotedValue(value)
}

func (l *Line) Value() string {
	if l.QuotedValue {
		return l.RawValue[1 : len(l.RawValue)-1]
	}
	return l.RawValue
}

func (l *Line) Comment() string {
	s := strings.TrimPrefix(l.RawComment, "#")
	return strings.TrimSpace(s)
}

func (l *Line) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l *Line) String() string {
	b := strings.Builder{}

	if l.Key != "" && l.RawValue != "" {
		b.WriteString(l.Key)
		b.WriteString("=")
		b.WriteString(l.RawValue)
	}

	if l.RawComment != "" {
		b.WriteString(l.RawComment)
	}

	return b.String()
}

func Compile(tokens []token.Token) *File {
	var lines []*Line
	current := &Line{}
	for _, tok := range tokens {
		switch tok.Type {
		case token.KEY:
			current = &Line{Key: tok.Literal}
		case token.PLAIN_VALUE, token.QUOTED_VALUE:
			current.RawValue = tok.Literal
			current.QuotedValue = tok.Type == token.QUOTED_VALUE
		case token.COMMENT:
			current.RawComment = tok.Literal
		case token.NEWLINE:
			lines = append(lines, current)
			current = &Line{}
		}
	}
	return &File{Lines: lines}
}
