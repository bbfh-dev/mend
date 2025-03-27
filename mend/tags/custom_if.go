package tags

import (
	"errors"
	"strings"
)

type CustomIfNode struct {
	*pairedNode
	Value  string
	Expect bool
}

func NewCustomIfNode(value string, expect bool) *CustomIfNode {
	return &CustomIfNode{
		pairedNode: newPairedNode(),
		Value:      value,
		Expect:     expect,
	}
}

func (node *CustomIfNode) check() bool {
	return (node.Value == "true") == node.Expect
}

func (node *CustomIfNode) Render(out writer, indent int) {
	if node.check() {
		node.renderMinimal(out, indent)
	}
}

func (node *CustomIfNode) Visible() bool {
	return node.check()
}

func (node *CustomIfNode) Clone() Node {
	clone := *node
	clone.pairedNode = clone.pairedNode.Clone()
	return &clone
}

func (node *CustomIfNode) ParseExpressions(source string, fn expressionFunc) (err error) {
	node.Value, err = fn(source, node.Value)
	return errors.Join(err, node.pairedNode.ParseExpressions(source, fn))
}

func (node *CustomIfNode) ReplaceText(text string, with string) {
	node.Value = strings.ReplaceAll(node.Value, text, with)
	node.pairedNode.ReplaceText(text, with)
}
