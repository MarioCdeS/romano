package point

import (
	"fmt"

	"github.com/MarioCdeS/romano/tracer/linalg/matrix"
	"github.com/MarioCdeS/romano/tracer/linalg/vector"
)

type Point matrix.Mat4x1

func New(x, y, z float64) *Point {
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

func (p *Point) Add(v *vector.Vector) *Point {
	return (*Point)((*matrix.Mat4x1)(p).Add((*matrix.Mat4x1)(v)))
}

func (p *Point) SubVector(v *vector.Vector) *Point {
	return (*Point)((*matrix.Mat4x1)(p).Sub((*matrix.Mat4x1)(v)))
}

func (p *Point) SubPoint(oth *Point) *vector.Vector {
	return (*vector.Vector)((*matrix.Mat4x1)(p).Sub((*matrix.Mat4x1)(oth)))
}

func (p *Point) Equal(oth *Point) bool {
	return (*matrix.Mat4x1)(p).Equal((*matrix.Mat4x1)(oth))
}

func (p *Point) String() string {
	return fmt.Sprintf("P(%g, %g, %g)", p[0], p[1], p[2])
}