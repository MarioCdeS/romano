package geometry

type Geometry interface {
	Intersections(ray Ray) []float64
}
