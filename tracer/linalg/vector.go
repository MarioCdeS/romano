package linalg

import "fmt"

type Vector Quadruple

func NewVector(x, y, z float64) *Vector {
	return &Vector{x, y, z, 0}
}

func (v *Vector) Add(oth *Vector) *Vector {
	return (*Vector)((*Quadruple)(v).Add((*Quadruple)(oth)))
}

func (v *Vector) Sub(oth *Vector) *Vector {
	return (*Vector)((*Quadruple)(v).Sub((*Quadruple)(oth)))
}

func (v *Vector) Mul(scalar float64) *Vector {
	return (*Vector)((*Quadruple)(v).Mul(scalar))
}

func (v *Vector) Div(scalar float64) *Vector {
	return (*Vector)((*Quadruple)(v).Div(scalar))
}

func (v *Vector) Neg() *Vector {
	return (*Vector)((*Quadruple)(v).Neg())
}

func (v *Vector) Dot(oth *Vector) float64 {
	return (*Quadruple)(v).Dot((*Quadruple)(oth))
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
	return (*Quadruple)(v).Magnitude()
}

func (v *Vector) Normalized() *Vector {
	return (*Vector)((*Quadruple)(v).Normalized())
}

func (v *Vector) Equal(oth *Vector) bool {
	return (*Quadruple)(v).Equal((*Quadruple)(oth))
}

func (v *Vector) String() string {
	return fmt.Sprintf("V(%g, %g, %g)", v[0], v[1], v[2])
}
