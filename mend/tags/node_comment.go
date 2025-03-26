package tags

import (
	"fmt"

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
