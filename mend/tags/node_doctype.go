package tags

import (
	"fmt"

	"github.com/bbfh-dev/mend/mend/settings"
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

func (node *DoctypeNode) ParseExpressions(source string, fn expressionFunc) (err error) {
	return nil
}

func (node *DoctypeNode) ReplaceText(text string, with string) {
}

func (node *DoctypeNode) Clone() Node {
	clone := *node
	return &clone
}
