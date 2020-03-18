package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 20; i++ {
		divBy5 := (i % 5) == 0
		divBy3 := (i % 3) == 0
		switch {
		case divBy5 && divBy3:
			fmt.Println("fizz buzz")
		case divBy5:
			fmt.Println("buzz")
		case divBy3:
			fmt.Println("fizz")
		default:
			fmt.Println(i)
		}
	}
}
