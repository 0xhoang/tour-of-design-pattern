package composite

type CompoundGraphic struct {
	children []InterfaceGraphic
}

func NewCompoundGraphic() *CompoundGraphic {
	return &CompoundGraphic{}
}

func (cg *CompoundGraphic) add(child InterfaceGraphic) {
	cg.children = append(cg.children, child)
}

func (cg *CompoundGraphic) remove(child InterfaceGraphic) {
	for i, i2 := range cg.children {
		if i2 == child {
			cg.children = append(cg.children[:i], cg.children[i+1:]...)
		}
	}
}

func (cg *CompoundGraphic) move(x, y int) {
	for _, i2 := range cg.children {
		i2.move(x, y)
	}
}

func (cg *CompoundGraphic) draw() {
	for _, i2 := range cg.children {
		i2.draw()
	}
}
