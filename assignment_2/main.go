package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type shape interface {
	getArea() float64
}

func printArea(s shape) float64 {
	return s.getArea()
}

func main() {
	s := square{
		sideLength: 5,
	}
	t := triangle{
		height: 4,
		base:   3,
	}
	fmt.Println(printArea(s))
	fmt.Println(printArea(t))
}
