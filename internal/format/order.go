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

func attrPriorities() map[string]int {
	priority := make(map[string]int)

	for i, a := range attrs() {
		priority[a] = i
	}

	return priority
}

func attrs() []string {
	return []string{
		// namespaces
		"android",
		"app",
		"tools",

		"style",

		"id",
		"layout_width",
		"layout_height",
		"layout_weight",

		"paddingTop",
		"clipToPadding",

		"orientation",

		"text",
		"textAppearance",
		"textColor",
		"textSize",

		"navigationContentDescription",
	}
}