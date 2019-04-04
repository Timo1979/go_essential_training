package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		ch2 <- 2
		ch1 <- 42
		ch1 <- 44
	}()
	select {
	case val := <-ch1:
		fmt.Printf("got %d from ch1\n", val)
	case val := <-ch2:
		fmt.Printf("got %d from ch2\n", val)
	}

	fmt.Println("####")

	out := make(chan float64)
	go func() {
		time.Sleep(300 * time.Millisecond)
		out <- 3.14
	}()
	select {
	case val := <-out:
		fmt.Printf("Got %v after 200ms\n", val)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("timeout!")
	}
}
