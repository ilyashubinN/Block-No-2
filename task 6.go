package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("tick")
			return
		default:
			fmt.Println("running...")
			time.Sleep(time.Second)
		}
		// TODO ЛОГИКА
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go worker(ctx)

	time.Sleep(4 * time.Second)
}
