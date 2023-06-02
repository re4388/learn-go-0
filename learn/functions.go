package learn

import "fmt"

// first class fn
func AddN(n int) func(x int) int {
	return func(x int) int {
		return n + n
	}
}

// It is possible, unlike in many other languages for functions in go
// to have named return values.
// Assigning a name to the type being returned in the function
// declaration line
// this allows us to easily return from multiple points in a function
// as well as to only use the return keyword, without anything further.
func NamedReturns(x, y int) (z int) {
	z = x * y
	return // z is implicit here, because we named it earlier.
}

// Functions can have variadic parameters.
func VariadicParams(strArr ...interface{}) {
	// Iterate each value of the variadic.
	for _, param := range strArr {
		fmt.Println("param:", param)
	}

	// Pass variadic value as a variadic parameter.
	fmt.Println("res:", fmt.Sprintln(strArr...))
}

func FunctionFactory() {
	// Next two are equivalent, with second being more practical
	fmt.Println(sentenceFactory("summer")("A beautiful", "day!"))

	fn1 := sentenceFactory("summer")
	fmt.Println(fn1("A beautiful", "day!"))
	fmt.Println(fn1("A lazy", "afternoon!"))
}

// Decorators are common in other languages. Same can be done in Go
// with function literals that accept arguments.
func sentenceFactory(s1 string) func(s2, s3 string) string {
	return func(s2, s3 string) string {
		return fmt.Sprintf("%s %s %s", s2, s1, s3) // new string
	}
}

// Go is fully garbage collected.
// It has pointers but no pointer arithmetic.
// You can make a mistake with a nil pointer, but not
// by incrementing a pointer.
// Unlike in C/Cpp taking and returning an address of a local variable
// is also safe.
func learnMemory() (p, q *int) { // *int means int ptr

	// new is built-in to new a allocated int memory
	// p = new(int)

	// The allocated int slice is initialized to 0
	s := make([]int, 20)
	s[3] = 7

	r := -2

	// return 2 int ptr
	return &s[3], &r
}

func GetPtr() {
	p, q := learnMemory() // Declares p, q to be type pointer to int.
	// * follows a pointer. This prints two ints.
	// like de-reference
	fmt.Println(*p, *q)
}

// return multiple value
func IsEven(n int) (bool, error) {
	if n <= 0 {
		return false, fmt.Errorf("error")
	}

	return n%2 == 0, nil // why above is error but here we can put nil?
}

/*
Go functions may be closures.
A closure is a function value that references variables from outside its body.
The function may access and assign to the referenced variables;
in this sense the function is "bound" to the variables.

For example, the adder function returns a closure.
Each closure is bound to its own sum variable.
*/

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func ClosureInAction() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
