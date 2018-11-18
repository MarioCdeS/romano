package linalg

import (
	"math"

	"github.com/MarioCdeS/romano/tracer"
)

type Quadruple [4]float64

func (q *Quadruple) Add(oth *Quadruple) *Quadruple {
	return &Quadruple{
		q[0] + oth[0],
		q[1] + oth[1],
		q[2] + oth[2],
		q[3] + oth[3],
	}
}

func (q *Quadruple) Sub(oth *Quadruple) *Quadruple {
	return &Quadruple{
		q[0] - oth[0],
		q[1] - oth[1],
		q[2] - oth[2],
		q[3] - oth[3],
	}
}

func (q *Quadruple) Mul(scalar float64) *Quadruple {
	return &Quadruple{
		q[0] * scalar,
		q[1] * scalar,
		q[2] * scalar,
		q[3] * scalar,
	}
}

func (q *Quadruple) Div(scalar float64) *Quadruple {
	if scalar == 0 {
		panic("divide by zero scalar")
	}

	return &Quadruple{
		q[0] / scalar,
		q[1] / scalar,
		q[2] / scalar,
		q[3] / scalar,
	}
}

func (q *Quadruple) Neg() *Quadruple {
	return &Quadruple{-q[0], -q[1], -q[2], -q[3]}
}

func (q *Quadruple) Dot(oth *Quadruple) float64 {
	return q[0]*oth[0] + q[1]*oth[1] + q[2]*oth[2] + q[3]*oth[3]
}

func (q *Quadruple) Hadamard(oth *Quadruple) *Quadruple {
	return &Quadruple{
		q[0] * oth[0],
		q[1] * oth[1],
		q[2] * oth[2],
		q[3] * oth[3],
	}
}

func (q *Quadruple) Magnitude() float64 {
	return math.Sqrt(q[0]*q[0] + q[1]*q[1] + q[2]*q[2] + q[3]*q[3])
}

func (q *Quadruple) Normalized() *Quadruple {
	return q.Div(q.Magnitude())
}

func (q *Quadruple) Equal(oth *Quadruple) bool {
	return tracer.Equalf64(q[0], oth[0]) &&
		tracer.Equalf64(q[1], oth[1]) &&
		tracer.Equalf64(q[2], oth[2]) &&
		tracer.Equalf64(q[3], oth[3])
}
