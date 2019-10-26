package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/rsookram/axmlfmt/internal/format"
	"github.com/rsookram/axmlfmt/internal/parse"
)

func main() {
	r, err := os.Open("view_main.xml")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	decoder := xml.NewDecoder(r)

	indent := "    "

	elements, err := parse.ReadXML(decoder)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}
	format.PrintXml(elements, indent)
}
