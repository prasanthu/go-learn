package main

import (
	"fmt"
	"io"
	"os"
)

// Capper will change characters to Uppercase
type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	converted := make([]byte, len(p))
	dist := byte('a' - 'A')
	for i, v := range p {
		if v >= 'a' && v <= 'z' {
			converted[i] = v - dist
		} else {
			converted[i] = v
		}
	}
	return c.wtr.Write(converted)
}

func main() {
	c := Capper{os.Stdout}
	fmt.Fprintln(&c, "Hello there")
}
