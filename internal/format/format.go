/*
Copyright 2017 Rashad Sookram

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package format

import (
	"encoding/xml"
	"fmt"

	"github.com/rsookram/axmlfmt/internal/parse"
)

func PrintXml(elements []parse.Element, indent string) {
	newLinePositions := determineNewLinePositions(elements)

	for i, ele := range elements {
		depth := ele.Depth

		switch token := ele.Token.(type) {
		case xml.StartElement:
			printStartElement(token.Name.Local, token.Attr, ele.ChildCount > 0, depth, indent)
		case xml.EndElement:
			printEndElement(token.Name.Local, depth, indent)
		case xml.CharData:
			// TODO: Need to handle this for string resources
		case xml.Comment:
			printComment(string(token), depth, indent)
		case xml.ProcInst:
			printProcInst(token.Target, string(token.Inst))
		}

		if newLinePositions[i] {
			fmt.Println()
		}
	}
}

// Returns whether a new lines should be printed after the element at a given
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

func printStartElement(name string, attrs []xml.Attr, hasChildren bool, depth int, indent string) {
	fmt.Printf(duplicate(indent, depth))
	fmt.Printf("<%s\n", name)

	attrIndent := duplicate(indent, depth+1)
	for i, a := range attrs {
		fmt.Printf("%s%s=\"%s\"", attrIndent, cleanAttrName(a.Name), a.Value)
		if i != len(attrs)-1 {
			fmt.Printf("\n")
		}
	}

	if hasChildren {
		fmt.Printf(">\n")
	} else {
		fmt.Printf("/>\n")
	}
}

func printEndElement(name string, depth int, indent string) {
	fmt.Printf(duplicate(indent, depth))
	fmt.Printf("</%s>\n", name)
}

func printComment(body string, depth int, indent string) {
	fmt.Printf(duplicate(indent, depth))
	fmt.Printf("<--%s-->\n", body)
}

func printProcInst(target string, inst string) {
	fmt.Printf("<?%s %s?>\n", target, inst)
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
