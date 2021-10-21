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

	expected := indent + `<xliff:g example="2" id="quantity">`
	if w.String() != expected {
		t.Errorf("got: %s, want %s", w.String(), expected)
	}
}

func TestStartAAPT(t *testing.T) {
	p := New(indent)

	w := &strings.Builder{}
	ee := []parse.Element{
		{
			Token: xml.StartElement{
				Name: xml.Name{
					Space: "http://schemas.android.com/aapt",
					Local: "attr",
				},
				Attr: []xml.Attr{
					{Name: xml.Name{Space: "", Local: "name"}, Value: "android:fillColor"},
				},
			},
			Depth:            2,
			IsSelfClosing:    false,
			ContainsCharData: false,
		},
	}

	err := p.Fprint(w, ee)
	requireNoError(t, err)

	expected := indent + indent + `<aapt:attr name="android:fillColor">` + "\n"
	if w.String() != expected {
		t.Errorf("got: %s, want %s", w.String(), expected)
	}
}

func TestStandardizeNamespaceName(t *testing.T) {
	p := New(indent)

	w := &strings.Builder{}
	ee := []parse.Element{
		{
			Token: xml.StartElement{
				Name: xml.Name{
					Space: "",
					Local: "androidx.cardview.widget.CardView",
				},
				Attr: []xml.Attr{
					{Name: xml.Name{Space: "xmlns", Local: "android"}, Value: "http://schemas.android.com/apk/res/android"},
					{Name: xml.Name{Space: "xmlns", Local: "card_view"}, Value: "http://schemas.android.com/apk/res-auto"},
					{Name: xml.Name{Space: "http://schemas.android.com/apk/res/android", Local: "layout_width"}, Value: "match_parent"},
					{Name: xml.Name{Space: "http://schemas.android.com/apk/res/android", Local: "layout_height"}, Value: "wrap_content"},
					{Name: xml.Name{Space: "http://schemas.android.com/apk/res-auto", Local: "cardCornerRadius"}, Value: "4dp"},
				},
			},
			Depth:            0,
			IsSelfClosing:    false,
			ContainsCharData: false,
		},
	}

	err := p.Fprint(w, ee)
	requireNoError(t, err)

	expected := `<androidx.cardview.widget.CardView
    xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:app="http://schemas.android.com/apk/res-auto"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    app:cardCornerRadius="4dp">
`
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

func TestEndAAPT(t *testing.T) {
	p := New(indent)

	w := &strings.Builder{}
	ee := []parse.Element{
		{
			Token: xml.EndElement{
				Name: xml.Name{
					Space: "http://schemas.android.com/aapt",
					Local: "attr",
				},
			},
			Depth:            2,
			IsSelfClosing:    false,
			ContainsCharData: false,
		},
	}

	err := p.Fprint(w, ee)
	requireNoError(t, err)

	expected := indent + indent + "</aapt:attr>"
	if w.String() != expected {
		t.Errorf("got: %s, want %s", w.String(), expected)
	}
}

func TestCharData(t *testing.T) {
	p := New(indent)

	{
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

	{
		w := &strings.Builder{}
		ee := []parse.Element{
			{
				Token:            xml.CharData("<b> & </b>"),
				Depth:            1,
				IsSelfClosing:    false,
				ContainsCharData: false,
			},
		}

		err := p.Fprint(w, ee)
		requireNoError(t, err)

		expected := "&lt;b&gt; &amp; &lt;/b&gt;"
		if w.String() != expected {
			t.Errorf("got: %s, want %s", w.String(), expected)
		}
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
