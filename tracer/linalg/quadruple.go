package linalg

import (
	"math"

	"github.com/MarioCdeS/romano/tracer"
)

type quadruple [4]float64

func (q *quadruple) Add(oth *quadruple) *quadruple {
	return &quadruple{
		q[0] + oth[0],
		q[1] + oth[1],
		q[2] + oth[2],
		q[3] + oth[3],
	}
}

func (q *quadruple) Sub(oth *quadruple) *quadruple {
	return &quadruple{
		q[0] - oth[0],
		q[1] - oth[1],
		q[2] - oth[2],
		q[3] - oth[3],
	}
}

func (q *quadruple) Mul(scalar float64) *quadruple {
	return &quadruple{
		q[0] * scalar,
		q[1] * scalar,
		q[2] * scalar,
		q[3] * scalar,
	}
}

func (q *quadruple) Div(scalar float64) *quadruple {
	if scalar == 0 {
		panic("divide by zero scalar")
	}

	return &quadruple{
		q[0] / scalar,
		q[1] / scalar,
		q[2] / scalar,
		q[3] / scalar,
	}
}

func (q *quadruple) Neg() *quadruple {
	return &quadruple{-q[0], -q[1], -q[2], -q[3]}
}

func (q *quadruple) Dot(oth *quadruple) float64 {
	return q[0]*oth[0] + q[1]*oth[1] + q[2]*oth[2] + q[3]*oth[3]
}

func (q *quadruple) Hadamard(oth *quadruple) *quadruple {
	return &quadruple{
		q[0] * oth[0],
		q[1] * oth[1],
		q[2] * oth[2],
		q[3] * oth[3],
	}
}

func (q *quadruple) Magnitude() float64 {
	return math.Sqrt(q[0]*q[0] + q[1]*q[1] + q[2]*q[2] + q[3]*q[3])
}

func (q *quadruple) Normalized() *quadruple {
	return q.Div(q.Magnitude())
}

func (q *quadruple) Equal(oth *quadruple) bool {
	return tracer.Equalf64(q[0], oth[0]) &&
		tracer.Equalf64(q[1], oth[1]) &&
		tracer.Equalf64(q[2], oth[2]) &&
		tracer.Equalf64(q[3], oth[3])
}
