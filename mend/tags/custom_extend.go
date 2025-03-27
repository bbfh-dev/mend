package tags

import (
	"errors"
	"fmt"
)

type CustomExtendNode struct {
	*pairedNode
	// Store all the nodes from the file it extends
	Inner *pairedNode
	// The node where contents will be placed
	Slot NodeWithChildren
}

func NewCustomExtendNode() *CustomExtendNode {
	return &CustomExtendNode{
		pairedNode: newPairedNode(),
		Inner:      newPairedNode(),
	}
}

func (node *CustomExtendNode) Render(out writer, indent int) {
	node.Inner.renderMinimal(out, indent)
}

func (node *CustomExtendNode) Visible() bool {
	return true
}

func (node *CustomExtendNode) Clone() Node {
	clone := *node
	clone.pairedNode = clone.pairedNode.Clone()
	return &clone
}

func (node *CustomExtendNode) ParseExpressions(source string, fn expressionFunc) (err error) {
	errs := make([]error, 0, len(node.Inner.Children))
	for _, child := range node.Inner.Children {
		err = child.ParseExpressions(source, fn)
		if err != nil {
			errs = append(errs, fmt.Errorf("<extend>: %w", err))
		}
	}
	return errors.Join(errs...)
}
