package main

import (
	"fmt"
)

func ret1() int {
	fmt.Println("ret1")
	return 1
}

func ret2() int {
	fmt.Println("ret2")
	return 2
}

func main() {
	x := 10
	if x > 5 {
		fmt.Println("x greater than 5")
	}

	if x > 100 {
		fmt.Println("x is very big")
	} else {
		fmt.Println("x is not very big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	if x < 20 || x > 30 {
		fmt.Println("x is out of range")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}

	s := []string{"1", "2"}
	if ret1() == 2 || s[2] == "1" {
		fmt.Printf("len = %d\n", len(s))
	}
}
