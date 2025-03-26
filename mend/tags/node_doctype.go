package tags

import (
	"fmt"

	"github.com/bbfh-dev/mend.html/mend/settings"
)

type DoctypeNode struct {
	Doctype string
}

func NewDoctypeNode(doctype string) *DoctypeNode {
	return &DoctypeNode{
		Doctype: doctype,
	}
}

func (node *DoctypeNode) Render(out writer, indent int) {
	settings.WriteIndent(out, indent)
	fmt.Fprintf(out, "<!DOCTYPE %s>", node.Doctype)
}

func (node *DoctypeNode) Visible() bool {
	return true
}
