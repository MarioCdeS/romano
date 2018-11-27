package geometry

import "github.com/MarioCdeS/romano/tracer/linalg"

type Ray struct {
	Origin    linalg.Point
	direction linalg.Vector
}

func NewRay(origin linalg.Point, direction linalg.Vector) *Ray {
	return &Ray{origin, direction.Normalized()}
}

func (r *Ray) Direction() linalg.Vector {
	return r.direction
}

func (r *Ray) SetDirection(direction linalg.Vector) {
	r.direction = direction.Normalized()
}

func (r *Ray) Position(t float64) linalg.Point {
	return r.Origin.AddVector(r.direction.Scale(t))
}
