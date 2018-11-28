package geometry

import "github.com/MarioCdeS/romano/tracer/linalg"

type Ray struct {
	Origin    linalg.Point
	Direction linalg.Vector
}

func NewRay(origin linalg.Point, direction linalg.Vector) *Ray {
	return &Ray{origin, direction}
}

func (r *Ray) Position(t float64) linalg.Point {
	return r.Origin.AddVector(r.Direction.Scale(t))
}

func (r *Ray) Transform(t *linalg.Mat4x4) *Ray {
	return &Ray{
		Origin:    t.DotPoint(r.Origin),
		Direction: t.DotVector(r.Direction),
	}
}
