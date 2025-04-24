package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	l1 := List[string]{nil, "hello"}
	l2 := List[string]{&l1, "world"}

	fmt.Println(l2.next.val, l2.val)
}
