package format

import (
	"encoding/xml"
	"sort"
)

var priorities map[string]int

func init() {
	priorities = attrPriorities()
}

type byPriority []xml.Attr

func (s byPriority) Len() int {
	return len(s)
}
func (s byPriority) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byPriority) Less(i, j int) bool {
	fst := s[i]
	snd := s[j]

	fstP, fstContains := priorities[fst.Name.Local]
	sndP, sndContains := priorities[snd.Name.Local]

	if fstContains && sndContains {
		return fstP < sndP
	}

	// Fallback to lexicographical order
	return fst.Value < snd.Value
}

func sortAttrs(attrs []xml.Attr) []xml.Attr {
	sorted := make([]xml.Attr, len(attrs))
	copy(sorted, attrs)

	sort.Sort(byPriority(sorted))

	return sorted
}
