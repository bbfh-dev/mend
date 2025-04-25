package tags

import (
	"fmt"

	"github.com/bbfh-dev/mend/lang/attrs"
	"github.com/bbfh-dev/mend/lang/printer"
)

type DefaultRootTag struct {
	*DefaultTag
}

func NewDefaultRoot(name string, attrs *attrs.Attributes) *DefaultRootTag {
	return &DefaultRootTag{
		DefaultTag: NewDefault(name, attrs),
	}
}

func (tag *DefaultRootTag) Render(writer printer.Writer, indent int) {
	tag.BaseTag.Render(writer, indent)

	fmt.Fprintf(writer, "<%s", tag.Name)
	tag.Attrs.Render(writer)
	writer.WriteString(">\n")

	writer.WriteString("\n")
	for _, child := range tag.Children {
		if child.Visibility() != INVISIBLE {
			child.Render(writer, indent)
			writer.WriteString("\n\n")
		}
	}

	tag.BaseTag.Render(writer, indent)
	fmt.Fprintf(writer, "</%s>", tag.Name)
}
