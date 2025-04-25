package templating

import (
	"fmt"

	"github.com/bbfh-dev/mend/lang/attrs"
	"github.com/bbfh-dev/mend/lang/printer"
)

type DefaultTag struct {
	*BasePairedTag
	Name  string
	Attrs *attrs.Attributes
}

func NewDefault(name string, attrs *attrs.Attributes) *DefaultTag {
	return &DefaultTag{
		BasePairedTag: NewPairedBase(),
		Name:          name,
		Attrs:         attrs,
	}
}

func (tag *DefaultTag) Render(writer printer.Writer, indent int) {
	tag.BaseTag.Render(writer, indent)

	fmt.Fprintf(writer, "<%s", tag.Name)
	tag.Attrs.Render(writer)
	writer.WriteString(">\n")

	tag.BasePairedTag.Render(writer, indent)

	tag.BaseTag.Render(writer, indent)
	fmt.Fprintf(writer, "</%s>", tag.Name)
}
