package tags

import (
	"fmt"

	"github.com/bbfh-dev/mend.html/mend/attrs"
)

type Node interface {
	Render(out writer, indent int)
}

type NodeWithChildren interface {
	Node
	Add(...Node)
}

// NOTE: Doesn't print the '>' in the end to allow for it to be used for both block and void tags
func renderOpeningTag(out writer, tag string, attributes attrs.Attributes) {
	if attributes.IsEmpty() {
		fmt.Fprintf(out, "<%s", tag)
		return
	}

	fmt.Fprintf(out, "<%s ", tag)
	attributes.Render(out)
}

func renderClosingTag(out writer, tag string) {
	fmt.Fprintf(out, "</%s>", tag)
}

// Shared fields/methods for paired nodes
//
// A paired node is: <tag ...></tag>
type pairedNode struct {
	Children []Node
}

func newPairedNode() *pairedNode {
	return &pairedNode{
		Children: []Node{},
	}
}

func (node *pairedNode) renderMinimal(out writer, indent int) {
	last := len(node.Children) - 1
	for i, child := range node.Children {
		child.Render(out, indent)
		if i != last {
			out.WriteString("\n")
		}
	}
}

func (node *pairedNode) renderList(out writer, indent int) {
	for _, child := range node.Children {
		child.Render(out, indent+1)
		out.WriteString("\n")
	}
}

func (node *pairedNode) renderPadded(out writer, indent int) {
	out.WriteString("\n")
	for _, child := range node.Children {
		child.Render(out, indent)
		out.WriteString("\n\n")
	}
}

func (node *pairedNode) Add(nodes ...Node) {
	node.Children = append(node.Children, nodes...)
}
