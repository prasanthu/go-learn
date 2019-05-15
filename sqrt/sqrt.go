package main

import (
	"fmt"
	"math"
)

func sqrt(val float64) (float64, error) {
	if val < 0 {
		return 0, fmt.Errorf("Cannot find square root of negative number %0.3f", val)
	}
	return math.Sqrt(val), nil
}
func main() {
	val := -9.2
	r, err := sqrt(val)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Square root of %0.3f == %0.3f\n", val, r)
	}
}
