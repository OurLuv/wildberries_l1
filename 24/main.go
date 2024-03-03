package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func (p *Point) SetPosition(x, y float64) {
	p.X = x
	p.Y = y
}

func (p *Point) DistanceTo(point Point) float64 {
	return math.Sqrt((point.X-p.X)*(point.X-p.X) + (point.Y-p.Y)*(point.Y-p.Y))
}

func NewPoint() *Point {
	return &Point{
		X: 0,
		Y: 0,
	}
}

func main() {
	p1 := NewPoint()
	p1.SetPosition(5, 5)
	p2 := NewPoint()
	p2.SetPosition(10, 10)

	fmt.Println(p1.DistanceTo(*p2))
}
