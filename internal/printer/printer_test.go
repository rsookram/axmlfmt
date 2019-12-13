package printer

import (
	"strings"
	"testing"
)

const indent = "    "

func TestCharData(t *testing.T) {
	p := New(indent)

	w := &strings.Builder{}
	p.charData(w, "a string")

	expected := "a string"
	if w.String() != expected {
		t.Errorf("got: %s, want %s", w.String(), expected)
	}
}

func TestComment(t *testing.T) {
	p := New(indent)

	{
		w := &strings.Builder{}
		p.comment(w, "a comment", 0)

		expected := "<!--a comment-->\n"
		if w.String() != expected {
			t.Errorf("got: %s, want %s", w.String(), expected)
		}
	}

	{
		w := &strings.Builder{}
		p.comment(w, "a comment", 1)

		expected := indent + "<!--a comment-->\n"
		if w.String() != expected {
			t.Errorf("got: %s, want %s", w.String(), expected)
		}
	}
}
