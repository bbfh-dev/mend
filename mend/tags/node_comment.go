package tags

import (
	"fmt"
	"strings"

	"github.com/bbfh-dev/mend.html/mend/settings"
)

// Represents a block of text
type CommentNode struct {
	Comment string
}

func NewCommentNode(comment string) *CommentNode {
	return &CommentNode{
		Comment: comment,
	}
}

func (node *CommentNode) Render(out writer, indent int) {
	settings.WriteIndent(out, indent)
	fmt.Fprintf(out, "<!-- %s -->", node.Comment)
}

func (node *CommentNode) Visible() bool {
	return true
}

func (node *CommentNode) ReplaceText(text string, with string) {
	node.Comment = strings.ReplaceAll(node.Comment, text, with)
}

func (node *CommentNode) Clone() Node {
	clone := *node
	return &clone
}
