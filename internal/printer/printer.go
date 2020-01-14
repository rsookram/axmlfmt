package printer

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/rsookram/axmlfmt/internal/parse"
)

type Printer struct {
	indent string
}

func New(indent string) Printer {
	return Printer{
		indent: indent,
	}
}

func (p Printer) Fprint(w io.Writer, elements []parse.Element) error {
	newLinePositions := determineNewLinePositions(elements)

	for i, ele := range elements {
		depth := ele.Depth

		var err error
		switch token := ele.Token.(type) {
		case xml.StartElement:
			attrs := sortAttrs(token.Attr)
			if isXLIFF(token.Name) {
				err = p.startXLIFF(w, attrs)
			} else {
				err = p.startElement(w, token.Name.Local, attrs, ele.IsSelfClosing, ele.ContainsCharData, depth)
			}
		case xml.EndElement:
			err = p.endElement(w, token.Name.Local, ele.ContainsCharData, depth, isXLIFF(token.Name))
		case xml.CharData:
			err = p.charData(w, string(token))
		case xml.Comment:
			err = p.comment(w, string(token), depth)
		case xml.ProcInst:
			err = printProcInst(w, token.Target, string(token.Inst))
		}

		if newLinePositions[i] {
			_, err = fmt.Fprintln(w)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// determineNewLinePositions returns whether a new line should be printed after
// the element at a given position
func determineNewLinePositions(elements []parse.Element) []bool {
	positions := make([]bool, len(elements))

	for i := 0; i < len(elements)-1; i++ {
		curr := elements[i].Token
		next := elements[i+1].Token

		switch curr.(type) {
		case xml.StartElement, xml.EndElement:
			switch next.(type) {
			case xml.StartElement, xml.Comment:
				positions[i] = true
			}
		}
	}

	return positions
}

func isXLIFF(name xml.Name) bool {
	return name.Space == "urn:oasis:names:tc:xliff:document:1.2" && name.Local == "g"
}

func (p Printer) startElement(w io.Writer, name string, attrs []xml.Attr, isSelfClosing, containsCharData bool, depth int) error {
	_, err := fmt.Fprint(w, duplicate(p.indent, depth))
	if err != nil {
		return err
	}

	// Elements without attrs look like `<requestFocus />` or `<resources>`
	// and elements with one attr look like
	// `<string name="app_name">` or `<menu xmlns:android="...">`
	hasAttrs := len(attrs) == 0
	isSingleAttr := len(attrs) == 1
	if hasAttrs {
		_, err = fmt.Fprintf(w, "<%s", name)
	} else {
		if isSingleAttr {
			_, err = fmt.Fprintf(w, "<%s", name)
		} else {
			_, err = fmt.Fprintf(w, "<%s\n", name)
		}
	}
	if err != nil {
		return err
	}

	attrIndent := duplicate(p.indent, depth+1)
	for i, a := range attrs {
		if isSingleAttr {
			_, err = fmt.Fprintf(w, " %s=\"%s\"", cleanAttrName(a.Name), a.Value)
		} else {
			_, err = fmt.Fprintf(w, "%s%s=\"%s\"", attrIndent, cleanAttrName(a.Name), a.Value)
		}

		// The last attribute is on the same line as the ">"
		if i != len(attrs)-1 {
			_, err = fmt.Fprintf(w, "\n")
		}

		if err != nil {
			return err
		}
	}

	if containsCharData {
		_, err = fmt.Fprintf(w, ">")
	} else if !isSelfClosing {
		_, err = fmt.Fprintf(w, ">\n")
	} else {
		_, err = fmt.Fprintf(w, " />\n")
	}

	return err
}

func (p Printer) startXLIFF(w io.Writer, attrs []xml.Attr) error {
	_, err := fmt.Fprintf(w, "<xliff:g")
	if err != nil {
		return err
	}

	for _, a := range attrs {
		_, err = fmt.Fprintf(w, " %s=\"%s\"", cleanAttrName(a.Name), a.Value)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(w, ">")
	return err
}

func (p Printer) endElement(w io.Writer, name string, containsCharData bool, depth int, isXLIFF bool) error {
	if !containsCharData {
		_, err := fmt.Fprint(w, duplicate(p.indent, depth))
		if err != nil {
			return err
		}
	}

	var err error
	if isXLIFF {
		_, err = fmt.Fprintf(w, "</xliff:%s>", name)
	} else {
		_, err = fmt.Fprintf(w, "</%s>\n", name)
	}
	return err
}

func (p Printer) charData(w io.Writer, value string) error {
	_, err := fmt.Fprint(w, value)
	return err
}

func (p Printer) comment(w io.Writer, body string, depth int) error {
	_, err := fmt.Fprint(w, duplicate(p.indent, depth))
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "<!--%s-->\n", body)
	return err
}

func printProcInst(w io.Writer, target string, inst string) error {
	_, err := fmt.Fprintf(w, "<?%s %s?>\n", target, inst)
	return err
}

func cleanAttrName(n xml.Name) string {
	space := n.Space
	if space == "" {
		// Attributes not in a namespace such as style
		return n.Local
	}

	switch space {
	case "http://schemas.android.com/apk/res/android":
		space = "android"
	case "http://schemas.android.com/apk/res-auto":
		space = "app"
	case "http://schemas.android.com/tools":
		space = "tools"
	}

	return fmt.Sprintf("%s:%s", space, n.Local)
}

func duplicate(s string, n int) string {
	ret := ""

	for i := 0; i < n; i++ {
		ret += s
	}

	return ret
}
