package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

func getArea(s Shape) float64 {
	return s.area()
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{10, 5}

	fmt.Println("circle area:", getArea(circle))
	fmt.Println("rectangle area:", getArea(rectangle))
}
