package templating

import (
	"fmt"

	"github.com/bbfh-dev/mend/lang/printer"
)

type CommentTag struct {
	*BaseTag
	Comment string
}

func NewComment(indent int, comment string) *CommentTag {
	return &CommentTag{
		BaseTag: NewBase(indent),
		Comment: comment,
	}
}

func (tag *CommentTag) Render(writer printer.Writer) {
	tag.BaseTag.Render(writer)
	fmt.Fprintf(writer, "<!-- %s -->", tag.Comment)
}
