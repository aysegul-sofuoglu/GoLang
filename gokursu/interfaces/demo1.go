package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) Area() float64 {
	return r.width * r.height
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func School(s shape) {
	fmt.Println("şeklin alanı: ", s.Area())
}

func Demo1() {
	r := rectangle{width: 10, height: 7}
	c := circle{radius: 1}
	School(r)
	School(c)
}
