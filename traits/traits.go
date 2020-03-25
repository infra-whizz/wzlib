package wzlib_traits

// WzTraits struct
type WzTraits struct {
	container      *WzTraitsContainer
	traitsFileName string
}

// NewTraits constructor.
/*
  WzTraits is a compount object that is loading all possible traits,
  registered to it and returns a self-contained instance.
*/
func NewTraits(fpath string) *WzTraits {
	t := new(WzTraits)
	t.traitsFileName = fpath
	t.container = NewWzTraitsContainer().LoadFromFile(t.traitsFileName)
	return t
}

// GetContainer returns a traits container
func (tl *WzTraits) GetContainer() *WzTraitsContainer {
	return tl.container
}

// LoadAttribute to the WzTraits container
func (tl *WzTraits) LoadAttribute(attr TraitsAttribute) {
	attr.Load(tl.container)
}

// Save traits data to a file
func (tl *WzTraits) Save() {
	tl.container.SaveToFile(tl.traitsFileName)
}

// Load traits data from a file
func (tl *WzTraits) Load() {
	tl.container.LoadFromFile(tl.traitsFileName)
}
