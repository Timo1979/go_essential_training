package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 353
	}()
	val := <-ch
	fmt.Printf("Got %d\n", val)
	fmt.Println("####")
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()
	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("got %d\n", val)
	}
	fmt.Println("####")
	go func() {
		for i := 10; i < 13; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()
	for val := range ch {
		fmt.Printf("got %d\n", val)
	}

}
