package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	table := make([][]uint8, dy)

	for y := range table {
		row := make([]uint8, dx)
		for x := range row {
			row[x] = uint8(x*x + y*y)
		}
		table[y] = row
		fmt.Println(row)
		fmt.Println(table)
	}

	return table
}

func main() {
	pic.Show(Pic)
}
