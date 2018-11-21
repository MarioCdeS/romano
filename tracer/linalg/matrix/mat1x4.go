package matrix

import (
	"fmt"

	"github.com/MarioCdeS/romano/tracer"
)

type Mat1x4 [4]float64

func New(elems ...float64) *Mat1x4 {
	if len(elems) > 4 {
		panic("too many elements given")
	}

	var mat Mat1x4

	for i, e := range elems {
		mat[i] = e
	}

	return &mat
}

func (m *Mat1x4) Copy() *Mat1x4 {
	res := *m
	return &res
}

func (m *Mat1x4) T() *Mat4x1 {
	return (*Mat4x1)(m)
}

func (m *Mat1x4) Dot(oth *Mat4x1) float64 {
	return m[0]*oth[0] + m[1]*oth[1] + m[2]*oth[2] + m[3]*oth[3]
}

func (m *Mat1x4) Equal(oth *Mat1x4) bool {
	return tracer.ApproxEqual(m[0], oth[0]) &&
		tracer.ApproxEqual(m[1], oth[1]) &&
		tracer.ApproxEqual(m[2], oth[2]) &&
		tracer.ApproxEqual(m[3], oth[3])
}

func (m *Mat1x4) String() string {
	return fmt.Sprintf("|%16.5f %16.5f %16.5f %16.5f|\n", m[0], m[1], m[2], m[3])
}
