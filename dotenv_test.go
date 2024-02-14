package dotenv_test

import (
	"os"
	"testing"

	"github.com/zeevallin/dotenv"
	"github.com/zeevallin/dotenv/ast"
)

func TestUnmarshal(t *testing.T) {
	t.Run("fixtures/comments.env", func(t *testing.T) {
		raw, err := os.ReadFile("fixtures/comments.env")
		if err != nil {
			t.Fatal(err)
		}
		actual, err := dotenv.Unmarshal(raw)
		if err != nil {
			t.Fatal(err)
		}

		expected := &ast.File{
			Lines: []*ast.Line{
				{
					RawComment: "# Full line comment",
				},
				{
					Key:        "one",
					RawValue:   "one ",
					RawComment: "# baz",
				},
				{
					Key:        "two",
					RawValue:   "two",
					RawComment: "#baz",
				},
				{
					Key:        "three",
					RawValue:   `"three"`,
					RawComment: "#bar",
				},
				{
					Key:        "four",
					RawValue:   "'four'",
					RawComment: "#bar",
				},
			},
		}

		if len(actual.Lines) != len(expected.Lines) {
			t.Fatalf("expected %d entries, got %d", len(expected.Lines), len(actual.Lines))
		}

		for i, line := range expected.Lines {
			compare(t, line, actual.Lines[i])
		}
	})
}

func compare(t testing.TB, a, b *ast.Line) {
	if a.Key != b.Key {
		t.Errorf("expected key %q, got %q", a.Key, b.Key)
	}
	if a.RawValue != b.RawValue {
		t.Errorf("expected value %q, got %q", a.RawValue, b.RawValue)
	}
	if a.RawComment != b.RawComment {
		t.Errorf("expected comment %q, got %q", a.RawComment, b.RawComment)
	}
}
