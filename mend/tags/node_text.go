package tags

import (
	"strings"

	"github.com/bbfh-dev/mend/mend/settings"
)

// Represents a block of text
type TextNode struct {
	Text string
}

func NewTextNode(text string) *TextNode {
	return &TextNode{
		Text: text,
	}
}

func (node *TextNode) Render(out writer, indent int) {
	settings.WriteIndent(out, indent)
	out.WriteString(node.Text)
}

func (node *TextNode) Visible() bool {
	return true
}

func (node *TextNode) ParseExpressions(source string, fn expressionFunc) (err error) {
	node.Text, err = fn(source, node.Text)
	return err
}

func (node *TextNode) ReplaceText(text string, with string) {
	node.Text = strings.ReplaceAll(node.Text, text, with)
}

func (node *TextNode) Clone() Node {
	clone := *node
	return &clone
}
