package linalg

import (
	"fmt"

	"github.com/MarioCdeS/romano/tracer/float"
)

type Point struct {
	X, Y, Z float64
}

func (p Point) AddVector(v Vector) Point {
	p.MutAddVector(v)
	return p
}

func (p *Point) MutAddVector(v Vector) *Point {
	p.X += v.X
	p.Y += v.Y
	p.Z += v.Z

	return p
}

func (p Point) SubVector(v Vector) Point {
	p.MutSubVector(v)
	return p
}

func (p *Point) MutSubVector(v Vector) *Point {
	p.X -= v.X
	p.Y -= v.Y
	p.Z -= v.Z

	return p
}

func (p Point) SubPoint(oth Point) Vector {
	return Vector{p.X - oth.X, p.Y - oth.Y, p.Z - oth.Z}
}

func (p Point) Equal(oth Point) bool {
	return float.ApproxEqual(p.X, oth.X) &&
		float.ApproxEqual(p.Y, oth.Y) &&
		float.ApproxEqual(p.Z, oth.Z)
}

func (p Point) String() string {
	return fmt.Sprintf("P(%g, %g, %g)", p.X, p.Y, p.Z)
}
