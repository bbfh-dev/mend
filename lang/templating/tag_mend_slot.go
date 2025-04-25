package templating

type MendSlotTag struct {
	*BasePairedTag
}

func NewMendSlot() *MendSlotTag {
	return &MendSlotTag{
		BasePairedTag: NewPairedBase(),
	}
}

func (tag *MendSlotTag) Visibility() visibility {
	return INLINE
}
