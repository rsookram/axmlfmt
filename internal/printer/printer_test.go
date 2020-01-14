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
	err := p.startXLIFF(
		w,
		[]xml.Attr{
			{Name: xml.Name{Space: "", Local: "example"}, Value: "2"},
			{Name: xml.Name{Space: "", Local: "id"}, Value: "quantity"},
		},
	)
	requireNoError(t, err)

	expected := `<xliff:g example="2" id="quantity">`
	if w.String() != expected {
		t.Errorf("got: %s, want %s", w.String(), expected)
	}
}

func TestEndXLIFF(t *testing.T) {
	p := New(indent)

	w := &strings.Builder{}
	err := p.endElement(w, "g", true, 2, true)
	requireNoError(t, err)

	expected := "</xliff:g>"
	if w.String() != expected {
		t.Errorf("got: %s, want %s", w.String(), expected)
	}
}

func TestCharData(t *testing.T) {
	p := New(indent)

	w := &strings.Builder{}
	err := p.charData(w, "a string")
	requireNoError(t, err)

	expected := "a string"
	if w.String() != expected {
		t.Errorf("got: %s, want %s", w.String(), expected)
	}
}

func TestComment(t *testing.T) {
	p := New(indent)

	{
		w := &strings.Builder{}
		err := p.comment(w, "a comment", 0)
		requireNoError(t, err)

		expected := "<!--a comment-->\n"
		if w.String() != expected {
			t.Errorf("got: %s, want %s", w.String(), expected)
		}
	}

	{
		w := &strings.Builder{}
		err := p.comment(w, "a comment", 1)
		requireNoError(t, err)

		expected := indent + "<!--a comment-->\n"
		if w.String() != expected {
			t.Errorf("got: %s, want %s", w.String(), expected)
		}
	}
}

func TestProcInst(t *testing.T) {
	w := &strings.Builder{}
	err := printProcInst(w, "xml", `version="1.0" encoding="utf-8"`)
	requireNoError(t, err)

	expected := `<?xml version="1.0" encoding="utf-8"?>` + "\n"
	if w.String() != expected {
		t.Errorf("got: %s, want %s", w.String(), expected)
	}
}

func requireNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
}
