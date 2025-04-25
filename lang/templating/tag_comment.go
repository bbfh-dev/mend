package templating

import (
	"fmt"

	"github.com/bbfh-dev/mend/lang/printer"
)

type CommentTag struct {
	*BaseTag
	Comment string
}

func NewComment(comment string) *CommentTag {
	return &CommentTag{
		BaseTag: NewBase(),
		Comment: comment,
	}
}

func (tag *CommentTag) Render(writer printer.Writer, indent int) {
	tag.BaseTag.Render(writer, indent)
	fmt.Fprintf(writer, "<!-- %s -->", tag.Comment)
}
