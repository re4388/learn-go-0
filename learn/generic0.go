package learn

import "fmt"

/*
Type parameters

Go functions can be written to work on multiple types using type parameters.
The type parameters of a function appear between "brackets", before the function's arguments.

func Index[T comparable](s []T, x T) int

This declaration means that s is a slice of any type T
that fulfills the built-in constraint comparable. x is also a value of the same type.

comparable is a useful constraint that makes it possible to use the == and != operators on values of the type.
In this example, we use it to compare a value to all slice elements until a match is found.
This Index function works for any type that supports comparison.
*/

// Index returns the index of target in slice, or -1 if not found.
func Index[T comparable](slice []T, target T) int {
	for idx, val := range slice {
		// val and target are type T, which has the comparable
		// constraint, so we can use == here.
		if val == target {
			return idx
		}
	}

	return -1
}

func RunIndex() {
	// Index works on a slice of ints
	slice0 := []int{10, 20, 15, -10}
	fmt.Println(Index(slice0, 15))

	// Index also works on a slice of strings
	slice1 := []string{"foo", "bar", "baz"}
	fmt.Println(Index(slice1, "hello"))
}

/*
In addition to "generic functions", Go also supports "generic types".
A type can be parameterized with a type parameter, which could be useful for implementing generic data structures.

This example demonstrates a simple type declaration for a "singly-linked list" holding "any type" of value.
As an exercise, add some functionality to this list implementation.
*/

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}
