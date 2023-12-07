package composite

import "fmt"

type Circle struct {
	Dot
	radius int
}

func NewCircle(x, y, radius int) *Circle {
	return &Circle{radius: radius, Dot: Dot{x: x, y: y}}
}

func (c *Circle) draw() {
	// Move circle
	fmt.Printf("Drawing a circle at (%d,%d) with radius %d\n", c.x, c.y, c.radius)
}
