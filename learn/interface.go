package learn

import "fmt"

// Define Stringer as an interface type with one method, String.
type Stringer interface {
	String() string
}

// Define pair as a struct with two fields, ints named x and y.
type pair2 struct {
	x, y int
}

// you can define a method on struct pair.
// Pair now implements Stringer because Pair has defined all the methods in the interface.
func (p pair2) String() string { // p is called the "receiver"
	// Sprintf is another public function in package fmt.
	// Dot syntax references fields of p.
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func LearnInterfaces() {
	// Brace syntax is a "struct literal".
	// It evaluates to an initialized struct.
	// The := syntax declares and initializes p1 to this struct.
	p1 := pair2{3, 4}

	// Call String method of p, of type pair.
	fmt.Println(p1.String())

	// Declare stringer of interface type Stringer.
	var stringer Stringer
	stringer = p1 // valid because pair implements Stringer

	// Call String method of stringer, of type Stringer.
	// Output same as above.
	fmt.Println(stringer.String())

	// Functions in the fmt package call the String method to ask an object
	// for a printable representation of itself.
	fmt.Println(p1)       // Output same as above. Println calls String method.
	fmt.Println(stringer) // Output same as above.

}
