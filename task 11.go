package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- job * job
		time.Sleep(time.Second)
	}

}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int)
	results := make(chan int, 20)
	nums := make([]int, 0, 20)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for j := 1; j <= 20; j++ {
		jobs <- j
	}
	close(jobs)
	for i := 1; i <= 20; i++ {
		value := <-results
		nums = append(nums, value)

	}

	wg.Wait()

	fmt.Println(nums)
}
