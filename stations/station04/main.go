package main

import (
	"fmt"
	"math"
)

func main() {
	rectangle := Rectangle{
		Width:  30,
		Height: 15.5,
	}

	circle := Circle{
		Radius: 10,
	}

	printArea(rectangle)
	printArea(circle)
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Println("面積: ", s.Area(), "㎠")
}
