package parse

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

const androidNS = "http://schemas.android.com/apk/res/android"

func TestEmptyXML(t *testing.T) {
	doc := ""

	ee, err := read(doc)
	if err != nil {
		t.Errorf("got %s", err)
	}

	if len(ee) != 0 {
		t.Errorf("got %d, want %d", len(ee), 0)
	}
}

func TestElementContainingCommentNoChildren(t *testing.T) {
	doc := `
<shape
    xmlns:android="http://schemas.android.com/apk/res/android"
    android:shape="rectangle">

    <!-- no children -->
</shape>
	`

	ee, err := read(doc)
	if err != nil {
		t.Errorf("got %s", err)
	}

	expected := []Element{
		{
			Token: xml.StartElement{
				Name: tagName("", "shape"),
				Attr: []xml.Attr{
					attr("xmlns", "android", androidNS),
					attr(androidNS, "shape", "rectangle"),
				}},
		},
		{
			Token: xml.Comment(" no children "),
			Depth: 1,
		},
		{
			Token: endElement("", "shape"),
		},
	}

	if !equal(expected, ee) {
		t.Errorf("got %s, want %s", str(ee), str(expected))
	}
}

func read(doc string) ([]Element, error) {
	d := xml.NewDecoder(strings.NewReader(doc))
	return ReadXML(d)
}

func attr(space, local, value string) xml.Attr {
	return xml.Attr{Name: tagName(space, local), Value: value}
}

func endElement(space, local string) xml.EndElement {
	return xml.EndElement{Name: tagName(space, local)}
}

func tagName(space, local string) xml.Name {
	return xml.Name{Space: space, Local: local}
}

func equal(expected, actual []Element) bool {
	return reflect.DeepEqual(expected, actual)
}

func str(ee []Element) string {
	s, _ := json.MarshalIndent(ee, "", "  ")
	return string(s)
}
