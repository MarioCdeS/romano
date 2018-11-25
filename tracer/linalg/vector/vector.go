package vector

import (
	"fmt"
	"math"

	"github.com/MarioCdeS/romano/tracer/linalg/matrix"
)

type Vector matrix.Mat4x1

func New(x, y, z float64) *Vector {
	return &Vector{x, y, z, 0}
}

func From4x1(m *matrix.Mat4x1) *Vector {
	return (*Vector)(m)
}

func (v *Vector) Copy() *Vector {
	res := *v
	return &res
}

func (v *Vector) X() float64 {
	return v[0]
}

func (v *Vector) Y() float64 {
	return v[1]
}

func (v *Vector) Z() float64 {
	return v[2]
}

func (v *Vector) Add(oth *Vector) *Vector {
	return (*Vector)((*matrix.Mat4x1)(v).Add((*matrix.Mat4x1)(oth)))
}

func (v *Vector) Sub(oth *Vector) *Vector {
	return (*Vector)((*matrix.Mat4x1)(v).Sub((*matrix.Mat4x1)(oth)))
}

func (v *Vector) Mul(scalar float64) *Vector {
	return (*Vector)((*matrix.Mat4x1)(v).Scale(scalar))
}

func (v *Vector) Div(scalar float64) *Vector {
	return (*Vector)((*matrix.Mat4x1)(v).Scale(1.0 / scalar))
}

func (v *Vector) Neg() *Vector {
	return (*Vector)((*matrix.Mat4x1)(v).Scale(-1))
}

func (v *Vector) Dot(oth *Vector) float64 {
	return (*matrix.Mat1x4)(v).Dot((*matrix.Mat4x1)(oth))
}

func (v *Vector) Cross(oth *Vector) *Vector {
	return &Vector{
		v[1]*oth[2] - v[2]*oth[1],
		v[2]*oth[0] - v[0]*oth[2],
		v[0]*oth[1] - v[1]*oth[0],
		0,
	}
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

func (v *Vector) Normalized() *Vector {
	return v.Div(v.Magnitude())
}

func (v *Vector) AsMat4x1() *matrix.Mat4x1 {
	return (*matrix.Mat4x1)(v)
}

func (v *Vector) Equal(oth *Vector) bool {
	return (*matrix.Mat4x1)(v).Equal((*matrix.Mat4x1)(oth))
}

func (v *Vector) String() string {
	return fmt.Sprintf("V(%g, %g, %g)", v[0], v[1], v[2])
}
