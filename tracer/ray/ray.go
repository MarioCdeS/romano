package ray

import (
	"github.com/MarioCdeS/romano/tracer/linalg/point"
	"github.com/MarioCdeS/romano/tracer/linalg/vector"
)

type Ray struct {
	Origin    point.Point
	Direction vector.Vector
}

func New(origin *point.Point, direction *vector.Vector) *Ray {
	return &Ray{*origin, *direction.Normalized()}
}

func (r *Ray) Position(t float64) *point.Point {
	return r.Origin.AddVector(r.Direction.Scale(t))
}
