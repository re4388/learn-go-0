package learn

import "fmt"

func LearnTypes() {
	str := "this is called short declaration"

	s2 := `A "raw" string literal
can include line breaks.` // Same string type.

	// Non-ASCII literal. Go source is UTF-8.
	// rune type, an alias for int32, holds a unicode code point.
	g := 'Σ'

	f := 3.14195       // float64, an IEEE-754 64-bit floating point number.
	complex0 := 3 + 4i // complex128, represented internally with two float64's.

	// var syntax with initializers.
	var u uint = 7 // Unsigned, but implementation dependent size as with int.
	var pi float32 = 22. / 7

	// Conversion syntax with a short declaration.
	n := byte('\n') // byte is an alias for uint8.

	// array is fixed size at compile time
	var arr4 [4]int // init 4 int with 0 value

	arr5 := [...]int{3, 1, 5, 10, 100} // init 5 val with specific value

	// Arrays have value semantics
	// (compare to reference semantics for slice, maps and channels)
	arr4_cpy := arr4 // a4_cpy is a copy of a4, 2 instance

	// modify
	arr4_cpy[0] = 25 // only a4_cpy is changed, a4 is untouched

	fmt.Println(arr4_cpy[0] == arr4[0]) // false

	// slices has dynamic size
	slice3 := []int{4, 5, 9} // Compare to a5 array. No ellipsis for slice

	// Allocates slice of 4 ints, initialized to all 0
	slice4 := make([]int, 4)

	// Declaration only, nothing allocated here.
	var slice_declare [][]float64

	// Type conversion syntax
	bs := []byte("a slice")

	// slice is ref sematic, so s3_cpy and s3 point to
	// the same instance
	s3_cpy := slice3

	s3_cpy[0] = 0
	fmt.Println(s3_cpy[0] == slice3[0]) // true

	// Because they are dynamic, slices can be appended to on-demand.
	// To append elements to a slice, the built-in append() function is used.
	// First argument is a slice to which we are appending. Commonly,
	// the array variable is updated in place, as in example below.
	slice0 := []int{1, 2, 3}         // Result is a slice of length 3.
	slice0 = append(slice0, 4, 5, 6) // Added 3 elements. Slice now has length of 6.
	fmt.Println(slice0)              // Updated slice is now [1 2 3 4 5 6]

	// To append another slice, instead of list of atomic elements we can
	// pass a reference to a slice or a slice literal like this, with a
	// trailing ellipsis, meaning take a slice and unpack its elements,
	// appending them to slice s.
	slice0 = append(slice0, []int{7, 8, 9}...) // Second argument is a slice literal.
	fmt.Println(slice0)                        // Updated slice is now [1 2 3 4 5 6 7 8 9]

	// Maps are a dynamically growable associative array type, like the
	// hash or dictionary types of some other languages.
	map0 := map[string]int{"three": 3, "four": 4}
	map0["one"] = 1

	// Unused variables are an error in Go.
	// The underscore lets you "use" a variable but discard its value.
	_, _, _, _, _, _, _, _, _, _ = str, s2, g, f, u, pi, n, arr5, slice4, bs

	// Output of course counts as using a variable.
	fmt.Println(slice0, complex0, arr4, slice3, slice_declare, map0)

	// type conversions
	var i int = 42
	var f3 float64 = float64(i)
	var u2 uint = uint(f3)
	_, _ = f3, u2

}
