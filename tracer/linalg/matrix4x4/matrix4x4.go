package matrix4x4

type Matrix4x4 struct {
	elems      [4][4]float64
	transposed bool
}

func New(elems ...float64) *Matrix4x4 {
	if len(elems) > 16 {
		panic("too many elements given")
	}

	var mat Matrix4x4

	for i, e := range elems {
		mat.elems[i/4][i%4] = e
	}

	return &mat
}

func Add(a *Matrix4x4, b *Matrix4x4) *Matrix4x4 {
	res := *a
	res.Add(b)
	return &res
}

func Transpose(m *Matrix4x4) *Matrix4x4 {
	res := *m
	res.Transpose()
	return &res
}

func (m *Matrix4x4) At(row, col int) float64 {
	row, col = elemIndices(row, col, m.transposed)
	return m.elems[row][col]
}

func (m *Matrix4x4) Set(row, col int, value float64) {
	row, col = elemIndices(row, col, m.transposed)
	m.elems[row][col] = value
}

func (m *Matrix4x4) Add(oth *Matrix4x4) {
	m.elems[0][0] += oth.elems[0][0]
	m.elems[0][1] += oth.elems[0][1]
	m.elems[0][2] += oth.elems[0][2]
	m.elems[0][3] += oth.elems[0][3]
	m.elems[1][0] += oth.elems[1][0]
	m.elems[1][1] += oth.elems[1][1]
	m.elems[1][2] += oth.elems[1][2]
	m.elems[1][3] += oth.elems[1][3]
	m.elems[2][0] += oth.elems[2][0]
	m.elems[2][1] += oth.elems[2][1]
	m.elems[2][2] += oth.elems[2][2]
	m.elems[2][3] += oth.elems[2][3]
	m.elems[3][0] += oth.elems[3][0]
	m.elems[3][1] += oth.elems[3][1]
	m.elems[3][2] += oth.elems[3][2]
	m.elems[3][3] += oth.elems[3][3]
}

func (m *Matrix4x4) Transpose() {
	m.transposed = !m.transposed
}

func elemIndices(row, col int, transposed bool) (int, int) {
	if transposed {
		return col, row
	} else {
		return row, col
	}
}
