package templating

import (
	"github.com/bbfh-dev/mend/lang/printer"
)

type MendExtendTag struct {
	*BasePairedTag
	Root PairedTag
	Slot *MendSlotTag
}

func NewMendExtend(root PairedTag, slot *MendSlotTag) *MendExtendTag {
	return &MendExtendTag{
		BasePairedTag: NewPairedBase(),
		Root:          root,
		Slot:          slot,
	}
}

func (tag *MendExtendTag) Render(writer printer.Writer, indent int) {
	tag.Root.Render(writer, indent)
}

func (tag *MendExtendTag) Visibility() visibility {
	return INLINE
}
