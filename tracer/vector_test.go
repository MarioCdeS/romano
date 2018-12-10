package tracer

import (
	"math"
	"testing"
)

func TestVector_Reflect(t *testing.T) {
	v := Vector{1, -1, 0}
	n := Vector{0, 1, 0}
	exp := Vector{1, 1, 0}
	got := v.Reflect(&n)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(&exp, got))
	}
}

func TestVector_Reflect2(t *testing.T) {
	v := Vector{0, -1, 0}
	n := Vector{math.Sqrt2 / 2, math.Sqrt2 / 2, 0}
	exp := Vector{1, 0, 0}
	got := v.Reflect(&n)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(&exp, got))
	}
}
