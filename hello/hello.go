package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/prasanthu/go-learn/average"
)

func main() {
	args := os.Args[1:]
	var numbers []float64
	for _, arg := range args {
		number, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	fmt.Println("Average = ", average.Average(numbers...))
}

/*
import (
	"fmt"
	"strings"
	"study/greeting"
	"time"
)
func main() {
	now := time.Now()
	broken := "Replace hashes ##"
	replacer := strings.NewReplacer("#", ":-)")
	fixed := replacer.Replace(broken)
	fmt.Println("hello world", now, fixed)
	greeting.Hi()
	greeting.Hello()
}
*/
