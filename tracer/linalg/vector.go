package linalg

import "fmt"

type Vector quadruple

func NewVector(x, y, z float64) *Vector {
	return &Vector{x, y, z, 0}
}

func (v *Vector) Add(other *Vector) *Vector {
	return (*Vector)((*quadruple)(v).Add((*quadruple)(other)))
}

func (v *Vector) Sub(other *Vector) *Vector {
	return (*Vector)((*quadruple)(v).Sub((*quadruple)(other)))
}

func (v *Vector) Mul(scalar float64) *Vector {
	return (*Vector)((*quadruple)(v).Mul(scalar))
}

func (v *Vector) Div(scalar float64) *Vector {
	return (*Vector)((*quadruple)(v).Div(scalar))
}

func (v *Vector) Neg() *Vector {
	return (*Vector)((*quadruple)(v).Neg())
}

func (v *Vector) Dot(other *Vector) float64 {
	return (*quadruple)(v).Dot((*quadruple)(other))
}

func (v *Vector) Cross(other *Vector) *Vector {
	return &Vector{
		v.y*other.z - v.z*other.y,
		v.z*other.x - v.x*other.z,
		v.x*other.y - v.y*other.x,
		0,
	}
}

func (v *Vector) Magnitude() float64 {
	return (*quadruple)(v).Magnitude()
}

func (v *Vector) Normalized() *Vector {
	return (*Vector)((*quadruple)(v).Normalized())
}

func (v *Vector) Equal(other *Vector) bool {
	return (*quadruple)(v).Equal((*quadruple)(other))
}

func (v *Vector) String() string {
	return fmt.Sprintf("V(%g, %g, %g)", v.x, v.y, v.z)
}
