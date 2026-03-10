package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SafeNumber struct {
	value int
	mu    sync.RWMutex
}

func (s *SafeNumber) Get() int {
	// ...
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.value
}

func (s *SafeNumber) Set(v int) {
	s.mu.Lock()
	s.value = v
	s.mu.Unlock()
	// ...
}

func main() {
	s := &SafeNumber{}
	// инициализация SafeNumber
	// запуск нескольких горутин для чтения

	for i := 1; i <= 5; i++ {
		go func(id int) {
			for {
				value := s.Get()
				fmt.Println("Reader", id, "read:", value)
				time.Sleep(500 * time.Microsecond)
			}
		}(i)
	}
	go func() {
		for {
			v := rand.Intn(100)
			s.Set(v)
			fmt.Println("Writer set:", v)
			time.Sleep(1 * time.Second)
		}
	}()
	select {}

	// запуск отдельной горутины для обновления
}
