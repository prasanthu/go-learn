package main

import (
	"fmt"
)

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	return area / 10.0
}

func main() {
	var amount, total float64

	amount = paintNeeded(4.2, 3.0)
	fmt.Printf("Amount of paint needed %.2f liters\n", amount)
	total += amount

	amount = paintNeeded(5.2, 3.5)
	fmt.Printf("Amount of paint needed %.2f liters\n", amount)
	total += amount

	fmt.Printf("Total Amount of paint needed %.2f liters\n", total)
}
