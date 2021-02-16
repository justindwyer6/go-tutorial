package main

import (
	"fmt"
	"math"
	"strconv"
)

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	tr := triangle{
		height: 10,
		base:   3,
	}
	sq := square{
		sideLength: 4,
	}

	printArea(tr)
	printArea(sq)
}

func (t triangle) getArea() float64 {
	return t.height * t.base * 0.5
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func printArea(sh shape) {
	fmt.Println("The area of this shape is " + strconv.FormatFloat(sh.getArea(), 'f', 5, 64))
}
