package linalg

import "fmt"

type Point quadruple

func NewPoint(x, y, z float64) *Point {
	return &Point{x, y, z, 1}
}

func (p *Point) Add(v *Vector) *Point {
	return (*Point)((*quadruple)(p).Add((*quadruple)(v)))
}

func (p *Point) SubVector(v *Vector) *Point {
	return (*Point)((*quadruple)(p).Sub((*quadruple)(v)))
}

func (p *Point) SubPoint(other *Point) *Vector {
	return (*Vector)((*quadruple)(p).Sub((*quadruple)(other)))
}

func (p *Point) Equal(other *Point) bool {
	return (*quadruple)(p).Equal((*quadruple)(other))
}

func (p *Point) String() string {
	return fmt.Sprintf("P(%g, %g, %g)", p.x, p.y, p.z)
}
