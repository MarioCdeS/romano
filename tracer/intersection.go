package tracer

import "math"

type Intersection struct {
	T   float64
	obj Object
}

func NewIntersection(t float64, obj Object) *Intersection {
	res := &Intersection{t, nil}
	res.SetObject(obj)
	return res
}

func (i *Intersection) Object() Object {
	return i.obj
}

func (i *Intersection) SetObject(obj Object) {
	if obj == nil {
		panic("intersection object cannot be nil")
	}

	i.obj = obj
}

func Hit(is []*Intersection) (*Intersection, bool) {
	res := &Intersection{math.MaxFloat64, nil}

	for _, i := range is {
		if i.T >= 0 && i.T < res.T {
			res = i
		}
	}

	return res, res.obj != nil
}
