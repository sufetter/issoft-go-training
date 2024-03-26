package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// Перезаписал метод String для красивого вывода
func (p *Point) String() string {
	return fmt.Sprintf("Point {X: %.5f, Y: %.5f}", p.X, p.Y)
}

type PointLabeled struct {
	Point
	Label string
}

func (p PointLabeled) String() string {
	return fmt.Sprintf("PointLabeled %v", p.Point.String())
}

type Normalizable interface {
	Normalize(minX, minY, rangeX, rangeY float64)
	Get() (float64, float64)
}

func (p *Point) Normalize(minX, minY, rangeX, rangeY float64) {
	if p != nil {
		p.X = (p.X - minX) / rangeX
		p.Y = (p.Y - minY) / rangeY
	}
}

func (p *Point) Get() (float64, float64) {
	return p.X, p.Y
}

func Normalize(points []Normalizable) {
	minX, minY, maxX, maxY := math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64
	for _, p := range points {
		if p != nil {
			x, y := p.Get()
			minX = math.Min(minX, x)
			minY = math.Min(minY, y)
			maxX = math.Max(maxX, x)
			maxY = math.Max(maxY, y)
		}
	}
	rangeX := maxX - minX
	rangeY := maxY - minY
	for _, p := range points {
		if p != nil {
			p.Normalize(minX, minY, rangeX, rangeY)
		}

	}
}

func main() {
	points := []Normalizable{
		&Point{X: 1, Y: 2},
		nil,
		&PointLabeled{Point: Point{X: -13, Y: 4}, Label: "Small"},
		&Point{X: 5, Y: 6},
		&PointLabeled{Point: Point{X: 7, Y: 888}, Label: "HUGE!"},
	}
	Normalize(points)
	for _, p := range points {
		fmt.Println(p)
	}
}
