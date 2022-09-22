package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type shape2 interface {
	perim() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

func (r *rect) area() float64 {
	return r.height * r.width
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r *rect) perim() float64 {
	return r.height + r.width
}

func getArea(s shape) float64 {
	return s.area()
}

func getPerim(s shape2) float64 {
	return s.perim()
}

func main() {
	// c1 := circle{4}
	r1 := rect{5, 7}
	// shapes := []shape{&c1, &r1}
	shapes2 := []shape2{&r1}
	for _, shape := range shapes2 {
		fmt.Println(getPerim(shape))
	}

}
