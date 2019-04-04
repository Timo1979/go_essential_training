package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int64
	var j int64
	count := 0
	for i = 1000; i <= 9999; i++ {
		for j = i; j <= 9999; j++ {
			x := strconv.FormatInt(i*j, 10)
			if x[0] == x[len(x)-1] {
				// fmt.Println(x)
				count++
			}
		}
	}
	fmt.Println(count)
}
