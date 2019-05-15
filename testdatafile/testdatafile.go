package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"study/average"
	"study/datafile"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal(errors.New("Provide the filename"))
	}
	numbers, err := datafile.GetFloats(args[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Numbers = ", numbers)
	fmt.Println("Average = ", average.Average(numbers...))
}
