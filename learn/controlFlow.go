package learn

import (
	"fmt"
	m "math" // Math library with local alias m.
)

func LearnControlFlow() {

	// If statements require brace brackets, and do not require parentheses.
	if true {
		fmt.Println("told ya")
	}
	// Formatting is standardized by the command line command "go fmt".
	if false {
		// Pout.
	} else {
		// Gloat.
	}
	// Use switch in preference to chained if statements.
	x := 42.0
	switch x {
	case 0:
	case 1, 2: // Can have multiple matches on one case
	case 42:
		// Cases don't "fall through".
		/*
		   There is a `fallthrough` keyword however, see:
		     https://github.com/golang/go/wiki/Switch#fall-through
		*/
	case 43:
		// Unreached.
	default:
		// Default case is optional.
	}

	// Type switch allows switching on the type of something instead of value
	var data interface{}
	data = ""
	switch c := data.(type) {
	case string:
		fmt.Println(c, "is a string")
	case int64:
		fmt.Printf("%d is an int64\n", c)
	default:
		// all other cases
	}

	// Like if, for doesn't use parens either.
	// Variables declared in for and if are local to their scope.
	for x := 0; x < 3; x++ { // ++ is a statement.
		fmt.Println("iteration", x)
	}
	// x == 42 here.

	// For is the only loop statement in Go, but it has alternate forms.
	for { // Infinite loop.
		break    // Just kidding.
		continue // Unreached.
	}

	// You can use range to iterate over an array, a slice, a string, a map, or a channel.
	// range returns one (channel) or two values (array, slice, string and map).
	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		// for each pair in the map, print key and value
		fmt.Printf("key=%s, value=%d\n", key, value)
	}
	// If you only need the value, use the underscore as the key
	for _, name := range []string{"Bob", "Bill", "Joe"} {
		fmt.Printf("Hello, %s\n", name)
	}

	// As with for, := in an if statement means to declare and assign
	// y first, then test y > x.
	if y := expensiveComputation(); y > x {
		x = y
	}

	// Function literals are closures.
	xBig := func() bool {
		return x > 10000 // References x declared above switch statement.
	}
	x = 99999
	fmt.Println("xBig:", xBig()) // true
	x = 1.3e3                    // This makes x == 1300
	// 跟 JS 的 closure 不一樣，by reference, 會每次呼叫都會去拿外面變動的x
	fmt.Println("xBig:", xBig()) // false now.

	// function literals
	// function can defined and called inline,
	// it act as an argument to function, as long as:
	// a) function literal is called immediately (),
	// b) result type matches expected type of argument.
	fmt.Println("Add + double two numbers: ",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2)) // Called with args 10 and 2
	// => Add + double two numbers: 24

	// goto
	goto love
love:
	learnDefer() // A quick detour to an important keyword.

}

func learnDefer() (ok bool) {
	// A defer statement pushes a function call onto a list. The list of saved
	// calls is executed AFTER the surrounding function returns.
	defer fmt.Println("deferred statements execute in reverse (LIFO) order.")
	defer fmt.Println("\nThis line is being printed first because")
	// Defer is commonly used to close a file, so the function closing the
	// file stays close to the function opening the file.

	fmt.Println("shall run first in this fn")
	fmt.Println("shall run second in this fn")
	return true
}

// Use the aliased math library (see imports, above)
func expensiveComputation() float64 {
	return m.Exp(10)
}
