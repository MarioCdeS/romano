package geometry

import "github.com/MarioCdeS/romano/linalg"

type Ray struct {
	Origin    linalg.Point
	Direction linalg.Vector
}

func (r *Ray) PointAt(t float64) *linalg.Point {
	return r.Origin.AddVector(r.Direction.Scale(t))
}

func (r *Ray) Transform(t *linalg.Mat4x4) *Ray {
	return &Ray{
		Origin:    *t.DotPoint(&r.Origin),
		Direction: *t.DotVector(&r.Direction),
	}
}
