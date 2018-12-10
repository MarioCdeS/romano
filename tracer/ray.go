package tracer

type Ray struct {
	Origin    Point
	Direction Vector
}

func (r *Ray) PointAt(t float64) *Point {
	return r.Origin.AddVector(r.Direction.Scale(t))
}

func (r *Ray) Transform(t *Mat4x4) *Ray {
	return &Ray{
		Origin:    *t.DotPoint(&r.Origin),
		Direction: *t.DotVector(&r.Direction),
	}
}
