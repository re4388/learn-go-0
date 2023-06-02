package learn

import (
	"fmt"
	"strconv"
)

func LearnErrorHandling() {

	map0 := map[int]string{
		3: "three",
		4: "four",
	}

	// ", ok" idiom used to tell if something worked or not.
	// ok will be false because 1 is not in the map.
	if x, ok := map0[3]; !ok {
		fmt.Println("no one there")
	} else {
		fmt.Println(x)  // three
		fmt.Println(ok) // true
	}

	// An error value communicates not just "ok"
	// but more about the problem.
	if _, err := strconv.Atoi("non-int"); err != nil { // _ discards value
		// prints 'strconv.ParseInt: parsing "non-int": invalid syntax'
		fmt.Println(err)
	}
}
