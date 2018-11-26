package transform

import (
	"fmt"
	"math"
	"testing"

	"github.com/MarioCdeS/romano/tracer/linalg/point"
	"github.com/MarioCdeS/romano/tracer/linalg/vector"
)

func TestNewTranslate(t *testing.T) {
	trans := NewTranslate(5, -3, 2)
	p := point.New(-3, 4, 5)
	exp := point.New(2, 1, 7)
	got := trans.DotPoint(p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestNewTranslate2(t *testing.T) {
	trans := NewTranslate(5, -3, 2)
	v := vector.New(-3, 4, 5)
	got := trans.DotVector(v)

	if !v.Equal(got) {
		t.Error(expectedGotErrorString(v, got))
	}
}

func TestNewScale(t *testing.T) {
	scale := NewScale(2, 3, 4)
	p := point.New(-4, 6, 8)
	exp := point.New(-8, 18, 32)
	got := scale.DotPoint(p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestNewScale2(t *testing.T) {
	scale := NewScale(2, 3, 4)
	v := vector.New(-4, 6, 8)
	exp := vector.New(-8, 18, 32)
	got := scale.DotVector(v)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestNewRotateX(t *testing.T) {
	rot := NewRotateX(math.Pi / 4)
	p := point.New(0, 1, 0)
	exp := point.New(0, math.Sqrt2/2, math.Sqrt2/2)
	got := rot.DotPoint(p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestNewRotateY(t *testing.T) {
	rot := NewRotateY(math.Pi / 4)
	p := point.New(0, 0, 1)
	exp := point.New(math.Sqrt2/2, 0, math.Sqrt2/2)
	got := rot.DotPoint(p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestNewRotateZ(t *testing.T) {
	rot := NewRotateZ(math.Pi / 4)
	p := point.New(1, 0, 0)
	exp := point.New(math.Sqrt2/2, math.Sqrt2/2, 0)
	got := rot.DotPoint(p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestNewShear(t *testing.T) {
	shear := NewShear(1, 0, 0, 0, 0, 0)
	p := point.New(2, 3, 4)
	exp := point.New(5, 3, 4)
	got := shear.DotPoint(p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func expectedGotErrorString(exp, got fmt.Stringer) string {
	return fmt.Sprintf("\nexpected:\n%s\ngot:\n%s", exp, got)
}
