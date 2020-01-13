package parse

import (
	"encoding/xml"
	"io"
	"strings"
)

// ReadXML processes tokens from the given reader and returns a slice of
// Elements corresponding to the tokens
func ReadXML(reader xml.TokenReader) ([]Element, error) {
	stack := make([]*Element, 0)
	elements := make([]*Element, 0)

	for t, err := reader.Token(); ; t, err = reader.Token() {
		if err == io.EOF {
			return elementsCopy(elements), nil
		}
		if err != nil {
			return nil, err
		}

		depth := len(stack)

		switch token := t.(type) {
		case xml.StartElement:
			containsCharData := false
			if depth > 0 {
				parent := stack[len(stack)-1]
				parent.IsSelfClosing = false
				containsCharData = parent.ContainsCharData
			}

			ele := Element{
				Token:            xml.CopyToken(token),
				Depth:            depth,
				IsSelfClosing:    true,
				ContainsCharData: containsCharData,
			}
			elements = append(elements, &ele)

			stack = append(stack, &ele)
		case xml.EndElement:
			start := stack[len(stack)-1]

			stack = stack[:len(stack)-1]

			// No end is needed for empty nodes
			if !start.IsSelfClosing || start.ContainsCharData {
				newDepth := len(stack)

				ele := Element{
					Token:            xml.CopyToken(token),
					Depth:            newDepth,
					IsSelfClosing:    false,
					ContainsCharData: start.ContainsCharData,
				}
				elements = append(elements, &ele)
			}
		case xml.CharData:
			s := string(token)
			if len(strings.TrimSpace(s)) != 0 {
				parent := stack[len(stack)-1]
				parent.ContainsCharData = true

				ele := Element{
					Token:         xml.CopyToken(xml.CharData(s)),
					Depth:         len(stack),
					IsSelfClosing: false,
				}
				elements = append(elements, &ele)
			}
		case xml.Comment:
			if depth > 0 {
				parent := stack[len(stack)-1]
				parent.IsSelfClosing = false
			}

			ele := Element{
				Token:         xml.CopyToken(token),
				Depth:         depth,
				IsSelfClosing: false,
			}
			elements = append(elements, &ele)
		case xml.ProcInst:
			ele := Element{
				Token:         xml.CopyToken(token),
				Depth:         depth,
				IsSelfClosing: false,
			}
			elements = append(elements, &ele)
		}
	}
}

func elementsCopy(ele []*Element) []Element {
	eleCopy := make([]Element, len(ele))
	for i, e := range ele {
		eleCopy[i] = *e
	}

	return eleCopy
}
