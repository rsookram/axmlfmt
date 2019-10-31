package parse

import "encoding/xml"

type Element struct {
	Token            xml.Token
	Depth            int
	IsSelfClosing    bool
	ContainsCharData bool
}
