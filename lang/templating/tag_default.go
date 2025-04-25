package templating

import (
	"fmt"

	"github.com/bbfh-dev/mend/lang/attrs"
	"github.com/bbfh-dev/mend/lang/printer"
)

type DefaultTag struct {
	*BasePairedTag
	BaseDefaultTag
}

func NewDefault(name string, attrs *attrs.Attributes) *DefaultTag {
	return &DefaultTag{
		BasePairedTag: NewPairedBase(),
		BaseDefaultTag: BaseDefaultTag{
			Name:  name,
			Attrs: attrs,
		},
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
