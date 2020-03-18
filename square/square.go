package main

import "fmt"

// Point structure
type Point struct {
	X int
	Y int
}

// Move the Point
func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

// Square structure
type Square struct {
	Center Point
	Length int
}

// Move for Square
func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

// Area of the Square
func (s *Square) Area() int {
	return s.Length * s.Length
}

// NewSquare creates a new Square
func NewSquare(x int, y int, length int) (*Square, error) {
	return &Square{Point{x, y}, length}, nil
}

func main() {
	s, _ := NewSquare(10, 10, 100)
	fmt.Println(s.Area())
}
