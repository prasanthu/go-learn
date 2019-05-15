package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	rand.Seed(int64(now.Nanosecond()))
	r := rand.Intn(100)
	fmt.Println("I have gussed a random number")
	fmt.Println("Can you guess it ?")
	reader := bufio.NewReader(os.Stdin)
	guessed := false

	for i := 1; i <= 3; i++ {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		n, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if n < r {
			fmt.Println("You guessed it LOW")
		} else if n > r {
			fmt.Println("You guessed it HIGH")
		} else {
			fmt.Println("You guessed it right")
			guessed = true
			break
		}
	}
	if guessed == false {
		fmt.Println("You failed to guess the number. It was ", r)
	}
}
