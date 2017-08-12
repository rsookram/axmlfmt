/*
Copyright 2017 Rashad Sookram

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
