package templating

import (
	"fmt"

	"github.com/bbfh-dev/mend/lang/attrs"
	"github.com/bbfh-dev/mend/lang/printer"
)

type DefaultRootTag struct {
	*DefaultTag
}

func NewDefaultRoot(indent int, name string, attrs *attrs.Attributes) *DefaultRootTag {
	return &DefaultRootTag{
		DefaultTag: NewDefault(indent, name, attrs),
	}
}

func (tag *DefaultRootTag) Render(writer printer.Writer) {
	tag.BaseTag.Render(writer)

	fmt.Fprintf(writer, "<%s", tag.Name)
	tag.Attrs.Render(writer)
	writer.WriteString(">\n")

	writer.WriteString("\n")
	for _, child := range tag.Children {
		if child.Visibility() != INVISIBLE {
			child.Shift(-1)
			child.Render(writer)
			writer.WriteString("\n\n")
		}
	}

	tag.BaseTag.Render(writer)
	fmt.Fprintf(writer, "</%s>", tag.Name)
}
