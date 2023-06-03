package learn

import (
	"fmt"
	"strings"
)

/*
Implement WordCount. It should return a map of the counts of each "word" in the string s.
*/
func wordCount(s string) map[string]int {
	map0 := make(map[string]int)
	words := strings.Fields(s)

	fmt.Printf("words %#v\n", words)

	for _, word := range words {
		map0[word] += 1
	}

	return map0
}

func Run_wordCount() {
	res := wordCount("hello world")
	fmt.Printf("res %#v\n", res)
}
