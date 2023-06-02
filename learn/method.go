package learn

import (
	"fmt"
	"math"
)

/*
multi line
comment*/

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ptr receiver, so we modify the ref one
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func printScaledABS() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

///

// You can declare a method on non-struct types, too.
type MyFloat float64

// You can only declare a method with a receiver whose type is defined in the same package as the method.
// You cannot declare a method with a receiver whose type is defined in another package
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func PrintMyFloat() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
