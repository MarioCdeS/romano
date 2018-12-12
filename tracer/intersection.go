package tracer

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
