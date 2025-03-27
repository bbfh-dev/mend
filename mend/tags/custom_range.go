package tags

import (
	"github.com/tidwall/gjson"
)

type CustomRangeNode struct {
	*pairedNode
	Name   string
	Values gjson.Result
}

func NewCustomRangeNode(name string, values gjson.Result) *CustomRangeNode {
	return &CustomRangeNode{
		pairedNode: newPairedNode(),
		Name:       name,
		Values:     values,
	}
}

func (node *CustomRangeNode) Render(out writer, indent int) {}

func (node *CustomRangeNode) Visible() bool {
	return false
}

func (node *CustomRangeNode) Clone() Node {
	clone := *node
	clone.pairedNode = clone.pairedNode.Clone()
	return &clone
}
