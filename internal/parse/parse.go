package parse

import (
	"encoding/xml"
	"io"
)

// ReadXML processes tokens from the given decoder and returns a slice of
// Elements corresponding to the tokens
func ReadXML(decoder *xml.Decoder) ([]Element, error) {
	stack := make([]*Element, 0)
	elements := make([]*Element, 0)

	for t, err := decoder.Token(); ; t, err = decoder.Token() {
		if err == io.EOF {
			return elementsCopy(elements), nil
		}
		if err != nil {
			return nil, err
		}

		depth := len(stack)

		switch token := t.(type) {
		case xml.StartElement:
			if depth > 0 {
				parent := stack[len(stack)-1]
				parent.ChildCount++
			}

			ele := Element{
				Token:      xml.CopyToken(token),
				Depth:      depth,
				ChildCount: 0,
			}
			elements = append(elements, &ele)

			stack = append(stack, &ele)
		case xml.EndElement:
			start := stack[len(stack)-1]

			stack = stack[:len(stack)-1]

			// No end is needed for empty nodes
			if start.ChildCount > 0 {
				newDepth := len(stack)

				ele := Element{
					Token:      xml.CopyToken(token),
					Depth:      newDepth,
					ChildCount: 0,
				}
				elements = append(elements, &ele)
			}
		case xml.CharData:
			// TODO: Need to handle this for string resources
		case xml.Comment:
			ele := Element{
				Token:      xml.CopyToken(token),
				Depth:      depth,
				ChildCount: 0,
			}
			elements = append(elements, &ele)
		case xml.ProcInst:
			ele := Element{
				Token:      xml.CopyToken(token),
				Depth:      depth,
				ChildCount: 0,
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
