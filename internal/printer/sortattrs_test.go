package printer

import (
	"encoding/xml"
	"fmt"
	"strings"
	"testing"
)

const androidNS = "http://schemas.android.com/apk/res/android"
const appNS = "http://schemas.android.com/apk/res-auto"
const toolsNS = "http://schemas.android.com/tools"

func TestNamespace(t *testing.T) {
	attrs := []xml.Attr{
		attr("xmlns", "unknown", "http://schemas.android.com/apk/custom"),
		attr("xmlns", "app", appNS),
		attr("xmlns", "tools", toolsNS),
		attr("xmlns", "android", androidNS),
	}

	sorted := sortAttrs(attrs)

	expected := []xml.Attr{
		attr("xmlns", "android", androidNS),
		attr("xmlns", "app", appNS),
		attr("xmlns", "tools", toolsNS),
		attr("xmlns", "unknown", "http://schemas.android.com/apk/custom"),
	}
	if str(sorted) != str(expected) {
		t.Errorf("got\n%s\n, want\n%s", str(sorted), str(expected))
	}
}

func TestAndroidAttrs(t *testing.T) {
	attrs := []xml.Attr{
		attr(androidNS, "layout_weight", "1"),
		attr(androidNS, "layout_width", "wrap_content"),
		attr(androidNS, "gravity", "center"),
		attr(androidNS, "id", "@+id/open"),
		attr(androidNS, "layout_height", "match_parent"),
	}

	sorted := sortAttrs(attrs)

	expected := []xml.Attr{
		attr(androidNS, "id", "@+id/open"),
		attr(androidNS, "layout_width", "wrap_content"),
		attr(androidNS, "layout_height", "match_parent"),
		attr(androidNS, "gravity", "center"),
		attr(androidNS, "layout_weight", "1"),
	}
	if str(sorted) != str(expected) {
		t.Errorf("got\n%s\n, want\n%s", str(sorted), str(expected))
	}
}

func TestAppAttrs(t *testing.T) {
	attrs := []xml.Attr{
		attr(appNS, "layout_constraintTop_toTopOf", "parent"),
		attr(appNS, "layoutManager", "androidx.recyclerview.widget.LinearLayoutManager"),
		attr(appNS, "layout_constraintBottom_toBottomOf", "@id/title"),
	}

	sorted := sortAttrs(attrs)

	expected := []xml.Attr{
		attr(appNS, "layoutManager", "androidx.recyclerview.widget.LinearLayoutManager"),
		attr(appNS, "layout_constraintBottom_toBottomOf", "@id/title"),
		attr(appNS, "layout_constraintTop_toTopOf", "parent"),
	}
	if str(sorted) != str(expected) {
		t.Errorf("got\n%s\n, want\n%s", str(sorted), str(expected))
	}
}

func TestNamespaceComparison(t *testing.T) {
	attrs := []xml.Attr{
		attr("xmlns", "app", appNS),
		attr(androidNS, "id", "@+id/open"),
		attr(androidNS, "layout_width", "wrap_content"),
		attr("", "style", "@style/list"),
		attr("xmlns", "tools", toolsNS),
		attr(toolsNS, "visibility", "gone"),
		attr(androidNS, "layout_height", "match_parent"),
		attr("xmlns", "android", androidNS),
		attr(appNS, "layoutManager", "androidx.recyclerview.widget.LinearLayoutManager"),
	}

	sorted := sortAttrs(attrs)

	expected := []xml.Attr{
		attr("xmlns", "android", androidNS),
		attr("xmlns", "app", appNS),
		attr("xmlns", "tools", toolsNS),
		attr(androidNS, "id", "@+id/open"),
		attr(androidNS, "layout_width", "wrap_content"),
		attr(androidNS, "layout_height", "match_parent"),
		attr(appNS, "layoutManager", "androidx.recyclerview.widget.LinearLayoutManager"),
		attr(toolsNS, "visibility", "gone"),
		attr("", "style", "@style/list"),
	}
	if str(sorted) != str(expected) {
		t.Errorf("got\n%s\n, want\n%s", str(sorted), str(expected))
	}
}

func attr(space, local, value string) xml.Attr {
	return xml.Attr{Name: tagName(space, local), Value: value}
}

func tagName(space, local string) xml.Name {
	return xml.Name{Space: space, Local: local}
}

func str(attrs []xml.Attr) string {
	strs := []string{}

	for _, a := range attrs {
		strs = append(strs, fmt.Sprintf(`%s:%s="%s"`, a.Name.Space, a.Name.Local, a.Value))
	}

	return strings.Join(strs, "\n")
}
