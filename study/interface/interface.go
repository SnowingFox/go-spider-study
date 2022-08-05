package main

import (
	"fmt"
	"math"
)

type rect struct {
	w, h float64
}

type circle struct {
	r float64
}

func (r rect) area() float64 {
	return r.h * r.w
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func measure(g interface {
	area() float64
}) {
	fmt.Println(g.area())
}

func main() {
	r := rect{w: 3, h: 2}
	c := circle{r: 1}

	measure(r)
	measure(c)
}
