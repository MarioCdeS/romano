package linalg

import (
	"math"

	"github.com/MarioCdeS/romano/tracer"
)

type quadruple struct {
	x float64
	y float64
	z float64
	w float64
}

func (q *quadruple) Add(other *quadruple) *quadruple {
	return &quadruple{
		q.x + other.x,
		q.y + other.y,
		q.z + other.z,
		q.w + other.w,
	}
}

func (q *quadruple) Sub(other *quadruple) *quadruple {
	return &quadruple{
		q.x - other.x,
		q.y - other.y,
		q.z - other.z,
		q.w - other.w,
	}
}

func (q *quadruple) Mul(scalar float64) *quadruple {
	return &quadruple{
		q.x * scalar,
		q.y * scalar,
		q.z * scalar,
		q.w * scalar,
	}
}

func (q *quadruple) Div(scalar float64) *quadruple {
	if scalar == 0 {
		panic("divide by zero scalar")
	}

	return &quadruple{
		q.x / scalar,
		q.y / scalar,
		q.z / scalar,
		q.w / scalar,
	}
}

func (q *quadruple) Neg() *quadruple {
	return &quadruple{-q.x, -q.y, -q.z, -q.w}
}

func (q *quadruple) Dot(other *quadruple) float64 {
	return q.x*other.x + q.y*other.y + q.z*other.z + q.w*other.w
}

func (q *quadruple) Magnitude() float64 {
	return math.Sqrt(q.x*q.x + q.y*q.y + q.z*q.z + q.w*q.w)
}

func (q *quadruple) Normalized() *quadruple {
	return q.Div(q.Magnitude())
}

func (q *quadruple) Equal(other *quadruple) bool {
	return tracer.Equalf64(q.x, other.x) &&
		tracer.Equalf64(q.y, other.y) &&
		tracer.Equalf64(q.z, other.z) &&
		tracer.Equalf64(q.w, other.w)
}
