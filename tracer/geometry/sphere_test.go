package geometry

import (
	"testing"

	"github.com/MarioCdeS/romano/tracer/float"
	"github.com/MarioCdeS/romano/tracer/linalg"
)

func TestSphere_Intersections(t *testing.T) {
	r := &Ray{linalg.Point{0, 1, -5}, linalg.Vector{0, 0, 1}}
	s := NewSphere()

	is := s.Intersections(r)

	if len(is) != 1 {
		t.Errorf("unexpected number of intersections (%d)", len(is))
		return
	}

	if !float.ApproxEqual(5, is[0].T) {
		t.Errorf("expected 5, got %g", is[0].T)
	}

	if s != is[0].Object() {
		t.Errorf("intersection object not set")
	}
}

func TestSphere_Intersections2(t *testing.T) {
	r := &Ray{linalg.Point{0, 0, -5}, linalg.Vector{0, 0, 1}}
	s := NewSphere()
	is := s.Intersections(r)

	if len(is) != 2 {
		t.Errorf("unexpected number of intersections (%d)", len(is))
		return
	}

	if !float.ApproxEqual(4, is[0].T) {
		t.Errorf("expected 4 for first, got %g", is[0].T)
	}

	if !float.ApproxEqual(6, is[1].T) {
		t.Errorf("expected 6 for second, got %g", is[1].T)
	}
}

func TestSphere_Intersections3(t *testing.T) {
	r := &Ray{linalg.Point{0, 2, -5}, linalg.Vector{0, 0, 1}}
	s := NewSphere()
	is := s.Intersections(r)

	if len(is) > 0 {
		t.Errorf("unexpected number of intersections (%d)", len(is))
		return
	}
}

func TestSphere_Intersections4(t *testing.T) {
	r := &Ray{linalg.Point{0, 0, 0}, linalg.Vector{0, 0, 1}}
	s := NewSphere()
	is := s.Intersections(r)

	if len(is) != 2 {
		t.Errorf("unexpected number of intersections (%d)", len(is))
		return
	}

	if !float.ApproxEqual(-1, is[0].T) {
		t.Errorf("expected -1 for first, got %g", is[0].T)
	}

	if !float.ApproxEqual(1, is[1].T) {
		t.Errorf("expected 1 for second, got %g", is[1].T)
	}
}

func TestSphere_Intersections5(t *testing.T) {
	r := &Ray{linalg.Point{0, 0, 5}, linalg.Vector{0, 0, 1}}
	s := NewSphere()
	is := s.Intersections(r)

	if len(is) != 2 {
		t.Errorf("unexpected number of intersections (%d)", len(is))
		return
	}

	if !float.ApproxEqual(-6, is[0].T) {
		t.Errorf("expected -6 for first, got %g", is[0].T)
	}

	if !float.ApproxEqual(-4, is[1].T) {
		t.Errorf("expected -4 for second, got %g", is[1].T)
	}
}

func TestSphere_Intersections6(t *testing.T) {
	r := &Ray{linalg.Point{0, 0, -5}, linalg.Vector{0, 0, 1}}
	s, err := NewTransformedSphere(linalg.NewScale(2, 2, 2))

	if err != nil {
		t.Error(err)
		return
	}

	is := s.Intersections(r)

	if len(is) != 2 {
		t.Errorf("unexpected number of intersections (%d)", len(is))
	}

	if !float.ApproxEqual(3, is[0].T) {
		t.Errorf("expected 3 for first, got %g", is[0].T)
	}

	if !float.ApproxEqual(7, is[1].T) {
		t.Errorf("expected 7 for second, got %g", is[1].T)
	}
}

func TestSphere_Intersections7(t *testing.T) {
	r := &Ray{linalg.Point{0, 0, -5}, linalg.Vector{0, 0, 1}}
	s, err := NewTransformedSphere(linalg.NewTranslate(5, 0, 0))

	if err != nil {
		t.Error(err)
		return
	}

	is := s.Intersections(r)

	if len(is) != 0 {
		t.Errorf("unexpected number of intersections (%d)", len(is))
	}
}
