package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//атомарный счётчик

func main() {
	var counter atomic.Int64
	var counter2 atomic.Int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}

	for i := 0; i < 1500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter2.Add(1)
		}()
	}

	wg.Wait()

	fmt.Println(counter.Load() + counter2.Load())
}
