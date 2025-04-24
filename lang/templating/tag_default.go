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

func NewDefault(indent int, name string, attrs *attrs.Attributes) *DefaultTag {
	return &DefaultTag{
		BasePairedTag: NewPairedBase(indent),
		BaseDefaultTag: BaseDefaultTag{
			Name:  name,
			Attrs: attrs,
		},
	}
}

func (tag *DefaultTag) Render(writer printer.Writer) {
	tag.BaseTag.Render(writer)

	fmt.Fprintf(writer, "<%s", tag.Name)
	tag.Attrs.Render(writer)
	writer.WriteString(">\n")

	tag.BasePairedTag.Render(writer)

	tag.BaseTag.Render(writer)
	fmt.Fprintf(writer, "</%s>", tag.Name)
}
