package main

import "fmt"

type square struct{ sidelength float64 }
type triangle struct {
	height float64
	base   float64
}
type shape interface {
	getArea() float64
}

func main() {
	sq := square{12.5}
	tr := triangle{17.45,
		5.2}
	printArea(sq)
	printArea(tr)

}

func (s square) getArea() float64 {

	return s.sidelength * s.sidelength
}

func (t triangle) getArea() float64 {

	return (t.base * 0.5) * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())

}
