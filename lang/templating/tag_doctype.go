package templating

import (
	"fmt"

	"github.com/bbfh-dev/mend/lang/printer"
)

type DoctypeTag struct {
	*BaseTag
	Doctype string
}

func NewDoctype(indent int, doctype string) *DoctypeTag {
	return &DoctypeTag{
		BaseTag: NewBase(indent),
		Doctype: doctype,
	}
}

func (tag *DoctypeTag) Render(writer printer.Writer) {
	tag.BaseTag.Render(writer)
	fmt.Fprintf(writer, "<!DOCTYPE %s>", tag.Doctype)
}
