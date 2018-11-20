package matrix

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

func TestNew4x4(t *testing.T) {
	elems := randomElements16()
	m := New4x4(elems...)

	for i, e := range elems {
		row := i / 4
		col := i % 4
		got := m[row][col]

		if !tracer.Equalf64(e, got) {
			t.Errorf("at (%d, %d), expected %g, got %g", row, col, e, got)
		}
	}
}

func TestMat4x4_Add(t *testing.T) {
	elems1 := randomElements16()
	elems2 := randomElements16()
	sums := make([]float64, 16)

	for i := 0; i < 16; i++ {
		sums[i] = elems1[i] + elems2[i]
	}

	exp := New4x4(sums...)
	got := New4x4(elems1...).Add(New4x4(elems2...))

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestMat4x4_Scale(t *testing.T) {
	elems := randomElements16()
	scaled := make([]float64, 16)

	for i := 0; i < 16; i++ {
		scaled[i] = elems[i] * 5
	}

	exp := New4x4(scaled...)
	got := New4x4(elems...).Scale(5)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestMat4x4_T(t *testing.T) {
	elems := randomElements16()
	transElems := make([]float64, 16)

	for i := 0; i < 16; i++ {
		transElems[(i%4)*4+(i/4)] = elems[i]
	}

	exp := New4x4(transElems...)
	got := New4x4(elems...).T()

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestMat4x4_Dot(t *testing.T) {
	m1 := New4x4(randomElements16()...)
	m2 := New4x4(randomElements16()...)
	exp := New4x4()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				exp[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}

	got := m1.Dot(m2)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestMatrix4x4_Equal(t *testing.T) {
	m := New4x4(randomElements16()...)

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

func randomElements16() []float64 {
	res := make([]float64, 16)

	for i := 0; i < 16; i++ {
		res[i] = rand.Float64() * 100
	}

	return res
}

func expectedGotErrorString(exp, got fmt.Stringer) string {
	return fmt.Sprintf("\nexpected:\n%sgot:\n%s", exp, got)
}
