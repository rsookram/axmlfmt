package printer

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/rsookram/axmlfmt/internal/parse"
)

const indent = "    "

func TestStartElement(t *testing.T) {
	p := New(indent)

	{
		w := &strings.Builder{}
		ee := []parse.Element{
			{
				Token: xml.StartElement{
					Name: xml.Name{
						Space: "",
						Local: "resources",
					},
					Attr: []xml.Attr{},
				},
				Depth:            0,
				IsSelfClosing:    false,
				ContainsCharData: false,
			},
		}

		err := p.Fprint(w, ee)
		requireNoError(t, err)

		expected := "<resources>\n"
		if w.String() != expected {
			t.Errorf("got: %s, want %s", w.String(), expected)
		}
	}
}

func TestStartXLIFF(t *testing.T) {
	p := New(indent)

	w := &strings.Builder{}
	ee := []parse.Element{
		{
			Token: xml.StartElement{
				Name: xml.Name{
					Space: "urn:oasis:names:tc:xliff:document:1.2",
					Local: "g",
				},
				Attr: []xml.Attr{
					{Name: xml.Name{Space: "", Local: "example"}, Value: "2"},
					{Name: xml.Name{Space: "", Local: "id"}, Value: "quantity"},
				},
			},
			Depth:            1,
			IsSelfClosing:    false,
			ContainsCharData: true,
		},
	}

	err := p.Fprint(w, ee)
	requireNoError(t, err)

	expected := `<xliff:g example="2" id="quantity">`
	if w.String() != expected {
		t.Errorf("got: %s, want %s", w.String(), expected)
	}
}

func TestEndElement(t *testing.T) {
	p := New(indent)

	w := &strings.Builder{}
	ee := []parse.Element{
		{
			Token: xml.EndElement{
				Name: xml.Name{
					Space: "",
					Local: "androidx.constraintlayout.widget.ConstraintLayout",
				},
			},
			Depth:            1,
			IsSelfClosing:    false,
			ContainsCharData: false,
		},
	}

	err := p.Fprint(w, ee)
	requireNoError(t, err)

	expected := indent + "</androidx.constraintlayout.widget.ConstraintLayout>\n"
	if w.String() != expected {
		t.Errorf("got: %s, want %s", w.String(), expected)
	}
}

func TestEndXLIFF(t *testing.T) {
	p := New(indent)

	w := &strings.Builder{}
	ee := []parse.Element{
		{
			Token: xml.EndElement{
				Name: xml.Name{
					Space: "urn:oasis:names:tc:xliff:document:1.2",
					Local: "g",
				},
			},
			Depth:            2,
			IsSelfClosing:    false,
			ContainsCharData: true,
		},
	}

	err := p.Fprint(w, ee)
	requireNoError(t, err)

	expected := "</xliff:g>"
	if w.String() != expected {
		t.Errorf("got: %s, want %s", w.String(), expected)
	}
}

func TestCharData(t *testing.T) {
	p := New(indent)

	w := &strings.Builder{}
	ee := []parse.Element{
		{
			Token:            xml.CharData("a string"),
			Depth:            1,
			IsSelfClosing:    false,
			ContainsCharData: false,
		},
	}

	err := p.Fprint(w, ee)
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
		ee := []parse.Element{
			{
				Token:            xml.Comment("a comment"),
				Depth:            0,
				IsSelfClosing:    false,
				ContainsCharData: false,
			},
		}

		err := p.Fprint(w, ee)
		requireNoError(t, err)

		expected := "<!--a comment-->\n"
		if w.String() != expected {
			t.Errorf("got: %s, want %s", w.String(), expected)
		}
	}

	{
		w := &strings.Builder{}
		ee := []parse.Element{
			{
				Token:            xml.Comment("a comment"),
				Depth:            1,
				IsSelfClosing:    false,
				ContainsCharData: false,
			},
		}

		err := p.Fprint(w, ee)
		requireNoError(t, err)

		expected := indent + "<!--a comment-->\n"
		if w.String() != expected {
			t.Errorf("got: %s, want %s", w.String(), expected)
		}
	}
}

func TestProcInst(t *testing.T) {
	p := New(indent)

	w := &strings.Builder{}
	ee := []parse.Element{
		{
			Token: xml.ProcInst{
				Target: "xml",
				Inst:   []byte(`version="1.0" encoding="utf-8"`),
			},
			Depth:            0,
			IsSelfClosing:    false,
			ContainsCharData: false,
		},
	}

	err := p.Fprint(w, ee)
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
