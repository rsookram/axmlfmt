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

var write = flag.Bool("w", false, "write result to (source) file instead of stdout")

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
		err = writeOutput(p, elements, *write, name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(3)
		}
	}
}

func writeOutput(p printer.Printer, elements []parse.Element, write bool, inputFileName string) error {
	if !write {
		p.Fprint(os.Stdout, elements)
		return nil
	}

	f, err := os.Create(inputFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	p.Fprint(f, elements)

	return nil
}
