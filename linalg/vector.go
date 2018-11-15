package linalg

import "fmt"

type Vector quadruple

func NewVector(x, y, z float64) *Vector {
	return &Vector{x, y, z, 0}
}

func (v *Vector) Add(other *Vector) *Vector {
	return (*Vector)((*quadruple)(v).add((*quadruple)(other)))
}

func (v *Vector) Sub(other *Vector) *Vector {
	return (*Vector)((*quadruple)(v).sub((*quadruple)(other)))
}

func (v *Vector) Mul(scalar float64) *Vector {
	return (*Vector)((*quadruple)(v).mul(scalar))
}

func (v *Vector) Div(scalar float64) *Vector {
	return (*Vector)((*quadruple)(v).div(scalar))
}

func (v *Vector) Neg() *Vector {
	return (*Vector)((*quadruple)(v).neg())
}

func (v *Vector) Dot(other *Vector) float64 {
	return (*quadruple)(v).dot((*quadruple)(other))
}

func (v *Vector) Cross(other *Vector) *Vector {
	return &Vector{
		v.Y*other.Z - v.Z*other.Y,
		v.Z*other.X - v.X*other.Z,
		v.X*other.Y - v.Y*other.X,
		0,
	}
}

func (v *Vector) Magnitude() float64 {
	return (*quadruple)(v).magnitude()
}

func (v *Vector) Normalized() *Vector {
	return (*Vector)((*quadruple)(v).normalized())
}

func (v *Vector) Equal(other *Vector) bool {
	return (*quadruple)(v).equal((*quadruple)(other))
}

func (v *Vector) String() string {
	return fmt.Sprintf("V(%g, %g, %g)", v.X, v.Y, v.Z)
}
