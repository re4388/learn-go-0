package learn

import "fmt"

/*
In Go, interfaces are used to define a set of methods that a type must implement.
*/

// Define an interface named Shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a struct named Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the Area method for the Rectangle struct
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implement the Perimeter method for the Rectangle struct
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Define a struct named Circle that implements the Shape interface
type Circle1 struct {
	Radius float64
}

// Implement the Area method for the Circle struct
func (c Circle1) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Implement the Perimeter method for the Circle1 struct
func (c Circle1) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func RUN_interface1() {
	// Create a slice of Shape interface that can hold different shapes
	// you can setup a shape, which put the struct that impl the shape interface
	// these struct need to impl the shape
	shapes := []Shape{
		Rectangle{Width: 4, Height: 5},
		Circle1{Radius: 3},
	}

	// Iterate over the shapes and print their area and perimeter
	for _, shape := range shapes {
		fmt.Println("Area:", shape.Area())
		fmt.Println("Perimeter:", shape.Perimeter())
		fmt.Println()
	}
}
