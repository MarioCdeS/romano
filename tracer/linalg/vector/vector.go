package vector

import (
	"fmt"
	"math"

	"github.com/MarioCdeS/romano/tracer/float"
)

type Vector struct {
	X, Y, Z float64
}

func New(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

func (v *Vector) Copy() *Vector {
	res := *v
	return &res
}

func (v *Vector) W() float64 {
	return 0
}

func (v *Vector) Add(oth *Vector) *Vector {
	return v.Copy().MutAdd(oth)
}

func (v *Vector) MutAdd(oth *Vector) *Vector {
	v.X += oth.X
	v.Y += oth.Y
	v.Z += oth.Z

	return v
}

func (v *Vector) Sub(oth *Vector) *Vector {
	return v.Copy().MutSub(oth)
}

func (v *Vector) MutSub(oth *Vector) *Vector {
	v.X -= oth.X
	v.Y -= oth.Y
	v.Z -= oth.Z

	return v
}

func (v *Vector) Scale(scalar float64) *Vector {
	return v.Copy().MutScale(scalar)
}

func (v *Vector) MutScale(scalar float64) *Vector {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar

	return v
}

func (v *Vector) Neg() *Vector {
	return v.Copy().MutNeg()
}

func (v *Vector) MutNeg() *Vector {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z

	return v
}

func (v *Vector) Dot(oth *Vector) float64 {
	return v.X*oth.X + v.Y*oth.Y + v.Z*oth.Z
}

func (v *Vector) Cross(oth *Vector) *Vector {
	return &Vector{
		X: v.Y*oth.Z - v.Z*oth.Y,
		Y: v.Z*oth.X - v.X*oth.Z,
		Z: v.X*oth.Y - v.Y*oth.X,
	}
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector) Normalized() *Vector {
	return v.Copy().MutNormalized()
}

func (v *Vector) MutNormalized() *Vector {
	return v.MutScale(1 / v.Magnitude())
}

func (v *Vector) Equal(oth *Vector) bool {
	return float.ApproxEqual(v.X, oth.X) &&
		float.ApproxEqual(v.Y, oth.Y) &&
		float.ApproxEqual(v.Z, oth.Z)
}

func (v *Vector) String() string {
	return fmt.Sprintf("V(%g, %g, %g)", v.X, v.Y, v.Z)
}
