package main

import "math"

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	ans := math.Pi * r.width * r.height
	return ans
}

func (r rect) perimeter() float64 {
	ans := 2 * (r.width + r.height)
	return ans
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	ans := math.Pi * c.radius * c.radius
	return ans
}

func (c circle) perimeter() float64 {
	ans := 2 * math.Pi * c.radius
	return ans
}
