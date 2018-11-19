package matrix4x4

import (
	"testing"

	"github.com/MarioCdeS/romano/tracer"
)

func TestNew(t *testing.T) {
	elems := make([]float64, 16)

	for i := 0; i < 16; i++ {
		elems[i] = float64(i)
	}

	m := New(elems...)

	for i, e := range elems {
		row := i / 4
		col := i % 4
		got := m.At(row, col)

		if !tracer.Equalf64(e, got) {
			t.Errorf("at (%d, %d), expected %g, got %g", row, col, e, got)
		}
	}
}

func TestAdd(t *testing.T) {
	elems1 := make([]float64, 16)
	elems2 := make([]float64, 16)

	for i := 0; i < 16; i++ {
		elems1[i] = float64(i)
		elems2[i] = float64(i + 1)
	}

	m1 := New(elems1...)
	m2 := New(elems2...)
	sum := Add(m1, m2)

	for i, e := range elems1 {
		row := i / 4
		col := i % 4
		got := m1.At(row, col)

		if !tracer.Equalf64(e, got) {
			t.Errorf("at (%d, %d), expected %g, got %g", row, col, e, got)
		}

		exp := e + m2.At(row, col)
		got = sum.At(row, col)

		if !tracer.Equalf64(exp, got) {
			t.Errorf("at (%d, %d), expected %g, got %g", row, col, exp, got)
		}
	}
}
