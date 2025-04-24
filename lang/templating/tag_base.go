package templating

import "github.com/bbfh-dev/mend/lang/printer"

type BaseTag struct {
	indent int
}

func NewBase(indent int) *BaseTag {
	return &BaseTag{
		indent: indent,
	}
}

func (tag *BaseTag) Render(writer printer.Writer) {
	printer.WriteIndent(writer, tag.indent)
}

func (tag *BaseTag) Visibility() visibility {
	return VISIBLE
}

func (tag *BaseTag) Indent() int {
	return tag.indent
}

func (tag *BaseTag) Shift(offset int) {
	tag.indent += offset
}

func (tag *BaseTag) Clone() Tag {
	clone := *tag
	return &clone
}
