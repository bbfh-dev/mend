package tags

import (
	"github.com/bbfh-dev/mend/mend/attrs"
	"github.com/bbfh-dev/mend/mend/settings"
)

// Represents a self-closing HTML tag
type VoidNode struct {
	Void       string
	Attributes attrs.Attributes
}

func NewVoidNode(tag string, attrs attrs.Attributes) *VoidNode {
	return &VoidNode{
		Void:       tag,
		Attributes: attrs,
	}
}

func (node *VoidNode) Render(out writer, indent int) {
	settings.WriteIndent(out, indent)
	renderOpeningTag(out, node.Void, node.Attributes)
	out.WriteString(" />")
}

func (node *VoidNode) Visible() bool {
	return true
}

func (node *VoidNode) ParseExpressions(source string, fn expressionFunc) (err error) {
	node.Attributes, err = node.Attributes.ParseExpressions(source, fn)
	return err
}

func (node *VoidNode) ReplaceText(text string, with string) {
	node.Attributes = node.Attributes.ReplaceText(text, with)
}

func (node *VoidNode) Clone() Node {
	clone := *node
	return &clone
}

func (node *VoidNode) MergeAttributes(attrs attrs.Attributes) bool {
	node.Attributes = node.Attributes.Merge(attrs)
	return true
}
