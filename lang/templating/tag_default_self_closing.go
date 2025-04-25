package templating

import (
	"fmt"

	"github.com/bbfh-dev/mend/lang/attrs"
	"github.com/bbfh-dev/mend/lang/printer"
)

type SelfClosingTag struct {
	*BaseTag
	Name  string
	Attrs *attrs.Attributes
}

func NewSelfClosing(name string, attrs *attrs.Attributes) *SelfClosingTag {
	return &SelfClosingTag{
		BaseTag: NewBase(),
		Name:    name,
		Attrs:   attrs,
	}
}

func (tag *SelfClosingTag) Render(writer printer.Writer, indent int) {
	tag.BaseTag.Render(writer, indent)

	fmt.Fprintf(writer, "<%s", tag.Name)
	tag.Attrs.Render(writer)
	writer.WriteString(" />")
}
