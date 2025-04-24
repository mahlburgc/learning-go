package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)

	if err != nil {
		return n, err
	}

	for i := range b {
		switch {
		case (b[i] >= 'A' && b[i] <= ('Z'-13)) || (b[i] >= 'a' && b[i] <= ('z'-13)):
			fmt.Printf("%v -> %v\n", string(b[i]), string(b[i]+13))
			b[i] += 13
		case (b[i] > 'z'-13 && b[i] <= 'z') || (b[i] > 'Z'-13 && b[i] <= 'Z'):
			fmt.Printf("%v -> %v\n", string(b[i]), string(b[i]-13))
			b[i] -= 13
		}
	}
	return len(b), io.EOF
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
