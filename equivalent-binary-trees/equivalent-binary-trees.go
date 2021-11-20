package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)
	var t1Vals [10]int
	go Walk(t1, ch)
	for i := 0; i < 10; i++ {
		t1Vals[i] = <-ch
	}
	var t2Vals [10]int
	go Walk(t2, ch)
	for i := 0; i < 10; i++ {
		t2Vals[i] = <-ch
	}
	return t1Vals == t2Vals
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(1)))
}
