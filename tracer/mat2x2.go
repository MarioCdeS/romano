package tracer

import "fmt"

type Mat2x2 [2][2]float64

func NewMat2x2(elems ...float64) *Mat2x2 {
	if len(elems) > 4 {
		panic("too many arguments")
	}

	var mat Mat2x2

	for i, e := range elems {
		mat[i/2][i%2] = e
	}

	return &mat
}

func (m *Mat2x2) Copy() *Mat2x2 {
	res := *m
	return &res
}

func (m *Mat2x2) Det() float64 {
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
}

func (m *Mat2x2) Equal(oth *Mat2x2) bool {
	return ApproxEqual(m[0][0], oth[0][0]) &&
		ApproxEqual(m[0][1], oth[0][1]) &&
		ApproxEqual(m[1][0], oth[1][0]) &&
		ApproxEqual(m[1][1], oth[1][1])
}

func (m *Mat2x2) String() string {
	return fmt.Sprintf("|%16.5f %16.5f|\n|%16.5f %16.5f|\n", m[0][0], m[0][1], m[1][0], m[1][1])
}
