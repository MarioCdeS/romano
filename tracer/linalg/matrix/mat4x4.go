package matrix

import (
	"fmt"
	"strings"

	"github.com/MarioCdeS/romano/tracer"
)

type Mat4x4 [4][4]float64

func New4x4(elems ...float64) *Mat4x4 {
	if len(elems) > 16 {
		panic("too many elements given")
	}

	var mat Mat4x4

	for i, e := range elems {
		mat[i/4][i%4] = e
	}

	return &mat
}

func (m *Mat4x4) Copy() *Mat4x4 {
	res := *m
	return &res
}

func (m *Mat4x4) Add(oth *Mat4x4) *Mat4x4 {
	m[0][0] += oth[0][0]
	m[0][1] += oth[0][1]
	m[0][2] += oth[0][2]
	m[0][3] += oth[0][3]
	m[1][0] += oth[1][0]
	m[1][1] += oth[1][1]
	m[1][2] += oth[1][2]
	m[1][3] += oth[1][3]
	m[2][0] += oth[2][0]
	m[2][1] += oth[2][1]
	m[2][2] += oth[2][2]
	m[2][3] += oth[2][3]
	m[3][0] += oth[3][0]
	m[3][1] += oth[3][1]
	m[3][2] += oth[3][2]
	m[3][3] += oth[3][3]

	return m
}

func (m *Mat4x4) Scale(scalar float64) *Mat4x4 {
	m[0][0] *= scalar
	m[0][1] *= scalar
	m[0][2] *= scalar
	m[0][3] *= scalar
	m[1][0] *= scalar
	m[1][1] *= scalar
	m[1][2] *= scalar
	m[1][3] *= scalar
	m[2][0] *= scalar
	m[2][1] *= scalar
	m[2][2] *= scalar
	m[2][3] *= scalar
	m[3][0] *= scalar
	m[3][1] *= scalar
	m[3][2] *= scalar
	m[3][3] *= scalar

	return m
}

func (m *Mat4x4) T() *Mat4x4 {
	m[0][1], m[1][0] = m[1][0], m[0][1]
	m[0][2], m[2][0] = m[2][0], m[0][2]
	m[0][3], m[3][0] = m[3][0], m[0][3]
	m[1][2], m[2][1] = m[2][1], m[1][2]
	m[1][3], m[3][1] = m[3][1], m[1][3]
	m[2][3], m[3][2] = m[3][2], m[2][3]

	return m
}

func (m *Mat4x4) Dot(oth *Mat4x4) *Mat4x4 {
	res := New4x4()

	res[0][0] = m[0][0]*oth[0][0] + m[0][1]*oth[1][0] + m[0][2]*oth[2][0] + m[0][3]*oth[3][0]
	res[0][1] = m[0][0]*oth[0][1] + m[0][1]*oth[1][1] + m[0][2]*oth[2][1] + m[0][3]*oth[3][1]
	res[0][2] = m[0][0]*oth[0][2] + m[0][1]*oth[1][2] + m[0][2]*oth[2][2] + m[0][3]*oth[3][2]
	res[0][3] = m[0][0]*oth[0][3] + m[0][1]*oth[1][3] + m[0][2]*oth[2][3] + m[0][3]*oth[3][3]
	res[1][0] = m[1][0]*oth[0][0] + m[1][1]*oth[1][0] + m[1][2]*oth[2][0] + m[1][3]*oth[3][0]
	res[1][1] = m[1][0]*oth[0][1] + m[1][1]*oth[1][1] + m[1][2]*oth[2][1] + m[1][3]*oth[3][1]
	res[1][2] = m[1][0]*oth[0][2] + m[1][1]*oth[1][2] + m[1][2]*oth[2][2] + m[1][3]*oth[3][2]
	res[1][3] = m[1][0]*oth[0][3] + m[1][1]*oth[1][3] + m[1][2]*oth[2][3] + m[1][3]*oth[3][3]
	res[2][0] = m[2][0]*oth[0][0] + m[2][1]*oth[1][0] + m[2][2]*oth[2][0] + m[2][3]*oth[3][0]
	res[2][1] = m[2][0]*oth[0][1] + m[2][1]*oth[1][1] + m[2][2]*oth[2][1] + m[2][3]*oth[3][1]
	res[2][2] = m[2][0]*oth[0][2] + m[2][1]*oth[1][2] + m[2][2]*oth[2][2] + m[2][3]*oth[3][2]
	res[2][3] = m[2][0]*oth[0][3] + m[2][1]*oth[1][3] + m[2][2]*oth[2][3] + m[2][3]*oth[3][3]
	res[3][0] = m[3][0]*oth[0][0] + m[3][1]*oth[1][0] + m[3][2]*oth[2][0] + m[3][3]*oth[3][0]
	res[3][1] = m[3][0]*oth[0][1] + m[3][1]*oth[1][1] + m[3][2]*oth[2][1] + m[3][3]*oth[3][1]
	res[3][2] = m[3][0]*oth[0][2] + m[3][1]*oth[1][2] + m[3][2]*oth[2][2] + m[3][3]*oth[3][2]
	res[3][3] = m[3][0]*oth[0][3] + m[3][1]*oth[1][3] + m[3][2]*oth[2][3] + m[3][3]*oth[3][3]

	return res
}

func (m *Mat4x4) Dot4x1(oth *Mat4x1) *Mat4x1 {
	res := New4x1()

	res[0] = m[0][0]*oth[0] + m[0][1]*oth[1] + m[0][2]*oth[2] + m[0][3]*oth[3]
	res[1] = m[1][0]*oth[0] + m[1][1]*oth[1] + m[1][2]*oth[2] + m[1][3]*oth[3]
	res[2] = m[2][0]*oth[0] + m[2][1]*oth[1] + m[2][2]*oth[2] + m[2][3]*oth[3]
	res[3] = m[3][0]*oth[0] + m[3][1]*oth[1] + m[3][2]*oth[2] + m[3][3]*oth[3]

	return res
}

func (m *Mat4x4) Equal(oth *Mat4x4) bool {
	return tracer.Equalf64(m[0][0], oth[0][0]) &&
		tracer.Equalf64(m[0][1], oth[0][1]) &&
		tracer.Equalf64(m[0][2], oth[0][2]) &&
		tracer.Equalf64(m[0][3], oth[0][3]) &&
		tracer.Equalf64(m[1][0], oth[1][0]) &&
		tracer.Equalf64(m[1][1], oth[1][1]) &&
		tracer.Equalf64(m[1][2], oth[1][2]) &&
		tracer.Equalf64(m[1][3], oth[1][3]) &&
		tracer.Equalf64(m[2][0], oth[2][0]) &&
		tracer.Equalf64(m[2][1], oth[2][1]) &&
		tracer.Equalf64(m[2][2], oth[2][2]) &&
		tracer.Equalf64(m[2][3], oth[2][3]) &&
		tracer.Equalf64(m[3][0], oth[3][0]) &&
		tracer.Equalf64(m[3][1], oth[3][1]) &&
		tracer.Equalf64(m[3][2], oth[3][2]) &&
		tracer.Equalf64(m[3][3], oth[3][3])
}

func (m *Mat4x4) String() string {
	sb := strings.Builder{}

	for i := 0; i < 4; i++ {
		sb.WriteString(fmt.Sprintf("|%16.5f %16.5f %16.5f %16.5f|\n", m[i][0], m[i][1], m[i][2], m[i][3]))
	}

	return sb.String()
}
