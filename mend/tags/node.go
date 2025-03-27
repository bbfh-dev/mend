package tags

import (
	"errors"
	"fmt"

	"github.com/bbfh-dev/mend.html/mend/attrs"
)

type expressionFunc func(source, text string) (string, error)

type Node interface {
	Render(out writer, indent int)
	Visible() bool
	ParseExpressions(string, expressionFunc) error
	ReplaceText(text string, with string)
	Clone() Node
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
		if !child.Visible() {
			continue
		}
		child.Render(out, indent)
		if i != last {
			out.WriteString("\n")
		}
	}
}

func (node *pairedNode) renderList(out writer, indent int) {
	for _, child := range node.Children {
		if !child.Visible() {
			continue
		}
		child.Render(out, indent+1)
		out.WriteString("\n")
	}
}

func (node *pairedNode) renderPadded(out writer, indent int) {
	out.WriteString("\n")
	for _, child := range node.Children {
		if !child.Visible() {
			continue
		}
		child.Render(out, indent)
		out.WriteString("\n\n")
	}
}

func (node *pairedNode) Add(nodes ...Node) {
	node.Children = append(node.Children, nodes...)
}

func (node *pairedNode) ParseExpressions(source string, fn expressionFunc) (err error) {
	errs := make([]error, 0, len(node.Children))
	for _, child := range node.Children {
		err = child.ParseExpressions(source, fn)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func (node *pairedNode) ReplaceText(text string, with string) {
	for _, child := range node.Children {
		child.ReplaceText(text, with)
	}
}

func (node *pairedNode) Clone() *pairedNode {
	clone := *node
	children := make([]Node, len(clone.Children))
	for i, child := range clone.Children {
		children[i] = child.Clone()
	}
	clone.Children = children
	return &clone
}
