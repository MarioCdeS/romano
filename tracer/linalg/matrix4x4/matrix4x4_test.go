package matrix4x4

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/MarioCdeS/romano/tracer"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().Unix())
	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	elems := randomElements()
	m := New(elems...)

	for i, e := range elems {
		row := i / 4
		col := i % 4
		got := m[row][col]

		if !tracer.Equalf64(e, got) {
			t.Errorf("at (%d, %d), expected %g, got %g", row, col, e, got)
		}
	}
}

func TestAdd(t *testing.T) {
	elems1 := randomElements()
	elems2 := randomElements()
	sums := make([]float64, 16)

	for i := 0; i < 16; i++ {
		sums[i] = elems1[i] + elems2[i]
	}

	m1 := New(elems1...)
	m2 := New(elems2...)
	rem := *m1
	exp := New(sums...)
	got := Add(m1, m2)

	if !rem.Equal(m1) {
		t.Error("first argument modified")
	}

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestScale(t *testing.T) {
	elems := randomElements()
	scaled := make([]float64, 16)

	for i := 0; i < 16; i++ {
		scaled[i] = elems[i] * 5
	}

	m := New(elems...)
	rem := *m
	exp := New(scaled...)
	got := Scale(m, 5)

	if !rem.Equal(m) {
		t.Error("argument modified")
	}

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestTranspose(t *testing.T) {
	elems := randomElements()
	transElems := make([]float64, 16)

	for i := 0; i < 16; i++ {
		transElems[(i%4)*4+(i/4)] = elems[i]
	}

	m := New(elems...)
	rem := *m
	exp := New(transElems...)
	got := Transpose(m)

	if !rem.Equal(m) {
		t.Error("argument modified")
	}

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestDot(t *testing.T) {
	m1 := New(randomElements()...)
	rem := *m1
	m2 := New(randomElements()...)
	exp := New()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				exp[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}

	got := Dot(m1, m2)

	if !rem.Equal(m1) {
		t.Error("argument modified")
	}

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestMatrix4x4_Equal(t *testing.T) {
	m := New(randomElements()...)

	if !m.Equal(m) {
		t.Error("not reflexive")
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			c := *m
			c[i][j] = m[i][j] + 1

			if c.Equal(m) {
				t.Errorf("at (%d, %d), %g not equal %g", i, j, m[i][j], c[i][j])
			}
		}
	}
}

func randomElements() []float64 {
	res := make([]float64, 16)

	for i := 0; i < 16; i++ {
		res[i] = rand.Float64() * 100
	}

	return res
}

func expectedGotErrorString(exp, got *Matrix4x4) string {
	return fmt.Sprintf("\nexpected:\n%sgot:\n%s", exp, got)
}
