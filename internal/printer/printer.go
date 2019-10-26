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

func (p Printer) Fprint(w io.Writer, elements []parse.Element) {
	newLinePositions := determineNewLinePositions(elements)

	for i, ele := range elements {
		depth := ele.Depth

		switch token := ele.Token.(type) {
		case xml.StartElement:
			p.startElement(w, token.Name.Local, sortAttrs(token.Attr), ele.ChildCount > 0, depth)
		case xml.EndElement:
			p.endElement(w, token.Name.Local, depth)
		case xml.CharData:
			// TODO: Need to handle this for string resources
		case xml.Comment:
			p.comment(w, string(token), depth)
		case xml.ProcInst:
			printProcInst(w, token.Target, string(token.Inst))
		}

		if newLinePositions[i] {
			fmt.Fprintln(w)
		}
	}
}

// Returns whether a new line should be printed after the element at a given
// position
func determineNewLinePositions(elements []parse.Element) []bool {
	positions := make([]bool, len(elements))

	for i := 0; i < len(elements)-1; i++ {
		curr := elements[i].Token
		next := elements[i+1].Token

		switch c := curr.(type) {
		case xml.StartElement:
			switch n := next.(type) {
			case xml.StartElement, xml.Comment:
				positions[i] = true
			default:
				_ = n
			}

		default:
			_ = c
		}
	}

	return positions
}

func (p Printer) startElement(w io.Writer, name string, attrs []xml.Attr, hasChildren bool, depth int) {
	fmt.Fprintf(w, duplicate(p.indent, depth))
	fmt.Fprintf(w, "<%s\n", name)

	attrIndent := duplicate(p.indent, depth+1)
	for i, a := range attrs {
		fmt.Fprintf(w, "%s%s=\"%s\"", attrIndent, cleanAttrName(a.Name), a.Value)
		if i != len(attrs)-1 {
			fmt.Fprintf(w, "\n")
		}
	}

	if hasChildren {
		fmt.Fprintf(w, ">\n")
	} else {
		fmt.Fprintf(w, "/>\n")
	}
}

func (p Printer) endElement(w io.Writer, name string, depth int) {
	fmt.Fprintf(w, duplicate(p.indent, depth))
	fmt.Fprintf(w, "</%s>\n", name)
}

func (p Printer) comment(w io.Writer, body string, depth int) {
	fmt.Fprintf(w, duplicate(p.indent, depth))
	fmt.Fprintf(w, "<--%s-->\n", body)
}

func printProcInst(w io.Writer, target string, inst string) {
	fmt.Fprintf(w, "<?%s %s?>\n", target, inst)
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
