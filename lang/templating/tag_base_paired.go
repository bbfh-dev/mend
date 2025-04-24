package templating

import (
	"github.com/bbfh-dev/mend/lang/printer"
)

type BasePairedTag struct {
	*BaseTag
	Children []Tag
}

func NewPairedBase(indent int) *BasePairedTag {
	return &BasePairedTag{
		BaseTag:  NewBase(indent),
		Children: []Tag{},
	}
}

func (tag *BasePairedTag) Render(writer printer.Writer) {
	for _, child := range tag.Children {
		// fmt.Fprintf(writer, "<!-- %s %+v -->\n", reflect.TypeOf(child), child)
		switch child.Visibility() {
		case VISIBLE:
			child.Render(writer)
			writer.WriteString("\n")
		case INLINE:
			child.Render(writer)
		}
	}
}

func (tag *BasePairedTag) Shift(offset int) {
	tag.BaseTag.Shift(offset)
	for _, child := range tag.Children {
		child.Shift(offset)
	}
}

func (tag *BasePairedTag) SetChildren(tags []Tag) {
	tag.Children = tags
}

func (tag *BasePairedTag) Append(tags ...Tag) {
	tag.Children = append(tag.Children, tags...)
}
