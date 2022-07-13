package main

import (
	"context"
	"fmt"
	"time"
)

//https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go

func doSomething2(ctx context.Context) {
	deadline := time.Now().Add(1500 * time.Millisecond)
	ctx, cancelCtx := context.WithDeadline(ctx, deadline)
	defer cancelCtx()

	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			break
		}
	}

	cancelCtx()

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("doSomething: finished\n")
}

func doSomething(ctx context.Context) {
	ctx, cancelCtx := context.WithCancel(ctx)

	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		printCh <- num
	}

	cancelCtx()

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("doSomething: finished\n")
}

func doAnother(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			fmt.Printf("doAnother: finished\n")
			return
		case num := <-printCh:
			fmt.Printf("doAnother--start: %d\n", num)
			//time.Sleep(2 * time.Second)
			doThird(ctx)
			fmt.Printf("doAnother--end: %d\n", num)
		}
	}
}

func doThird(ctx context.Context) {
	select {
	case <-ctx.Done():
		if err := ctx.Err(); err != nil {
			fmt.Printf("doThird err: %s\n", err)
		}
		fmt.Print("doThird: finished\n")
	default:
		fmt.Print("do third default start\n")
		time.Sleep(1 * time.Second)
		fmt.Print("do third default end\n")
	}

}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "myKey", "myValue")

	doSomething2(ctx)
}
