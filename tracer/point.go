package tracer

import "fmt"

type Point struct {
	X, Y, Z float64
}

func Origin() *Point {
	return &Point{0, 0, 0}
}

func (p Point) AddVector(v *Vector) *Point {
	return p.MutAddVector(v)
}

func (p *Point) MutAddVector(v *Vector) *Point {
	p.X += v.X
	p.Y += v.Y
	p.Z += v.Z

	return p
}

func (p Point) SubVector(v *Vector) *Point {
	return p.MutSubVector(v)
}

func (p *Point) MutSubVector(v *Vector) *Point {
	p.X -= v.X
	p.Y -= v.Y
	p.Z -= v.Z

	return p
}

func (p *Point) SubPoint(oth *Point) *Vector {
	return &Vector{p.X - oth.X, p.Y - oth.Y, p.Z - oth.Z}
}

func (p *Point) Equal(oth *Point) bool {
	return ApproxEqual(p.X, oth.X) &&
		ApproxEqual(p.Y, oth.Y) &&
		ApproxEqual(p.Z, oth.Z)
}

func (p *Point) String() string {
	return fmt.Sprintf("P(%g, %g, %g)", p.X, p.Y, p.Z)
}