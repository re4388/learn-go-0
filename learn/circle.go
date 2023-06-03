package learn

import (
	"fmt"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(%0.2f)", c.Radius)
}

func PrintCircle() {
	c := Circle{Radius: 5.0}
	fmt.Printf("Area of %s: %0.2f\n", c, c.Area())
}
