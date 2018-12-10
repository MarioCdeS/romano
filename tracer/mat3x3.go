package tracer

type Mat3x3 [3][3]float64

func NewMat3x3(elems ...float64) *Mat3x3 {
	if len(elems) > 9 {
		panic("too many arguments")
	}

	var mat Mat3x3

	for i, e := range elems {
		mat[i/3][i%3] = e
	}

	return &mat
}

func (m *Mat3x3) Copy() *Mat3x3 {
	res := *m
	return &res
}

func (m *Mat3x3) SubMat(i, j int) *Mat2x2 {
	if i < 0 || i >= 3 {
		panic("row index out of bounds")
	}

	if j < 0 || j >= 3 {
		panic("column index out of bounds")
	}

	res := NewMat2x2()

	for di, si := 0, 0; di < 2; di, si = di+1, si+1 {
		if si == i {
			si++
		}

		for dj, sj := 0, 0; dj < 2; dj, sj = dj+1, sj+1 {
			if sj == j {
				sj++
			}

			res[di][dj] = m[si][sj]
		}
	}

	return res
}

func (m *Mat3x3) Minor(i, j int) float64 {
	return m.SubMat(i, j).Det()
}

func (m *Mat3x3) Cofactor(i, j int) float64 {
	minor := m.Minor(i, j)

	if (i+j)%2 == 0 {
		return minor
	} else {
		return -minor
	}
}

func (m *Mat3x3) Det() float64 {
	return m[0][0]*m.Cofactor(0, 0) + m[0][1]*m.Cofactor(0, 1) + m[0][2]*m.Cofactor(0, 2)
}
