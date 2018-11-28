package geometry

type Object interface {
	Intersections(r *Ray) []Intersection
}
