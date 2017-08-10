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

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func main() {
	r, err := os.Open("view_main.xml")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	decoder := xml.NewDecoder(r)

	indent := "    "
	formatXml(decoder, indent)
}

func formatXml(decoder *xml.Decoder, indent string) {
	stack := make([]xml.StartElement, 0)

	for t, err := decoder.Token(); ; t, err = decoder.Token() {
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			return
		}

		depth := len(stack)

		switch token := t.(type) {
		case xml.StartElement:
			fmt.Printf(duplicate(indent, depth))
			fmt.Printf("<%s\n", token.Name.Local)

			attrIndent := duplicate(indent, depth+1)
			attrs := token.Attr
			for i, a := range attrs {
				fmt.Printf("%s%s=\"%s\"", attrIndent, cleanAttrName(a.Name), a.Value)
				if i != len(attrs)-1 {
					fmt.Printf("\n")
				}
			}
			fmt.Printf(">\n")

			stack = append(stack, token)
		case xml.EndElement:
			start := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			newDepth := len(stack)
			fmt.Printf(duplicate(indent, newDepth))
			fmt.Printf("</%s>\n", start.Name.Local)
		case xml.CharData:
			// TODO: Need to handle this for string resources
		case xml.Comment:
			fmt.Printf("\n")
			fmt.Printf(duplicate(indent, depth))
			fmt.Printf("<--%s-->\n", string(token))
		case xml.ProcInst:
			fmt.Printf("<?%s %s?>\n", token.Target, string(token.Inst))
		}
	}
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
