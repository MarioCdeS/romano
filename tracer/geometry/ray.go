package geometry

import (
	"github.com/MarioCdeS/romano/tracer/linalg/point"
	"github.com/MarioCdeS/romano/tracer/linalg/vector"
)

type Ray struct {
	Origin    point.Point
	direction vector.Vector
}

func NewRay(origin point.Point, direction vector.Vector) *Ray {
	return &Ray{origin, direction.Normalized()}
}

func (r *Ray) Direction() vector.Vector {
	return r.direction
}

func (r *Ray) SetDirection(direction *vector.Vector) {
	r.direction = *direction.Normalized()
}

func (r *Ray) Position(t float64) *point.Point {
	return r.Origin.AddVector(r.direction.Scale(t))
}
