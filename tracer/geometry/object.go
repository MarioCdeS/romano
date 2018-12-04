package geometry

import "github.com/MarioCdeS/romano/tracer/linalg"

type Object interface {
	Intersections(r *Ray) []Intersection
	NormalAt(p *linalg.Point) *linalg.Vector
}
