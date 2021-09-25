package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	with, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.with * r.height
}
func (r rect) perim() float64 {
	return (r.with + r.height) * 2
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return math.Pi * c.radius * 2
}

func messure(g geometry) {
	fmt.Printf("g=%v \n", g)
	fmt.Printf("perim=%v \n", g.perim())
	fmt.Printf("area=%v \n", g.area())
}

func main() {
	r := rect{
		with:   3,
		height: 4,
	}
	c := circle{
		radius: 4,
	}
	messure(r)
	messure(c)
}
