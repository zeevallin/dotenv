package ast_test

import (
	"testing"

	"github.com/zeevallin/dotenv/ast"
	"github.com/zeevallin/dotenv/token"
)

func TestCompile(t *testing.T) {
	input := []token.Token{
		{Type: token.WHITESPACE, Literal: " "},
		{Type: token.COMMENT, Literal: "# Comment 1 "},
		{Type: token.NEWLINE, Literal: "\n"},
		{Type: token.KEY, Literal: "FOO_ONE"},
		{Type: token.EQUALS, Literal: "="},
		{Type: token.PLAIN_VALUE, Literal: "BAR_ONE "},
		{Type: token.COMMENT, Literal: "# Comment 2"},
		{Type: token.NEWLINE, Literal: "\n"},
		{Type: token.KEY, Literal: "FOO_TWO"},
		{Type: token.WHITESPACE, Literal: " "},
		{Type: token.EQUALS, Literal: "="},
		{Type: token.WHITESPACE, Literal: " "},
		{Type: token.PLAIN_VALUE, Literal: "BAR_TWO"},
		{Type: token.NEWLINE, Literal: "\n"},
		{Type: token.KEY, Literal: "FOO_THREE"},
		{Type: token.EQUALS, Literal: "="},
		{Type: token.WHITESPACE, Literal: " "},
		{Type: token.PLAIN_VALUE, Literal: "BAR_THREE"},
		{Type: token.NEWLINE, Literal: "\n"},
		{Type: token.KEY, Literal: "FOO_FOUR"},
		{Type: token.WHITESPACE, Literal: " "},
		{Type: token.EQUALS, Literal: "="},
		{Type: token.PLAIN_VALUE, Literal: "BAR_FOUR"},
		{Type: token.NEWLINE, Literal: "\n"},
		{Type: token.NEWLINE, Literal: "\n"},
		{Type: token.COMMENT, Literal: "# Comment 3"},
		{Type: token.NEWLINE, Literal: "\n"},
		{Type: token.KEY, Literal: "FOO_FIVE"},
		{Type: token.EQUALS, Literal: "="},
		{Type: token.PLAIN_VALUE, Literal: "BAR_FIVE"},
		{Type: token.NEWLINE, Literal: "\n"},
		{Type: token.KEY, Literal: "FOO_SIX"},
		{Type: token.EQUALS, Literal: "="},
		{Type: token.QUOTED_VALUE, Literal: "\"BAR_SIX\""},
		{Type: token.NEWLINE, Literal: "\n"},
		{Type: token.KEY, Literal: "FOO_SEVEN"},
		{Type: token.EQUALS, Literal: "="},
		{Type: token.QUOTED_VALUE, Literal: "\"BAR_SEVEN\""},
		{Type: token.WHITESPACE, Literal: " "},
		{Type: token.COMMENT, Literal: "# Comment 4"},
		{Type: token.NEWLINE, Literal: "\n"},
		{Type: token.KEY, Literal: "FOO_EIGHT"},
		{Type: token.EQUALS, Literal: "="},
		{Type: token.QUOTED_VALUE, Literal: "\"BAR_\\\"EIGHT\\\"\""},
		{Type: token.NEWLINE, Literal: "\n"},
		{Type: token.EOF, Literal: ""},
	}

	expected := ast.File{
		Lines: []*ast.Line{
			{Key: "", RawValue: "", QuotedValue: false, RawComment: "# Comment 1 "},                     // test 1
			{Key: "FOO_ONE", RawValue: "BAR_ONE ", QuotedValue: false, RawComment: "# Comment 2"},       // test 2
			{Key: "FOO_TWO", RawValue: "BAR_TWO", QuotedValue: false, RawComment: ""},                   // test 3
			{Key: "FOO_THREE", RawValue: "BAR_THREE", QuotedValue: false, RawComment: ""},               // test 4
			{Key: "FOO_FOUR", RawValue: "BAR_FOUR", QuotedValue: false, RawComment: ""},                 // test 5
			{Key: "", RawValue: "", QuotedValue: false, RawComment: ""},                                 // test 6
			{Key: "", RawValue: "", QuotedValue: false, RawComment: "# Comment 3"},                      // test 7
			{Key: "FOO_FIVE", RawValue: "BAR_FIVE", QuotedValue: false, RawComment: ""},                 // test 8
			{Key: "FOO_SIX", RawValue: "\"BAR_SIX\"", QuotedValue: true, RawComment: ""},                // test 9
			{Key: "FOO_SEVEN", RawValue: "\"BAR_SEVEN\"", QuotedValue: true, RawComment: "# Comment 4"}, // test 10
			{Key: "FOO_EIGHT", RawValue: "\"BAR_\\\"EIGHT\\\"\"", QuotedValue: true, RawComment: ""},    // test 11
		},
	}

	file := ast.Compile(input)

	if len(file.Lines) != len(expected.Lines) {
		t.Fatalf("expected %d lines, got %d", len(expected.Lines), len(file.Lines))
	}

	for i, line := range file.Lines {
		if line.Key != expected.Lines[i].Key {
			t.Fatalf("[test %d] expected key %q, got %q", i, expected.Lines[i].Key, line.Key)
		}
		if line.RawValue != expected.Lines[i].RawValue {
			t.Fatalf("[test %d] expected raw value %q, got %q", i, expected.Lines[i].RawValue, line.RawValue)
		}
		if line.QuotedValue != expected.Lines[i].QuotedValue {
			t.Fatalf("[test %d] expected quoted value %t, got %t", i, expected.Lines[i].QuotedValue, line.QuotedValue)
		}
		if line.RawComment != expected.Lines[i].RawComment {
			t.Fatalf("[test %d] expected comment %q, got %q", i, expected.Lines[i].RawComment, line.RawComment)
		}
	}
}
