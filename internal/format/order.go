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
