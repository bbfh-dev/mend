package tags

import (
	"github.com/bbfh-dev/mend.html/mend/attrs"
	"github.com/bbfh-dev/mend.html/mend/settings"
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
