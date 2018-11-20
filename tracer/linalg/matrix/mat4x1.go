package matrix

import (
	"fmt"
	"strings"

	"github.com/MarioCdeS/romano/tracer"
)

type Mat4x1 [4]float64

func New4x1(elems ...float64) *Mat4x1 {
	if len(elems) > 4 {
		panic("too many elements given")
	}

	var mat Mat4x1

	for i, e := range elems {
		mat[i] = e
	}

	return &mat
}

func (m *Mat4x1) Copy() *Mat4x1 {
	res := *m
	return &res
}

func (m *Mat4x1) Add(oth *Mat4x1) *Mat4x1 {
	m[0] += oth[0]
	m[1] += oth[1]
	m[2] += oth[2]
	m[3] += oth[3]

	return m
}

func (m *Mat4x1) T() *Mat1x4 {
	return (*Mat1x4)(m)
}

func (m *Mat4x1) Equal(oth *Mat4x1) bool {
	return tracer.Equalf64(m[0], oth[0]) &&
		tracer.Equalf64(m[1], oth[1]) &&
		tracer.Equalf64(m[2], oth[2]) &&
		tracer.Equalf64(m[3], oth[3])
}

func (m *Mat4x1) String() string {
	sb := strings.Builder{}

	for i := 0; i < 4; i++ {
		sb.WriteString(fmt.Sprintf("|%16.5f|\n", m[i]))
	}

	return sb.String()
}
