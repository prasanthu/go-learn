package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Request struct {
	args          []int
	f             func([]int) int
	resultChannel chan int
}

func sum(args []int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}

func makeRange(max int) []int {
	a := make([]int, max)
	for i := range a {
		a[i] = i + 1
	}
	return a
}

func createRequests(count int) []*Request {
	requests := make([]*Request, count)
	for i := 0; i < count; i++ {
		args := makeRange(rand.Intn(10) + 1)
		resultChannel := make(chan int, 1)
		requests[i] = &Request{args, sum, resultChannel}
	}
	return requests
}

func handle(queue chan *Request) {
	fmt.Println("Handler waiting for requests")
	for req := range queue {
		// If there is no buffer, it will wait for somebody to pickup
		req.resultChannel <- req.f(req.args)
	}
	fmt.Println("Handler done")
}

const MaxOutstanding = 10

func main() {
	rand.Seed(time.Now().UnixNano())
	queue := make(chan *Request, MaxOutstanding)
	requests := createRequests(rand.Intn(50) + 1)
	go handle(queue)
	for _, r := range requests {
		queue <- r
	}
	fmt.Printf("Created total %d requests\n", len(requests))
	for _, r := range requests {
		// Pickup result from the channel
		result := <-r.resultChannel
		fmt.Printf("Sum of %v = %d\n", r.args, result)
	}
}
