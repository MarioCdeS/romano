package sphere

import (
	"math"

	"github.com/MarioCdeS/romano/tracer/float"
	"github.com/MarioCdeS/romano/tracer/linalg/point"
	"github.com/MarioCdeS/romano/tracer/ray"
)

type Sphere struct {
	center point.Point
	radius float64
}

func New(center *point.Point, radius float64) *Sphere {
	return &Sphere{*center, radius}
}

func (s *Sphere) Intersections(ray *ray.Ray) []float64 {
	centerToRayOrig := ray.Origin.SubPoint(&s.center)
	// a = 1 in this case, because ray direction is a unit vector
	b := 2 * ray.Direction.Dot(centerToRayOrig)
	c := centerToRayOrig.Dot(centerToRayOrig) - s.radius*s.radius
	disc := b*b - 4*c

	if disc < 0 {
		return nil
	}

	res := make([]float64, 1, 2)
	sqrtDisc := math.Sqrt(disc)
	res[0] = (-b - sqrtDisc) / 2

	if !float.ApproxEqual(disc, 0) {
		res = append(res, (-b+sqrtDisc)/2)
	}

	return res
}
