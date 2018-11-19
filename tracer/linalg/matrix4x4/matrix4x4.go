package matrix4x4

import (
	"fmt"
	"github.com/MarioCdeS/romano/tracer"
	"strings"
)

type Matrix4x4 [4][4]float64

func New(elems ...float64) *Matrix4x4 {
	if len(elems) > 16 {
		panic("too many elements given")
	}

	var mat Matrix4x4

	for i, e := range elems {
		mat[i/4][i%4] = e
	}

	return &mat
}

func Add(a *Matrix4x4, b *Matrix4x4) *Matrix4x4 {
	res := *a
	res.Add(b)
	return &res
}

func Scale(m *Matrix4x4, scalar float64) *Matrix4x4 {
	res := *m
	res.Scale(scalar)
	return &res
}

func Transpose(m *Matrix4x4) *Matrix4x4 {
	res := *m
	res.Transpose()
	return &res
}

func Dot(a *Matrix4x4, b *Matrix4x4) *Matrix4x4 {
	res := New()

	res[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0] + a[0][3]*a[3][0]
	res[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1] + a[0][2]*b[2][1] + a[0][3]*a[3][1]
	res[0][2] = a[0][0]*b[0][2] + a[0][1]*b[1][2] + a[0][2]*b[2][2] + a[0][3]*a[3][2]
	res[0][3] = a[0][0]*b[0][3] + a[0][1]*b[1][3] + a[0][2]*b[2][3] + a[0][3]*a[3][3]
	res[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0] + a[1][3]*a[3][0]
	res[1][1] = a[1][0]*b[0][1] + a[1][1]*b[1][1] + a[1][2]*b[2][1] + a[1][3]*a[3][1]
	res[1][2] = a[1][0]*b[0][2] + a[1][1]*b[1][2] + a[1][2]*b[2][2] + a[1][3]*a[3][2]
	res[1][3] = a[1][0]*b[0][3] + a[1][1]*b[1][3] + a[1][2]*b[2][3] + a[1][3]*a[3][3]
	res[2][0] = a[2][0]*b[0][0] + a[2][1]*b[1][0] + a[2][2]*b[2][0] + a[2][3]*a[3][0]
	res[2][1] = a[2][0]*b[0][1] + a[2][1]*b[1][1] + a[2][2]*b[2][1] + a[2][3]*a[3][1]
	res[2][2] = a[2][0]*b[0][2] + a[2][1]*b[1][2] + a[2][2]*b[2][2] + a[2][3]*a[3][2]
	res[2][3] = a[2][0]*b[0][3] + a[2][1]*b[1][3] + a[2][2]*b[2][3] + a[2][3]*a[3][3]
	res[3][0] = a[3][0]*b[0][0] + a[3][1]*b[1][0] + a[3][2]*b[2][0] + a[3][3]*a[3][0]
	res[3][1] = a[3][0]*b[0][1] + a[3][1]*b[1][1] + a[3][2]*b[2][1] + a[3][3]*a[3][1]
	res[3][2] = a[3][0]*b[0][2] + a[3][1]*b[1][2] + a[3][2]*b[2][2] + a[3][3]*a[3][2]
	res[3][3] = a[3][0]*b[0][3] + a[3][1]*b[1][3] + a[3][2]*b[2][3] + a[3][3]*a[3][3]

	return res
}

func (m *Matrix4x4) Add(oth *Matrix4x4) {
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
}

func (m *Matrix4x4) Scale(scalar float64) {
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
}

func (m *Matrix4x4) Transpose() {
	m[0][1], m[1][0] = m[1][0], m[0][1]
	m[0][2], m[2][0] = m[2][0], m[0][2]
	m[0][3], m[3][0] = m[3][0], m[0][3]
	m[1][2], m[2][1] = m[2][1], m[1][2]
	m[1][3], m[3][1] = m[3][1], m[1][3]
	m[2][3], m[3][2] = m[3][2], m[2][3]
}

func (m *Matrix4x4) Equal(oth *Matrix4x4) bool {
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

func (m *Matrix4x4) String() string {
	sb := strings.Builder{}

	for i := 0; i < 4; i++ {
		sb.WriteString(fmt.Sprintf("|%16.5f %16.5f %16.5f %16.5f|\n", m[i][0], m[i][1], m[i][2], m[i][3]))
	}

	return sb.String()
}
