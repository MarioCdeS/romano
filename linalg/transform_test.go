package linalg

import (
	"math"
	"testing"
)

func TestNewTranslate(t *testing.T) {
	trans := NewTranslate(5, -3, 2)
	p := Point{-3, 4, 5}
	exp := Point{2, 1, 7}
	got := trans.DotPoint(&p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(&exp, got))
	}
}

func TestNewTranslate2(t *testing.T) {
	trans := NewTranslate(5, -3, 2)
	v := Vector{-3, 4, 5}
	got := trans.DotVector(&v)

	if !v.Equal(got) {
		t.Error(expectedGotErrorString(&v, got))
	}
}

func TestNewScale(t *testing.T) {
	scale := NewScale(2, 3, 4)
	p := Point{-4, 6, 8}
	exp := Point{-8, 18, 32}
	got := scale.DotPoint(&p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(&exp, got))
	}
}

func TestNewScale2(t *testing.T) {
	scale := NewScale(2, 3, 4)
	v := Vector{-4, 6, 8}
	exp := Vector{-8, 18, 32}
	got := scale.DotVector(&v)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(&exp, got))
	}
}

func TestNewRotateX(t *testing.T) {
	rot := NewRotateX(math.Pi / 4)
	p := Point{0, 1, 0}
	exp := Point{0, math.Sqrt2 / 2, math.Sqrt2 / 2}
	got := rot.DotPoint(&p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(&exp, got))
	}
}

func TestNewRotateY(t *testing.T) {
	rot := NewRotateY(math.Pi / 4)
	p := Point{0, 0, 1}
	exp := Point{math.Sqrt2 / 2, 0, math.Sqrt2 / 2}
	got := rot.DotPoint(&p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(&exp, got))
	}
}

func TestNewRotateZ(t *testing.T) {
	rot := NewRotateZ(math.Pi / 4)
	p := Point{1, 0, 0}
	exp := Point{math.Sqrt2 / 2, math.Sqrt2 / 2, 0}
	got := rot.DotPoint(&p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(&exp, got))
	}
}

func TestNewShear(t *testing.T) {
	shear := NewShear(1, 0, 0, 0, 0, 0)
	p := Point{2, 3, 4}
	exp := Point{5, 3, 4}
	got := shear.DotPoint(&p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(&exp, got))
	}
}
