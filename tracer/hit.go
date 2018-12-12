package tracer

import "math"

type Hit struct {
	Intersection
	Point  *Point
	Normal *Vector
	Camera *Vector
	Inside bool
}

func FindHit(ray *Ray, is []*Intersection) (*Hit, bool) {
	nearest := &Intersection{math.MaxFloat64, nil}

	for _, i := range is {
		if i.T >= 0 && i.T < nearest.T {
			nearest = i
		}
	}

	if nearest.obj == nil {
		return nil, false
	}

	point := ray.PointAt(nearest.T)
	normal := nearest.obj.NormalAt(point)
	camera := ray.Direction.Neg().MutNormalized()
	var inside bool

	if camera.Dot(normal) < 0 {
		inside = true
		normal.MutNeg()
	}

	return &Hit{
		Intersection: *nearest,
		Point:        point,
		Normal:       normal,
		Camera:       camera,
		Inside:       inside,
	}, true
}
