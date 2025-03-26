package tags

import (
	"github.com/bbfh-dev/mend.html/mend/attrs"
	"github.com/bbfh-dev/mend.html/mend/settings"
)

// Represents a regular HTML paired node
type TagNode struct {
	*pairedNode
	Tag        string
	Attributes attrs.Attributes
}

func NewTagNode(tag string, attributes attrs.Attributes) *TagNode {
	return &TagNode{
		Tag:        tag,
		Attributes: attributes,
		pairedNode: newPairedNode(),
	}
}

func (node *TagNode) Render(out writer, indent int) {
	settings.WriteIndent(out, indent)
	renderOpeningTag(out, node.Tag, node.Attributes)
	out.WriteString(">\n")

	if node.Tag == "html" {
		node.renderPadded(out, indent)
	} else {
		node.renderList(out, indent)
	}

	settings.WriteIndent(out, indent)
	renderClosingTag(out, node.Tag)
}
