package templating

import (
	"github.com/bbfh-dev/mend/lang/printer"
)

type visibility int

const (
	INVISIBLE visibility = iota
	VISIBLE
	INLINE
)

type Tag interface {
	Render(writer printer.Writer, indent int)
	Visibility() visibility
	Clone() Tag
}

type PairedTag interface {
	Tag
	SetChildren(tags []Tag)
	Append(tags ...Tag)
}
