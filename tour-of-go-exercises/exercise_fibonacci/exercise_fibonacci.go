package main

import (
	"fmt"
)

func fibonacci() func() int {
	f0 := 0
	f1 := 1
	return func() int {
		n := f0 + f1
		f0 = f1
		f1 = n
		return n
	}
}

func main() {
	a := fibonacci()
	for range 10 {
		fmt.Println(a())
	}
}
