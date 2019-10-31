package parse

import "encoding/xml"

type Element struct {
	Token       xml.Token
	Depth       int
	ChildCount  int
	HasCharData bool
}
