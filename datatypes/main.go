package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

type Rectangle struct {
	TopLeft     Point
	BottomRight Point
}

func (r Rectangle) Width() float64 {
	return r.BottomRight.X - r.TopLeft.X
}

func (r Rectangle) Height() float64 {
	return r.BottomRight.Y - r.TopLeft.Y
}

func (r Rectangle) Area() float64 {
	return r.Width() * r.Height()
}

func (r *Rectangle) Enlarge(factor float64) {
	// Enlarge the width and the height by the given factor
	r.BottomRight.X = r.TopLeft.X + r.Width()*factor
	r.BottomRight.Y = r.TopLeft.Y + r.Height()*factor
}

type Circle struct {
	Center Point
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

const (
	WHITE int = 0xFFFFFF
	BLACK int = 0x000000
	RED   int = 0xFF0000
	GREEN int = 0x00FF00
	BLUE  int = 0x0000FF
)

type ColoredCircle struct {
	Circle
	Color int
}

func (c ColoredCircle) GetColor() int {
	return c.Color
}

type Shape interface {
	Area() float64
}

type Colored interface {
	GetColor() int
}

func main() {
	point := Point{X: 1.0, Y: 2.0}
	fmt.Println(point)

	rectangle := Rectangle{
		TopLeft:     Point{X: 1.0, Y: 2.0},
		BottomRight: Point{X: 3.0, Y: 4.0},
	}
	fmt.Println(rectangle)
	fmt.Printf("Width: %f, Height: %f, Area: %f\n", rectangle.Width(), rectangle.Height(), rectangle.Area())

	circle := Circle{
		Center: point,
		Radius: 3.0,
	}
	fmt.Println(circle)
	fmt.Printf("Area: %f\n", circle.Area())

	coloredCircle := ColoredCircle{
		Circle: circle,
		Color:  RED,
	}
	fmt.Println(coloredCircle)
	fmt.Printf("Area: %f\n", coloredCircle.Area())
	inlinePoint := struct {
		X float64
		Y float64
	}{X: 1.0, Y: 2.0}
	fmt.Println(inlinePoint)

	rectangleToEnlarge := Rectangle{
		TopLeft:     Point{X: 1.0, Y: 1.0},
		BottomRight: Point{X: 3.0, Y: 4.0},
	}
	fmt.Println(rectangleToEnlarge)
	rectangleToEnlarge.Enlarge(2.0)
	fmt.Println(rectangleToEnlarge)

	shapes := []Shape{circle, rectangle, coloredCircle}
	for _, shape := range shapes {
		fmt.Printf("Area: %f\n", shape.Area())
		if colored, ok := shape.(Colored); ok {
			fmt.Printf("Color: %d\n", colored.GetColor())
		}
	}

	var anything any
	anything = "asdf"
	fmt.Println(anything)
	anything = 123
	fmt.Println(anything)
	anything = Circle{}
	fmt.Println(anything)
}
