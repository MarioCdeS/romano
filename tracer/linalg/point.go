package linalg

import "fmt"

type Point Quadruple

func NewPoint(x, y, z float64) *Point {
	return &Point{x, y, z, 1}
}

func (p *Point) X() float64 {
	return p[0]
}

func (p *Point) Y() float64 {
	return p[1]
}

func (p *Point) Z() float64 {
	return p[2]
}

func (p *Point) Add(v *Vector) *Point {
	return (*Point)((*Quadruple)(p).Add((*Quadruple)(v)))
}

func (p *Point) SubVector(v *Vector) *Point {
	return (*Point)((*Quadruple)(p).Sub((*Quadruple)(v)))
}

func (p *Point) SubPoint(oth *Point) *Vector {
	return (*Vector)((*Quadruple)(p).Sub((*Quadruple)(oth)))
}

func (p *Point) Equal(oth *Point) bool {
	return (*Quadruple)(p).Equal((*Quadruple)(oth))
}

func (p *Point) String() string {
	return fmt.Sprintf("P(%g, %g, %g)", p[0], p[1], p[2])
}
