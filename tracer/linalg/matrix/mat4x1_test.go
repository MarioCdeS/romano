package matrix

import (
	"math/rand"
	"testing"

	"github.com/MarioCdeS/romano/tracer"
)

func TestNew4x1(t *testing.T) {
	elems := randomElements4()
	m := New4x1(elems...)

	for i, e := range elems {
		got := m[i]

		if !tracer.Equalf64(e, got) {
			t.Errorf("at (%d, 0), expected %g, got %g", i, e, got)
		}
	}
}

func TestMat4x1_Add(t *testing.T) {
	elems1 := randomElements4()
	elems2 := randomElements4()
	sums := make([]float64, 4)

	for i := 0; i < 4; i++ {
		sums[i] = elems1[i] + elems2[i]
	}

	exp := New4x1(sums...)
	got := New4x1(elems1...).Add(New4x1(elems2...))

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestMatrix4x1_Equal(t *testing.T) {
	m := New4x1(randomElements4()...)

	if !m.Equal(m) {
		t.Error("not reflexive")
	}

	for i := 0; i < 4; i++ {
		c := *m
		c[i] = m[i] + 1

		if c.Equal(m) {
			t.Errorf("at (%d, 0), %g not equal %g", i, m[i], c[i])
		}
	}
}

func randomElements4() []float64 {
	res := make([]float64, 4)

	for i := 0; i < 4; i++ {
		res[i] = rand.Float64() * 100
	}

	return res
}
