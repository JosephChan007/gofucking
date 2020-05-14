package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type UserId int64

var wg sync.WaitGroup

func worker(ctx context.Context) {

	userVal := ctx.Value(UserId(111111)).(int64)
	fmt.Printf("user_id's value is: %d\n", userVal)

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
	ctx = context.WithValue(ctx, UserId(111111), int64(313324123412431131))

	wg.Add(1)
	go worker(ctx)

	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over....")

}
