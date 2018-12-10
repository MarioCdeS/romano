package tracer

type Object interface {
	Material() *Material
	Intersections(r *Ray) []*Intersection
	NormalAt(p *Point) *Vector
}
