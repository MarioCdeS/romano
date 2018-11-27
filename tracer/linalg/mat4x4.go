package linalg

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MarioCdeS/romano/tracer/float"
)

type Mat4x4 [4][4]float64

func NewMat4x4(elems ...float64) *Mat4x4 {
	if len(elems) > 16 {
		panic("too many elements given")
	}

	var mat Mat4x4

	for i, e := range elems {
		mat[i/4][i%4] = e
	}

	return &mat
}

func NewMat4x4ID() *Mat4x4 {
	return &Mat4x4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func (m *Mat4x4) Copy() *Mat4x4 {
	res := *m
	return &res
}

func (m *Mat4x4) Add(oth *Mat4x4) *Mat4x4 {
	return m.Copy().MutAdd(oth)
}

func (m *Mat4x4) MutAdd(oth *Mat4x4) *Mat4x4 {
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
	return m.Copy().MutScale(scalar)
}

func (m *Mat4x4) MutScale(scalar float64) *Mat4x4 {
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
	return m.Copy().MutT()
}

func (m *Mat4x4) MutT() *Mat4x4 {
	m[0][1], m[1][0] = m[1][0], m[0][1]
	m[0][2], m[2][0] = m[2][0], m[0][2]
	m[0][3], m[3][0] = m[3][0], m[0][3]
	m[1][2], m[2][1] = m[2][1], m[1][2]
	m[1][3], m[3][1] = m[3][1], m[1][3]
	m[2][3], m[3][2] = m[3][2], m[2][3]

	return m
}

func (m *Mat4x4) Dot(oth *Mat4x4) *Mat4x4 {
	var res Mat4x4

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

	return &res
}

func (m *Mat4x4) DotVector(v Vector) Vector {
	return Vector{
		X: m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z,
		Y: m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z,
		Z: m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z,
	}
}

func (m *Mat4x4) DotPoint(p Point) Point {
	return Point{
		X: m[0][0]*p.X + m[0][1]*p.Y + m[0][2]*p.Z + m[0][3],
		Y: m[1][0]*p.X + m[1][1]*p.Y + m[1][2]*p.Z + m[1][3],
		Z: m[2][0]*p.X + m[2][1]*p.Y + m[2][2]*p.Z + m[2][3],
	}
}

func (m *Mat4x4) SubMat(i, j int) *Mat3x3 {
	if i < 0 || i >= 4 {
		panic("row index out of bounds")
	}

	if j < 0 || j >= 4 {
		panic("column index out of bounds")
	}

	var res Mat3x3

	for di, si := 0, 0; di < 3; di, si = di+1, si+1 {
		if si == i {
			si++
		}

		for dj, sj := 0, 0; dj < 3; dj, sj = dj+1, sj+1 {
			if sj == j {
				sj++
			}

			res[di][dj] = m[si][sj]
		}
	}

	return &res
}

func (m *Mat4x4) Minor(i, j int) float64 {
	return m.SubMat(i, j).Det()
}

func (m *Mat4x4) Cofactor(i, j int) float64 {
	minor := m.Minor(i, j)

	if (i+j)%2 == 0 {
		return minor
	} else {
		return -minor
	}
}

func (m *Mat4x4) Det() float64 {
	return m[0][0]*m.Cofactor(0, 0) + m[0][1]*m.Cofactor(0, 1) + m[0][2]*m.Cofactor(0, 2) + m[0][3]*m.Cofactor(0, 3)
}

func (m *Mat4x4) Inv() (*Mat4x4, error) {
	det := m.Det()

	if det == 0 {
		return nil, errors.New("matrix is not invertible")
	}

	var res Mat4x4

	// Cofactors, transpose, and division by determinant in one
	res[0][0] = m.Cofactor(0, 0) / det
	res[0][1] = m.Cofactor(1, 0) / det
	res[0][2] = m.Cofactor(2, 0) / det
	res[0][3] = m.Cofactor(3, 0) / det
	res[1][0] = m.Cofactor(0, 1) / det
	res[1][1] = m.Cofactor(1, 1) / det
	res[1][2] = m.Cofactor(2, 1) / det
	res[1][3] = m.Cofactor(3, 1) / det
	res[2][0] = m.Cofactor(0, 2) / det
	res[2][1] = m.Cofactor(1, 2) / det
	res[2][2] = m.Cofactor(2, 2) / det
	res[2][3] = m.Cofactor(3, 2) / det
	res[3][0] = m.Cofactor(0, 3) / det
	res[3][1] = m.Cofactor(1, 3) / det
	res[3][2] = m.Cofactor(2, 3) / det
	res[3][3] = m.Cofactor(3, 3) / det

	return &res, nil
}

func (m *Mat4x4) Equal(oth *Mat4x4) bool {
	return float.ApproxEqual(m[0][0], oth[0][0]) &&
		float.ApproxEqual(m[0][1], oth[0][1]) &&
		float.ApproxEqual(m[0][2], oth[0][2]) &&
		float.ApproxEqual(m[0][3], oth[0][3]) &&
		float.ApproxEqual(m[1][0], oth[1][0]) &&
		float.ApproxEqual(m[1][1], oth[1][1]) &&
		float.ApproxEqual(m[1][2], oth[1][2]) &&
		float.ApproxEqual(m[1][3], oth[1][3]) &&
		float.ApproxEqual(m[2][0], oth[2][0]) &&
		float.ApproxEqual(m[2][1], oth[2][1]) &&
		float.ApproxEqual(m[2][2], oth[2][2]) &&
		float.ApproxEqual(m[2][3], oth[2][3]) &&
		float.ApproxEqual(m[3][0], oth[3][0]) &&
		float.ApproxEqual(m[3][1], oth[3][1]) &&
		float.ApproxEqual(m[3][2], oth[3][2]) &&
		float.ApproxEqual(m[3][3], oth[3][3])
}

func (m *Mat4x4) String() string {
	sb := strings.Builder{}

	for i := 0; i < 4; i++ {
		sb.WriteString(fmt.Sprintf("|%16.5f %16.5f %16.5f %16.5f|\n", m[i][0], m[i][1], m[i][2], m[i][3]))
	}

	return sb.String()
}
