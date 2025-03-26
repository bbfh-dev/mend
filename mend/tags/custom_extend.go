package tags

// Represents a regular HTML paired node
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
