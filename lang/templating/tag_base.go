package templating

import (
	"github.com/bbfh-dev/mend/lang/printer"
)

type BaseTag struct {
}

func NewBase() *BaseTag {
	return &BaseTag{}
}

func (tag *BaseTag) Render(writer printer.Writer, indent int) {
	printer.WriteIndent(writer, indent)
}

func (tag *BaseTag) Visibility() visibility {
	return VISIBLE
}

func (tag *BaseTag) Clone() Tag {
	clone := *tag
	return &clone
}

func (tag *BaseTag) OverrideAttr(key string, value string) bool {
	return false
}
