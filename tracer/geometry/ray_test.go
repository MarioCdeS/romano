package geometry

import (
	"testing"

	"github.com/MarioCdeS/romano/tracer/linalg"
)

func TestRay_Transform(t *testing.T) {
	r := &Ray{linalg.Point{1, 2, 3}, linalg.Vector{0, 1, 0}}
	tl := linalg.NewTranslate(3, 4, 5)
	exp := Ray{linalg.Point{4, 6, 8}, linalg.Vector{0, 1, 0}}
	got := r.Transform(tl)

	if exp != *got {
		t.Errorf("expected %v, got %v", exp, *got)
	}
}

func TestRay_Transform2(t *testing.T) {
	r := &Ray{linalg.Point{1, 2, 3}, linalg.Vector{0, 1, 0}}
	tl := linalg.NewScale(2, 3, 4)
	exp := Ray{linalg.Point{2, 6, 12}, linalg.Vector{0, 3, 0}}
	got := r.Transform(tl)

	if exp != *got {
		t.Errorf("expected %v, got %v", exp, *got)
	}
}
