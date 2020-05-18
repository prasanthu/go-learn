package main

import (
	"fmt"
	"sync"
)

var numPoolsCreated = 0

func NewMemPool() interface{} {
	numPoolsCreated++
	mem := make([]byte, 1024)
	return &mem
}
func main() {
	memPool := sync.Pool{New: NewMemPool}
	memPool.Put(memPool.New())
	memPool.Put(memPool.New())
	memPool.Put(memPool.New())
	memPool.Put(memPool.New())
	const numWorkers = 1024 * 1024
	wg := sync.WaitGroup{}
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			mem := memPool.Get()
			defer memPool.Put(mem)
			// Do something with the memory
		}()
	}
	wg.Wait()
	fmt.Printf("Pools created %d by %d workers\n", numPoolsCreated, numWorkers)
}
