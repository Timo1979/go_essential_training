package main

import (
	"fmt"
	"log"
)

type Point struct {
	X int
	Y int
}

type Square struct {
	Center Point
	Length int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

func NewSquare(x int, y int, len int) (*Square, error) {
	if len < 0 {
		return nil, fmt.Errorf("lenght must be >= 0")
	}
	return &Square{Point{x, y}, len}, nil
}

func main() {
	s, err := NewSquare(1, 2, 7)
	if err != nil {
		log.Fatalf("ERROR: can't create square: %v\n", err)
	}
	fmt.Printf("%+v\n", s)

	fmt.Printf("%d\n", s.Area())
	s.Center.Move(-10, -10)
	fmt.Printf("%v\n", s)
}
