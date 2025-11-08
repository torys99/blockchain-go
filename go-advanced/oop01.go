package main

import "fmt"

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Area() {
	fmt.Printf("Rectangle Area: %.2f\n", r.Width*r.Height)
}
func (r *Rectangle) Perimeter() {
	fmt.Printf("Rectangle Perimeter: %.2f\n", 2*(r.Width+r.Height))
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() {
	fmt.Printf("Circle Area: %.2f\n", 3.14*c.Radius*c.Radius)
}
func (c *Circle) Perimeter() {
	fmt.Printf("Circle Perimeter: %.2f\n", 2*3.14*c.Radius)
}

func main() {
	var s Shape
	s = &Rectangle{Width: 5, Height: 10}
	s.Area()
	s.Perimeter()
	s = &Circle{Radius: 7}
	s.Area()
	s.Perimeter()
}
