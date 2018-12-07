package geometry

import (
	"math"
	"testing"

	"github.com/MarioCdeS/romano/float"
	"github.com/MarioCdeS/romano/linalg"
)

func TestSphere_Intersections(t *testing.T) {
	r := &Ray{linalg.Point{0, 1, -5}, linalg.Vector{0, 0, 1}}
	s := NewSphere(DefaultMaterial())

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
	s := NewSphere(DefaultMaterial())
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
	s := NewSphere(DefaultMaterial())
	is := s.Intersections(r)

	if len(is) > 0 {
		t.Errorf("unexpected number of intersections (%d)", len(is))
		return
	}
}

func TestSphere_Intersections4(t *testing.T) {
	r := &Ray{linalg.Point{0, 0, 0}, linalg.Vector{0, 0, 1}}
	s := NewSphere(DefaultMaterial())
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
	s := NewSphere(DefaultMaterial())
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
	s, err := NewTransformedSphere(linalg.NewScale(2, 2, 2), DefaultMaterial())

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
	s, err := NewTransformedSphere(linalg.NewTranslate(5, 0, 0), DefaultMaterial())

	if err != nil {
		t.Error(err)
		return
	}

	is := s.Intersections(r)

	if len(is) != 0 {
		t.Errorf("unexpected number of intersections (%d)", len(is))
	}
}

func TestSphere_NormalAt(t *testing.T) {
	s := NewSphere(DefaultMaterial())
	exp := linalg.Vector{1, 0, 0}
	got := s.NormalAt(&linalg.Point{1, 0, 0})

	if !exp.Equal(got) {
		t.Errorf("expected %s, got %s", &exp, got)
	}
}

func TestSphere_NormalAt2(t *testing.T) {
	s, err := NewTransformedSphere(linalg.NewTranslate(0, 1, 0), DefaultMaterial())

	if err != nil {
		t.Error(err)
		return
	}

	exp := linalg.Vector{0, 0.7071067811865475, -0.7071067811865476}
	got := s.NormalAt(&linalg.Point{0, 1.70711, -0.70711})

	if !exp.Equal(got) {
		t.Errorf("expected %s, got %s", &exp, got)
	}
}

func TestSphere_NormalAt3(t *testing.T) {
	s, err := NewTransformedSphere(
		linalg.NewScale(1, 0.5, 1).Dot(linalg.NewRotateZ(math.Pi/5)),
		DefaultMaterial(),
	)

	if err != nil {
		t.Error(err)
		return
	}

	exp := linalg.Vector{0, 0.9701425001453319, -0.2425356250363329}
	got := s.NormalAt(&linalg.Point{0, math.Sqrt2 / 2, -math.Sqrt2 / 2})

	if !exp.Equal(got) {
		t.Errorf("expected %s, got %s", &exp, got)
	}
}
