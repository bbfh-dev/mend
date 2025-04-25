package tags

type MendVoidTag struct {
	*BasePairedTag
}

func NewMendVoid() *MendVoidTag {
	return &MendVoidTag{
		BasePairedTag: NewPairedBase(),
	}
}

func (tag *MendVoidTag) Visibility() visibility {
	return INVISIBLE
}
