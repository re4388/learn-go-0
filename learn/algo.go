package learn

import (
	"fmt"

	"golang.org/x/tour/tree"
)

/*
Implement the Walk function and Same function using Walk to determine whether t1 and t2 store the same values.
*/

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkHelper(t, ch)
	close(ch)
}

func WalkHelper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	WalkHelper(t.Left, ch)
	ch <- t.Value
	WalkHelper(t.Right, ch)
}

// isSameTree determines whether the trees
// t1 and t2 contain the same values.
func isSameTree(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 2 goroutine go walk 2 tree and send to 2 channels
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1 {
		v2 := <-ch2
		if v1 != v2 {
			return false
		}
	}

	return true
}

func Is_the_same_tree() {
	fmt.Println(isSameTree(tree.New(1), tree.New(1)))
	fmt.Println(isSameTree(tree.New(1), tree.New(2)))
}
