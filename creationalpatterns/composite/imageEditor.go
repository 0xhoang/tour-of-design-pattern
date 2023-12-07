package composite

type ImageEditor struct {
	all CompoundGraphic
}

func NewImageEditor() *ImageEditor {
	return &ImageEditor{}
}

func (ie *ImageEditor) load() {
	all := NewCompoundGraphic()
	all.add(NewDot(1, 2))
	all.add(NewCircle(5, 3, 10))
	ie.all = *all
}

func (ie *ImageEditor) groupSelected(components []InterfaceGraphic) {
	group := NewCompoundGraphic()
	for _, i2 := range components {
		group.add(i2)
		ie.all.remove(i2)
	}

	ie.all.add(group)
	ie.all.draw()
}
