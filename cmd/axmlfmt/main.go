package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"

	"github.com/rsookram/axmlfmt/internal/parse"
	"github.com/rsookram/axmlfmt/internal/printer"
)

const indent = "    "

func main() {
	flag.Parse()

	filenames := flag.Args()

	for _, name := range filenames {
		r, err := os.Open(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}

		decoder := xml.NewDecoder(r)

		elements, err := parse.ReadXML(decoder)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(2)
		}

		p := printer.New(indent)
		p.Fprint(os.Stdout, elements)
	}
}
