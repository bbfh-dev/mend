package templating

import (
	"fmt"

	"github.com/bbfh-dev/mend/lang/attrs"
	"github.com/bbfh-dev/mend/lang/printer"
)

type SelfClosingTag struct {
	*BaseTag
	BaseDefaultTag
}

func NewSelfClosing(indent int, name string, attrs *attrs.Attributes) *SelfClosingTag {
	return &SelfClosingTag{
		BaseTag: NewBase(indent),
		BaseDefaultTag: BaseDefaultTag{
			Name:  name,
			Attrs: attrs,
		},
	}
}

func (tag *SelfClosingTag) Render(writer printer.Writer) {
	tag.BaseTag.Render(writer)

	fmt.Fprintf(writer, "<%s", tag.Name)
	tag.Attrs.Render(writer)
	writer.WriteString(" />")
}
