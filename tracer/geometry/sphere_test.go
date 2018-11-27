package geometry

import (
	"testing"

	"github.com/MarioCdeS/romano/tracer/float"
	"github.com/MarioCdeS/romano/tracer/linalg"
)

func TestSphere_Intersections(t *testing.T) {
	r := NewRay(linalg.Point{0, 1, -5}, linalg.Vector{0, 0, 1})
	s := NewSphere(linalg.Point{0, 0, 0}, 1)
	is := s.Intersections(r)

	if len(is) != 1 {
		t.Errorf("unexpected number of intersections (%d)", len(is))
		return
	}

	if !float.ApproxEqual(5, is[0]) {
		t.Errorf("expected 5, got %g", is[0])
	}
}

func TestSphere_Intersections2(t *testing.T) {
	r := NewRay(linalg.Point{0, 0, -5}, linalg.Vector{0, 0, 1})
	s := NewSphere(linalg.Point{0, 0, 0}, 1)
	is := s.Intersections(r)

	if len(is) != 2 {
		t.Errorf("unexpected number of intersections (%d)", len(is))
		return
	}

	if !float.ApproxEqual(4, is[0]) {
		t.Errorf("expected 4 for first, got %g", is[0])
	}

	if !float.ApproxEqual(6, is[1]) {
		t.Errorf("expected 6 for second, got %g", is[1])
	}
}

func TestSphere_Intersections3(t *testing.T) {
	r := NewRay(linalg.Point{0, 2, -5}, linalg.Vector{0, 0, 1})
	s := NewSphere(linalg.Point{0, 0, 0}, 1)
	is := s.Intersections(r)

	if len(is) > 0 {
		t.Errorf("unexpected number of intersections (%d)", len(is))
		return
	}
}

func TestSphere_Intersections4(t *testing.T) {
	r := NewRay(linalg.Point{0, 0, 0}, linalg.Vector{0, 0, 1})
	s := NewSphere(linalg.Point{0, 0, 0}, 1)
	is := s.Intersections(r)

	if len(is) != 2 {
		t.Errorf("unexpected number of intersections (%d)", len(is))
		return
	}

	if !float.ApproxEqual(-1, is[0]) {
		t.Errorf("expected -1 for first, got %g", is[0])
	}

	if !float.ApproxEqual(1, is[1]) {
		t.Errorf("expected 1 for second, got %g", is[1])
	}
}

func TestSphere_Intersections5(t *testing.T) {
	r := NewRay(linalg.Point{0, 0, 5}, linalg.Vector{0, 0, 1})
	s := NewSphere(linalg.Point{0, 0, 0}, 1)
	is := s.Intersections(r)

	if len(is) != 2 {
		t.Errorf("unexpected number of intersections (%d)", len(is))
		return
	}

	if !float.ApproxEqual(-6, is[0]) {
		t.Errorf("expected -6 for first, got %g", is[0])
	}

	if !float.ApproxEqual(-4, is[1]) {
		t.Errorf("expected -4 for second, got %g", is[1])
	}
}
