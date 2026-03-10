package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	counter := 0
	// TODO: sync.Mutex для защиты counter
	// TODO: sync.WaitGroup для ожидания завершения всех горутин

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// TODO: Запустить горутину, которая увеличит counter
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	// TODO: Дождаться завершения всех горутин

	fmt.Println("Result:", counter)
}
