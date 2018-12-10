package tracer

type World struct {
	Objects []Object
	Lights  []*PointLight
}

func (w *World) Intersections(ray *Ray) []Intersection {
	var res []Intersection

	for _, obj := range w.Objects {
		res = append(res, obj.Intersections(ray)...)
	}

	return res
}
