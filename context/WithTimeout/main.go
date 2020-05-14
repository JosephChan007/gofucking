package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {

LOOP:
	for {
		fmt.Println("connetion....")
		time.Sleep(time.Millisecond * 50)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done....")
	wg.Done()
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)

	wg.Add(1)
	go worker(ctx)

	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over....")

}
