package printer

import (
	"encoding/xml"
	"strings"
	"testing"
)

const indent = "    "

func TestStartXLIFF(t *testing.T) {
	p := New(indent)

	w := &strings.Builder{}
	p.startXLIFF(
		w,
		[]xml.Attr{
			{Name: xml.Name{Space: "", Local: "example"}, Value: "2"},
			{Name: xml.Name{Space: "", Local: "id"}, Value: "quantity"},
		},
	)

	expected := `<xliff:g example="2" id="quantity">`
	if w.String() != expected {
		t.Errorf("got: %s, want %s", w.String(), expected)
	}
}

func TestEndXLIFF(t *testing.T) {
	p := New(indent)

	w := &strings.Builder{}
	p.endElement(w, "g", true, 2, true)

	expected := "</xliff:g>"
	if w.String() != expected {
		t.Errorf("got: %s, want %s", w.String(), expected)
	}
}

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

func TestProcInst(t *testing.T) {
	w := &strings.Builder{}
	printProcInst(w, "xml", `version="1.0" encoding="utf-8"`)

	expected := `<?xml version="1.0" encoding="utf-8"?>` + "\n"
	if w.String() != expected {
		t.Errorf("got: %s, want %s", w.String(), expected)
	}
}
