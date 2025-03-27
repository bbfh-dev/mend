package tags

import (
	"errors"

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

func (node *TagNode) Visible() bool {
	return true
}

func (node *TagNode) ParseExpressions(source string, fn expressionFunc) (err error) {
	node.Attributes, err = node.Attributes.ParseExpressions(source, fn)
	return errors.Join(err, node.pairedNode.ParseExpressions(source, fn))
}

func (node *TagNode) ReplaceText(text string, with string) {
	node.Attributes = node.Attributes.ReplaceText(text, with)
	node.pairedNode.ReplaceText(text, with)
}

func (node *TagNode) Clone() Node {
	clone := *node
	clone.pairedNode = clone.pairedNode.Clone()
	return &clone
}
