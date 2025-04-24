package templating

import "github.com/bbfh-dev/mend/lang/printer"

type MendExtendTag struct {
	*BasePairedTag
	Root PairedTag
	Slot *MendSlotTag
}

func NewExtendSlot(indent int, root PairedTag, slot *MendSlotTag) *MendExtendTag {
	return &MendExtendTag{
		BasePairedTag: NewPairedBase(indent),
		Root:          root,
		Slot:          slot,
	}
}

func (tag *MendExtendTag) Render(writer printer.Writer) {
	tag.Root.Render(writer)
}

func (tag *MendExtendTag) Visibility() visibility {
	return INLINE
}
