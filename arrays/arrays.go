package main

import "fmt"

func writeToArray(a *[2]string) {
	a[0] = "bbc"
	a[1] = "cnn"
}
func writeToSlice(s *[]string) {
	*s = append(*s, "bbc")
	*s = append(*s, "cnn")
}
func main() {
	fmt.Println("arrays")
	a := [2]string{"dd", "cc"}
	writeToArray(&a)
	for k, v := range a {
		fmt.Println(k, v)
	}
	fmt.Println("Slices")
	s := make([]string, 0)
	writeToSlice(&s)
	for k, v := range s {
		fmt.Println(k, v)
	}
	dd := New2DArray(3, 3)
}
