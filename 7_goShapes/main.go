package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct{
	baseLength float64
	height float64
}
type rectangle struct{
	length float64
	height float64
}

func main() {
	triangle := triangle{
		baseLength: 10,
		height: 20,
	}
	rectangle := rectangle{
		length: 20,
		height: 20,
	}

	printArea(triangle)
	printArea(rectangle)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return t.baseLength * t.height / 2
}

func (r rectangle) getArea() float64 {
	return r.length * r.height
}