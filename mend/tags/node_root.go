package tags

// Represents a simple paired node
type RootNode struct {
	*pairedNode
}

func NewRootNode() *RootNode {
	return &RootNode{
		pairedNode: newPairedNode(),
	}
}

func (node *RootNode) Render(out writer, indent int) {
	node.renderMinimal(out, indent)
}

func (node *RootNode) Visible() bool {
	return true
}
