package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	atom := atomic.Int64{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atom.Add(1)
		}()
	}

	wg.Wait()
	fmt.Println("Result:", atom.Load())
}
