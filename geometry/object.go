package geometry

import "github.com/MarioCdeS/romano/linalg"

type Object interface {
	Material() *Material
	Intersections(r *Ray) []Intersection
	NormalAt(p *linalg.Point) *linalg.Vector
}
