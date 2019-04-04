package main

import (
	"fmt"
)

func main() {
	book := "Цвет магии"
	// book := "The colour of magic"
	fmt.Println(book)
	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])

	fmt.Println(book[4:11])
	fmt.Println(book[4:])
	fmt.Println(book[:4])

	fmt.Println("t" + book[1:])

	// Unicode
	fmt.Println("It was ½ price!")

	poem := `
	Стих строфа1
	Стих строфа2
	..
	`
	fmt.Println(poem)

}
