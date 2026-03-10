// гонки горутин
package main

import (
	"context"
	"fmt"
	"time"
)

func number1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("tick")
			return
		default:
			fmt.Println("running...")
			time.Sleep(time.Second)
		}
	}

}

func number2(ddx context.Context) {
	for {
		select {
		case <-ddx.Done():
			fmt.Println("tick")
			return
		default:
			fmt.Println("running...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	ddx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go number1(ctx)
	go number2(ddx)

	time.Sleep(4 * time.Second)
}
