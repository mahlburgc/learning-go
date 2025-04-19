package main

import "fmt"

type IPAddr [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main() {
	hosts := map[string]IPAddr{
		"google":   {8, 8, 8, 8},
		"nextbike": {12, 243, 23, 23},
	}

	for name, ip := range hosts {
		fmt.Printf("Address: %v, IP: %v\n", name, ip)
	}
}
