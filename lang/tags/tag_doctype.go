package tags

import (
	"fmt"

	"github.com/bbfh-dev/mend/lang/printer"
)

type DoctypeTag struct {
	*BaseTag
	Doctype string
}

func NewDoctype(doctype string) *DoctypeTag {
	return &DoctypeTag{
		BaseTag: NewBase(),
		Doctype: doctype,
	}
}

func (tag *DoctypeTag) Render(writer printer.Writer, indent int) {
	tag.BaseTag.Render(writer, indent)
	fmt.Fprintf(writer, "<!DOCTYPE %s>", tag.Doctype)
}
