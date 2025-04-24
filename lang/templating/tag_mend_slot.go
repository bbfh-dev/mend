package templating

type MendSlotTag struct {
	*BasePairedTag
}

func NewMendSlot(indent int) *MendSlotTag {
	return &MendSlotTag{
		BasePairedTag: NewPairedBase(indent),
	}
}

func (tag *MendSlotTag) Visibility() visibility {
	return INLINE
}
