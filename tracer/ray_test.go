package tracer

import "testing"

func TestRay_Transform(t *testing.T) {
	r := &Ray{Point{1, 2, 3}, Vector{0, 1, 0}}
	tl := NewTranslate(3, 4, 5)
	exp := Ray{Point{4, 6, 8}, Vector{0, 1, 0}}
	got := r.Transform(tl)

	if exp != *got {
		t.Errorf("expected %v, got %v", exp, *got)
	}
}

func TestRay_Transform2(t *testing.T) {
	r := &Ray{Point{1, 2, 3}, Vector{0, 1, 0}}
	tl := NewScale(2, 3, 4)
	exp := Ray{Point{2, 6, 12}, Vector{0, 3, 0}}
	got := r.Transform(tl)

	if exp != *got {
		t.Errorf("expected %v, got %v", exp, *got)
	}
}
