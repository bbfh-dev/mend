package attrs

import "fmt"

func (attrs Attributes) Render(output writer) {
	last := len(attrs.order) - 1

	for i, key := range attrs.order {
		attrs.renderKey(output, key)
		if i != last {
			output.WriteString(" ")
		}
	}
}

func (attrs Attributes) renderKey(output writer, key string) {
	if len(attrs.values[key]) == 0 {
		output.WriteString(key)
		return
	}

	fmt.Fprintf(output, "%s=%q", key, attrs.values[key])
}
