package main

import "fmt"

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(i int) {
	i *= 2
}

func doublePtr(i *int) {
	*i *= 2
}

func main() {
	values := []int{1, 2, 3, 4}
	doubleAt(values, 2)
	fmt.Println(values)

	val := 10
	double(val)
	fmt.Println(val)

	doublePtr(&val)
	fmt.Println(val)

}
