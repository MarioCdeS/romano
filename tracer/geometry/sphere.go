package geometry

import (
	"math"

	"github.com/MarioCdeS/romano/tracer/float"
)

type Sphere struct{}

func (s Sphere) Intersections(ray *Ray) []Intersection {
	centerToRayOrig := ray.Origin.SubPoint(WorldOrigin) // Sphere center is at origin
	a := ray.Direction.Dot(ray.Direction)
	b := 2 * ray.Direction.Dot(centerToRayOrig)
	c := centerToRayOrig.Dot(centerToRayOrig) - 1 // Sphere radius is 1
	disc := b*b - 4*a*c

	if disc < 0 {
		return nil
	}

	sqrtDisc := math.Sqrt(disc)
	res := make([]Intersection, 1, 2)
	res[0] = NewIntersection((-b-sqrtDisc)/2, s)

	if !float.ApproxEqual(disc, 0) {
		res = append(res, NewIntersection((-b+sqrtDisc)/2, s))
	}

	return res
}
