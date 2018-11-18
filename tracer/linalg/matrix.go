package linalg

import (
	"github.com/MarioCdeS/romano/tracer"
)

type Matrix [][]float64

func NewMatrix(rows, cols int, elems ...float64) Matrix {
	if len(elems) > rows*cols {
		panic("too many elements given")
	}

	mat := make(Matrix, rows)

	for i := 0; i < rows; i++ {
		mat[i] = make([]float64, cols)
	}

	for i, e := range elems {
		mat[i/4][i%4] = e
	}

	return mat
}

func NewMatrix4x4(elems ...float64) Matrix {
	return NewMatrix(4, 4, elems...)
}

func NewMatrix3x3(elems ...float64) Matrix {
	return NewMatrix(3, 3, elems...)
}

func NewMatrix2x2(elems ...float64) Matrix {
	return NewMatrix(2, 2, elems...)
}

func NewColumnMatrix4(elems ...float64) Matrix {
	return NewMatrix(4, 1, elems...)
}

func (m Matrix) Dimensions() (int, int) {
	rows := len(m)
	var cols int

	if rows > 0 {
		cols = len(m[0])
	}

	return rows, cols
}

func (m Matrix) Add(oth Matrix) Matrix {
	if !m.EqualDim(oth) {
		panic("dimensions of matrices are not equal")
	}

	return m.Map(func(elem float64, row, col int) float64 {
		return elem + oth[row][col]
	})
}

func (m Matrix) Sub(oth Matrix) Matrix {
	if !m.EqualDim(oth) {
		panic("dimensions of matrices are not equal")
	}

	return m.Map(func(elem float64, row, col int) float64 {
		return elem - oth[row][col]
	})
}

func (m Matrix) Neg() Matrix {
	return m.Map(func(elem float64, _, _ int) float64 {
		return -elem
	})
}

func (m Matrix) Scale(scalar float64) Matrix {
	return m.Map(func(elem float64, _, _ int) float64 {
		return elem * scalar
	})
}

func (m Matrix) Dot(oth Matrix) Matrix {
	numRows, numCols := m.Dimensions()
	othRows, othCols := oth.Dimensions()

	if numCols != othRows {
		panic("dimensions of matrices are incompatible")
	}

	res := NewMatrix(numRows, othCols)

	for row := 0; row < numRows; row++ {
		for col := 0; col < othCols; col++ {
			var elem float64

			for i := 0; i < numCols; i++ {
				elem += m[row][i] * oth[i][col]
			}

			res[row][col] = elem
		}
	}

	return res
}

func (m Matrix) Equal(oth Matrix) bool {
	numRows, numCols := m.Dimensions()

	if othRows, othCols := oth.Dimensions(); !(numRows == othRows && numCols == othCols) {
		return false
	}

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if !tracer.Equalf64(m[row][col], oth[row][col]) {
				return false
			}
		}
	}

	return true
}

func (m Matrix) EqualDim(oth Matrix) bool {
	numRows, numCols := m.Dimensions()
	othRows, othCols := oth.Dimensions()
	return numRows == othRows && numCols == othCols
}

func (m Matrix) Map(proj func(elem float64, row, col int) float64) Matrix {
	res := NewMatrix(m.Dimensions())

	for i, row := range m {
		for j, elem := range row {
			res[i][j] = proj(elem, i, j)
		}
	}

	return res
}
