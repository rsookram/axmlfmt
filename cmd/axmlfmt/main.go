package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/rsookram/axmlfmt/internal/parse"
	"github.com/rsookram/axmlfmt/internal/printer"
)

func main() {
	r, err := os.Open("view_main.xml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	decoder := xml.NewDecoder(r)

	indent := "    "

	elements, err := parse.ReadXML(decoder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(2)
	}
	printer.PrintXml(elements, indent)
}
