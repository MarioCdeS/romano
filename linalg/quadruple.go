package linalg

import "math"

type quadruple struct {
	X float64
	Y float64
	Z float64
	w float64
}

func (q *quadruple) add(other *quadruple) *quadruple {
	return &quadruple{
		q.X + other.X,
		q.Y + other.Y,
		q.Z + other.Z,
		q.w + other.w,
	}
}

func (q *quadruple) sub(other *quadruple) *quadruple {
	return &quadruple{
		q.X - other.X,
		q.Y - other.Y,
		q.Z - other.Z,
		q.w - other.w,
	}
}

func (q *quadruple) mul(scalar float64) *quadruple {
	return &quadruple{
		scalar * q.X,
		scalar * q.Y,
		scalar * q.Z,
		scalar * q.w,
	}
}

func (q *quadruple) div(scalar float64) *quadruple {
	if scalar == 0 {
		panic("divide by zero scalar")
	}

	return &quadruple{
		q.X / scalar,
		q.Y / scalar,
		q.Z / scalar,
		q.w / scalar,
	}
}

func (q *quadruple) neg() *quadruple {
	return &quadruple{-q.X, -q.Y, -q.Z, -q.w}
}

func (q *quadruple) dot(other *quadruple) float64 {
	return q.X*other.X + q.Y*other.Y + q.Z*other.Z + q.w*other.w
}

func (q *quadruple) magnitude() float64 {
	return math.Sqrt(q.X*q.X + q.Y*q.Y + q.Z*q.Z + q.w*q.w)
}

func (q *quadruple) normalized() *quadruple {
	return q.div(q.magnitude())
}

func (q *quadruple) equal(other *quadruple) bool {
	return Equalf64(q.X, other.X) && Equalf64(q.Y, other.Y) && Equalf64(q.Z, other.Z) && Equalf64(q.w, other.w)
}
