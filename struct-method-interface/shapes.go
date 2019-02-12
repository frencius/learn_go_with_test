package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

func (rec Rectangle) Area() float64 {
	return rec.Width * rec.Height
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func (tri Triangle) Area() float64 {
	return (tri.Base * tri.Height) / 2
}
