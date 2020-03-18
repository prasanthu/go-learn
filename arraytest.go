package main

import (
	"fmt"

	"github.com/prasanthu/go-learn/arrays"
)

func main() {
	fmt.Println("arrays")
	a := [2]string{"dd", "cc"}
	arrays.WriteToArray(&a)
	for k, v := range a {
		fmt.Println(k, v)
	}
	fmt.Println("Slices")
	s := make([]string, 0)
	arrays.WriteToSlice(&s)
	for k, v := range s {
		fmt.Println(k, v)
	}
	test2DSlices()
}

func test2DSlices() {
	dd := arrays.New2DSlice(5, 5)
	arrays.IdentityFill(dd)
	for k := range dd {
		for _, v := range dd[k] {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
