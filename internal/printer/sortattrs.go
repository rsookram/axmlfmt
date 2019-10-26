package printer

import (
	"encoding/xml"
	"sort"
)

var nsPriority = []string{
	"xmlns",
	"http://schemas.android.com/apk/res/android",
	"http://schemas.android.com/apk/res-auto",
	"http://schemas.android.com/tools",
	"", // for attributes without a namespace like style or layout
}

var androidPriority = []string{
	"id",
	"layout_width",
	"layout_height",
}

// sortAttrs returns a sorted slice of the given attributes in the following
// order:
//
//   - xmlns:android
//   - xmlns:* (alphabetic)
//   - android:id
//   - android:layout_width
//   - android:layout_height
//   - android:* (alphabetic)
//   - app:* (alphabetic)
//   - tools:* (alphabetic)
//   - :* (alphabetic)
func sortAttrs(attrs []xml.Attr) []xml.Attr {
	sorted := make([]xml.Attr, len(attrs))
	copy(sorted, attrs)

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i].Name, sorted[j].Name)
	})

	return sorted
}
func less(fst, snd xml.Name) bool {
	if fst.Space != snd.Space {
		fstP, hasFst := indexOf(nsPriority, fst.Space)
		sndP, hasSnd := indexOf(nsPriority, snd.Space)

		if hasFst && hasSnd {
			return fstP < sndP
		}
		if hasFst {
			return true
		}
		if hasSnd {
			return false
		}
	}

	ns := fst.Space
	switch ns {
	case "xmlns":
		// "android" is always first
		if fst.Local == "android" && snd.Local != "android" {
			return true
		}
		return fst.Local < snd.Local
	case "http://schemas.android.com/apk/res/android":
		fstP, hasFst := indexOf(androidPriority, fst.Local)
		sndP, hasSnd := indexOf(androidPriority, snd.Local)

		if hasFst && hasSnd {
			return fstP < sndP
		}
		if hasFst {
			return true
		}
		if hasSnd {
			return false
		}
	}
	return fst.Local < snd.Local
}

func indexOf(haystack []string, needle string) (int, bool) {
	for i, s := range haystack {
		if s == needle {
			return i, true
		}
	}

	return -1, false
}
