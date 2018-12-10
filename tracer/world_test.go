package tracer

import "testing"

func TestWorld_Intersections(t *testing.T) {
	w, err := NewTestWorld()

	if err != nil {
		t.Error(err)
		return
	}

	r := Ray{
		Origin:    Point{0, 0, -5},
		Direction: Vector{0, 0, 1},
	}

	is := w.Intersections(&r)

	if len(is) != 4 {
		t.Errorf("unexpected number of intersections (%d)", len(is))
		return
	}

	ts := [4]float64{4, 4.5, 5.5, 6}

	for i, exp := range ts {
		if !ApproxEqual(exp, is[i].T) {
			t.Errorf("at %d, expected %g, got %g", i, exp, is[i].T)
		}
	}
}

func NewTestWorld() (*World, error) {
	l := PointLight{
		Intensity: Color{1, 1, 1, 1},
		Position:  Point{-10, -10, -10},
	}

	s1 := NewSphere(NewMaterial(&Color{0.8, 1.0, 0.6, 1.0}, 0.1, 0.7, 0.2, 200))
	s2, err := NewTransformedSphere(
		NewUniformScale(0.5),
		NewMaterial(&Color{1, 1, 1, 1}, 0.1, 0.9, 0.9, 200),
	)

	if err != nil {
		return nil, err
	}

	w := World{
		Objects: []Object{s1, s2},
		Lights:  []*PointLight{&l},
	}

	return &w, nil
}
