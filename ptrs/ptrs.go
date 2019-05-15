package main

import (
	"fmt"
)

func createPtrs(val float64) *float64 {
	var myVal = val
	return &myVal
}

func negate(myBoolean *bool) bool {
	*myBoolean = !(*myBoolean)
	return *myBoolean
}

func main() {
	fmt.Println(createPtrs(10.5))
	fmt.Println(createPtrs(45.3))
	fmt.Println(createPtrs(56.3))
	fmt.Println(createPtrs(156.3))
	myBoolean := false
	fmt.Println(negate(&myBoolean))
	fmt.Println(negate(&myBoolean))
	fmt.Println(negate(&myBoolean))
	fmt.Println(negate(&myBoolean))
}
