package linalg

import "fmt"

type Point quadruple

func NewPoint(x, y, z float64) *Point {
	return &Point{x, y, z, 1}
}

func (p *Point) AddVector(v *Vector) *Point {
	return (*Point)((*quadruple)(p).add((*quadruple)(v)))
}

func (p *Point) SubVector(v *Vector) *Point {
	return (*Point)((*quadruple)(p).sub((*quadruple)(v)))
}

func (p *Point) SubPoint(other *Point) *Vector {
	return (*Vector)((*quadruple)(p).sub((*quadruple)(other)))
}

func (p *Point) Equal(other *Point) bool {
	return (*quadruple)(p).equal((*quadruple)(other))
}

func (p *Point) String() string {
	return fmt.Sprintf("P(%g, %g, %g)", p.X, p.Y, p.Z)
}
