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

//////////

type Sender interface {
	Send()
}

type Mail struct {
	mail string
}

// Interfaces are implemented implicitly:
// This method means type Mail implements the Sender Logger,
// but we don't need to explicitly declare that it does so.
func (mail Mail) Send() {
	fmt.Println(mail.mail)
}

func Run0() {
	var sender Sender = Mail{"hello"}
	sender.Send()
}

////////////

// what is empty interface
// The interface type that specifies zero methods is known as the empty interface
// An empty interface may hold values of any type. (Every type implements at least zero methods.)

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func seeEmptyInterface() {

	// Empty interfaces are used by code that handles values of unknown type.
	// For example, fmt.Print takes any number of arguments of type interface{}.
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

// Type assertions
// https://go.dev/tour/methods/15
// A type assertion provides access to an interface value's underlying concrete value.

// t := i.(T)
// This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

// If i does not hold a T, the statement will trigger a panic.

// To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

// t, ok := i.(T)
// If i holds a T, then t will be the underlying value and ok will be true.

// If not, ok will be false and t will be the zero value of type T, and no panic occurs.

// Note the similarity between this syntax and that of reading from a map.

// Type switches
// A type switch is a construct that permits several type assertions in series.

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func typeAsserts() {
	do(21)
	do("hello")
	do(true)
}

///////////////////////

/*
One of the most ubiquitous interfaces is Stringer defined by the fmt package.

type Stringer interface {
    String() string
}
A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
*/

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func SeePerson() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
