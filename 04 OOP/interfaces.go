package main

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

type Square struct {
	Length float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c *Square) Area() float64 {
	return c.Length * c.Length
}

func sumAreas(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

type Shape interface {
	Area() float64
}

func main() {
	s := &Square{20}
	fmt.Println(s.Area())
	c := &Circle{20}
	fmt.Println(c.Area())
	p := Point{10, 11}
	fmt.Println(p)
	shapes := []Shape{s, c}
	fmt.Println(sumAreas(shapes))
}
