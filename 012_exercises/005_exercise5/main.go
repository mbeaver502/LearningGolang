package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func main() {
	s := square{
		side: 5,
	}

	c := circle{
		radius: 7,
	}

	info(s)
	info(c)
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func info(s shape) {
	fmt.Println(s.area())
}
