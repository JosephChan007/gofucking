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
			fmt.Println(ctx.Err())
			break LOOP
		default:
		}
	}
	fmt.Println("worker done....")
	wg.Done()
}

func master() {

	d := time.Now().Add(time.Millisecond * 100)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	wg.Add(1)
	go worker(ctx)

	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over....")

}

func main() {
	//master()

	d := time.Now().Add(time.Millisecond * 2000)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(time.Second * 1):
		fmt.Println("over sleep....")
	}

}
