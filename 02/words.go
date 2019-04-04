package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	counts := map[string]int{}
	for _, val := range strings.Fields(text) {
		counts[strings.ToLower(val)]++
	}
	fmt.Println(counts)
}
