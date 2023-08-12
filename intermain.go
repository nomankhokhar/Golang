package main

import "fmt"

// Define an interface named "Shape"
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a struct type "Circle"
type Circle struct {
	Radius float64
}

// Define methods for the Circle type to implement the Shape interface
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// Define a struct type "Rectangle"
type Rectangle struct {
	Width  float64
	Height float64
}

// Define methods for the Rectangle type to implement the Shape interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Both Circle and Rectangle implement the Shape interface,
	// so they can be used interchangeably with printShapeInfo
	printShapeInfo(circle)
	printShapeInfo(rectangle)
}
