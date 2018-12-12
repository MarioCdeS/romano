package tracer

import "math"

type Hit struct {
	Intersection
	Point  *Point
	Normal *Vector
	Camera *Vector
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

	point := ray.Origin.AddVector(ray.Direction.Scale(nearest.T))
	normal := nearest.obj.NormalAt(point)
	camera := ray.Direction.Neg().MutNormalized()

	return &Hit{
		Intersection: *nearest,
		Point:        point,
		Normal:       normal,
		Camera:       camera,
	}, true
}
