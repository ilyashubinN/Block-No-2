package main

import (
	"fmt"
	"sync"
)

func producer(ch1 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch1 <- i
	}
	close(ch1)
}
func square(ch1 chan int, ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch1 {
		ch2 <- v * v
	}
	close(ch2)
}
func printer(ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch2 {
		fmt.Println(v)
	}
}
func main() {
	wg := sync.WaitGroup{}
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg.Add(3)
	go producer(ch1, &wg)
	go square(ch1, ch2, &wg)
	go printer(ch2, &wg)
	wg.Wait()
	// ожидание завершения (sync.WaitGroup)
}
