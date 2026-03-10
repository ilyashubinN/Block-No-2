package main

import (
	"fmt"
	"time"
)

func generator(ch chan int) {
	// TODO: Отправить числа 1-5 в канал
	// TODO: Закрыть канал
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int) {
	// TODO: Считать из канала и печатать числа
	for v := range ch {
		fmt.Println(v)
	}

}

func main() {
	ch := make(chan int)

	go generator(ch)
	go consumer(ch)
	time.Sleep(2 * time.Second)
	// TODO: Подождать завершения (через time.Sleep или sync.WaitGroup)
}
