package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sidelength float64
}

func main() {
	t := triangle{
		base:   10.23,
		height: 12.34,
	}

	s := square{
		sidelength: 20.34,
	}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
