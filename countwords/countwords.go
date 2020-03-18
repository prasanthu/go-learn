package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	words := make(map[string]int, 0)
	fields := strings.Fields(text)
	for _, v := range fields {
		word := strings.ToLower(v)
		words[word]++
	}
	fmt.Println(words)
}
