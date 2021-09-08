package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
