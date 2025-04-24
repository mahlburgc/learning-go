package main

import (
	"fmt"
	"sort"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, c chan int, cls bool) {
	if t != nil {
		c <- t.Value
		Walk(t.Left, c, false)
		Walk(t.Right, c, false)
	}

	if cls {
		close(c)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1, true)
	go Walk(t2, c2, true)

	var arr1 []int
	for i := range c1 {
		arr1 = append(arr1, i)
	}
	sort.Ints(arr1)
	fmt.Println(arr1)

	var arr2 []int
	for i := range c2 {
		arr2 = append(arr2, i)
	}
	sort.Ints(arr2)
	fmt.Println(arr2)

	if len(arr1) != len(arr2) {
		return false
	}

	for i := range len(arr1) {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
