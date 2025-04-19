package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(p []byte) (int, error) {
	i := 0
	for i = range p {
		p[i] = 'A'
	}
	return i, nil
}

func main() {
	reader.Validate(MyReader{})
}
