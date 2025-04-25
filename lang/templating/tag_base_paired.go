package templating

import (
	"github.com/bbfh-dev/mend/lang/printer"
)

type BasePairedTag struct {
	*BaseTag
	Children []Tag
}

func NewPairedBase() *BasePairedTag {
	return &BasePairedTag{
		BaseTag:  NewBase(),
		Children: []Tag{},
	}
}

func (tag *BasePairedTag) Render(writer printer.Writer, indent int) {
	for _, child := range tag.Children {
		// fmt.Fprintf(writer, "<!-- %s %+v -->\n", reflect.TypeOf(child), child)
		switch child.Visibility() {
		case VISIBLE:
			child.Render(writer, indent+1)
			writer.WriteString("\n")
		case INLINE:
			child.Render(writer, indent)
		}
	}
}

func (tag *BasePairedTag) SetChildren(tags []Tag) {
	tag.Children = tags
}

func (tag *BasePairedTag) Append(tags ...Tag) {
	tag.Children = append(tag.Children, tags...)
}
