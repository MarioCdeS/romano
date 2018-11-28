package geometry

import (
	"testing"

	"github.com/MarioCdeS/romano/tracer/linalg"
)

func TestHit(t *testing.T) {
	r := NewRay(linalg.Point{0, 0, 0}, linalg.Vector{0, 0, 1})
	var s Sphere
	is := s.Intersections(r)
	hit, ok := Hit(is)

	if !ok {
		t.Error("no hit found")
		return
	}

	if is[1] != hit {
		t.Errorf("expected %v, got %v", is[1], hit)
	}
}

func TestHit2(t *testing.T) {
	r := NewRay(linalg.Point{0, 2, -5}, linalg.Vector{0, 0, 1})
	var s Sphere
	is := s.Intersections(r)
	_, ok := Hit(is)

	if ok {
		t.Error("false hit found")
	}
}
