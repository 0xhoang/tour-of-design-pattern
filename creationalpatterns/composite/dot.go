package composite

import "fmt"

type Dot struct {
	x, y int
}

func NewDot(x, y int) *Dot {
	return &Dot{x, y}
}

func (d *Dot) move(x, y int) {
	d.x += x
	d.y += y
}

func (d *Dot) draw() {
	fmt.Printf("Drawing a dot at (%d,%d)\n", d.x, d.y)
}
