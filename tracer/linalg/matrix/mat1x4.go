package matrix

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

func Transpose(m *Mat1x4) *Mat4x1 {
	return (*Mat4x1)(m)
}
