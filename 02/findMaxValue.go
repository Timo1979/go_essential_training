package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	max := -10000000
	for _, val := range nums {
		if val > max {
			max = val
		}
	}
	fmt.Println(max)
}
