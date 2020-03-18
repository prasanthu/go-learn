package main

import "fmt"

func New2DSlice(x, y int) [][]int {
	var a = make([][]int, y)
	var backing = make([]int, x*y)
	for i := range a {
		a[i], backing = backing[:x], backing[x:]
	}
	return a
}

func IdentityFill(a [][]int) {
	for i := range a {
		for j := range a[i] {
			a[i][j] = 0
		}
		a[i][i] = 1
	}
}

func main() {
	dd := New2DSlice(5, 5)
	IdentityFill(dd)
	for k := range dd {
		for _, v := range dd[k] {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
