package main

import (
	"context"
	"fmt"
)

func run(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			case dst <- n:
				n++
			}
		}
	}()

	return dst
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := range run(ctx) {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}

}
