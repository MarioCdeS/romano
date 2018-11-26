package geometry

import (
	"github.com/MarioCdeS/romano/tracer/ray"
)

type Geometry interface {
	Intersections(ray *ray.Ray) []float64
}
