package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	slicedString := strings.Fields(s)
	// fmt.Println(slicedString)
	m := make(map[string]int)

	for _, v := range slicedString {
		m[v] += 1
	}
	// fmt.Println(m)

	return m
}

func main() {
	wc.Test(WordCount)
}
