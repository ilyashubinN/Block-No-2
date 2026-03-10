package main

//Сумма чисел с распределением задач

import (
	"fmt"
	"math/rand"
)

func sumPart(nums []int, resultCh chan int) {
	// TODO: Посчитать сумму и отправить в канал

	sum := 0
	for _, v := range nums {
		sum += v
	}
	resultCh <- sum
}

func main() {
	const parts = 10
	nums := make([]int, 1000)

	for i := range nums {
		nums[i] = rand.Intn(100)
	}

	resultCh := make(chan int, parts)

	chunk := len(nums) / parts

	for i := 0; i < len(nums); i += chunk {
		part := nums[i : i+chunk]

		go sumPart(part, resultCh)

	}
	total := 0
	for i := 0; i < parts; i++ {
		total += <-resultCh
	}
	// TODO: Разделить nums на 10 частей и запустить 10 горутин sumPart

	// TODO: Собрать результаты из канала и вывести общую сумму

	fmt.Println(total)

}
