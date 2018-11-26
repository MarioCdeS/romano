package point

import (
	"fmt"

	"github.com/MarioCdeS/romano/tracer/float"
	"github.com/MarioCdeS/romano/tracer/linalg/vector"
)

type Point struct {
	X, Y, Z float64
}

func New(x, y, z float64) *Point {
	return &Point{x, y, z}
}

func (p *Point) Copy() *Point {
	res := *p
	return &res
}

func (p *Point) W() float64 {
	return 1
}

func (p *Point) AddVector(v *vector.Vector) *Point {
	return p.Copy().MutAddVector(v)
}

func (p *Point) MutAddVector(v *vector.Vector) *Point {
	p.X += v.X
	p.Y += v.Y
	p.Z += v.Z

	return p
}

func (p *Point) SubVector(v *vector.Vector) *Point {
	return p.Copy().MutSubVector(v)
}

func (p *Point) MutSubVector(v *vector.Vector) *Point {
	p.X -= v.X
	p.Y -= v.Y
	p.Z -= v.Z

	return p
}

func (p *Point) SubPoint(oth *Point) *vector.Vector {
	return vector.New(p.X-oth.X, p.Y-oth.Y, p.Z-oth.Z)
}

func (p *Point) Equal(oth *Point) bool {
	return float.ApproxEqual(p.X, oth.X) &&
		float.ApproxEqual(p.Y, oth.Y) &&
		float.ApproxEqual(p.Z, oth.Z)
}

func (p *Point) String() string {
	return fmt.Sprintf("P(%g, %g, %g)", p.X, p.Y, p.Z)
}
