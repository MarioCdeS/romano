package transform

import (
	"math"

	"github.com/MarioCdeS/romano/tracer/linalg/matrix"
)

func NewTranslate(dx, dy, dz float64) *matrix.Mat4x4 {
	return &matrix.Mat4x4{
		{1, 0, 0, dx},
		{0, 1, 0, dy},
		{0, 0, 1, dz},
		{0, 0, 0, 1},
	}
}

func NewScale(sx, sy, sz float64) *matrix.Mat4x4 {
	return &matrix.Mat4x4{
		{sx, 0, 0, 0},
		{0, sy, 0, 0},
		{0, 0, sz, 0},
		{0, 0, 0, 1},
	}
}

func NewRotateX(rad float64) *matrix.Mat4x4 {
	sin, cos := math.Sincos(rad)

	return &matrix.Mat4x4{
		{1, 0, 0, 0},
		{0, cos, -sin, 0},
		{0, sin, cos, 0},
		{0, 0, 0, 1},
	}
}

func NewRotateY(rad float64) *matrix.Mat4x4 {
	sin, cos := math.Sincos(rad)

	return &matrix.Mat4x4{
		{cos, 0, sin, 0},
		{0, 1, 0, 0},
		{-sin, 0, cos, 0},
		{0, 0, 0, 1},
	}
}

func NewRotateZ(rad float64) *matrix.Mat4x4 {
	sin, cos := math.Sincos(rad)

	return &matrix.Mat4x4{
		{cos, -sin, 0, 0},
		{sin, cos, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func NewShear(xy, xz, yx, yz, zx, zy float64) *matrix.Mat4x4 {
	return &matrix.Mat4x4{
		{1, xy, xz, 0},
		{yx, 1, yz, 0},
		{zx, zy, 1, 0},
		{0, 0, 0, 1},
	}
}
