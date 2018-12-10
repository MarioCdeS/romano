package tracer

import "sort"

type World struct {
	Objects []Object
	Lights  []*PointLight
}

type byT []*Intersection

func (w *World) Intersections(ray *Ray) []*Intersection {
	var res []*Intersection

	for _, obj := range w.Objects {
		res = append(res, obj.Intersections(ray)...)
	}

	sort.Sort(byT(res))

	return res
}

func (t byT) Len() int {
	return len(t)
}

func (t byT) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t byT) Less(i, j int) bool {
	return t[i].T < t[j].T
}
