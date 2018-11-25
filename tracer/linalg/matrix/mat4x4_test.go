package matrix

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/MarioCdeS/romano/tracer/float"
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

		if !float.ApproxEqual(e, got) {
			t.Errorf("at (%d, %d), expected %g, got %g", row, col, e, got)
		}
	}
}

func TestNew4x4Translate(t *testing.T) {
	trans := New4x4Translate(5, -3, 2)
	p := New4x1(-3, 4, 5, 1)
	exp := New4x1(2, 1, 7, 1)
	got := trans.Dot4x1(p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestNew4x4Translate2(t *testing.T) {
	trans := New4x4Translate(5, -3, 2)
	v := New4x1(-3, 4, 5, 0)
	got := trans.Dot4x1(v)

	if !v.Equal(got) {
		t.Error(expectedGotErrorString(v, got))
	}
}

func TestNew4x4Scale(t *testing.T) {
	scale := New4x4Scale(2, 3, 4)
	p := New4x1(-4, 6, 8, 1)
	exp := New4x1(-8, 18, 32, 1)
	got := scale.Dot4x1(p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestNew4x4Scale2(t *testing.T) {
	scale := New4x4Scale(2, 3, 4)
	v := New4x1(-4, 6, 8, 0)
	exp := New4x1(-8, 18, 32, 0)
	got := scale.Dot4x1(v)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestNew4x4RotateX(t *testing.T) {
	rot := New4x4RotateX(math.Pi / 4)
	p := New4x1(0, 1, 0, 1)
	exp := New4x1(0, math.Sqrt2/2, math.Sqrt2/2, 1)
	got := rot.Dot4x1(p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestNew4x4RotateY(t *testing.T) {
	rot := New4x4RotateY(math.Pi / 4)
	p := New4x1(0, 0, 1, 1)
	exp := New4x1(math.Sqrt2/2, 0, math.Sqrt2/2, 1)
	got := rot.Dot4x1(p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestNew4x4RotateZ(t *testing.T) {
	rot := New4x4RotateZ(math.Pi / 4)
	p := New4x1(1, 0, 0, 1)
	exp := New4x1(math.Sqrt2/2, math.Sqrt2/2, 0, 1)
	got := rot.Dot4x1(p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestNew4x4Shear(t *testing.T) {
	shear := New4x4Shear(1, 0, 0, 0, 0, 0)
	p := New4x1(2, 3, 4)
	exp := New4x1(5, 3, 4)
	got := shear.Dot4x1(p)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
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

func TestMat4x4_Dot4x1(t *testing.T) {
	m1 := New4x4(randomElements16()...)
	m2 := New4x1(randomElements4()...)
	exp := New4x1()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			exp[i] += m1[i][j] * m2[j]
		}
	}

	got := m1.Dot4x1(m2)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestMat4x4_Dot_Mat4x4ID(t *testing.T) {
	m := New4x4(randomElements16()...)
	got := m.Dot(New4x4ID())

	if !m.Equal(got) {
		t.Error(expectedGotErrorString(m, got))
	}
}

func TestMat4x4_Det(t *testing.T) {
	m := New4x4(-2, -8, 3, 5, -3, 1, 7, 3, 1, 2, -9, 6, -6, 7, 7, -9)
	got := m.Det()

	if !float.ApproxEqual(-4071, got) {
		t.Errorf("expected -4071, got %g", got)
	}
}

func TestMat4x4_Inv(t *testing.T) {
	const det = 532

	m := New4x4(-5, 2, 6, -8, 1, -5, 1, 8, 7, 7, -6, -7, 1, -3, 7, 4)

	exp := New4x4(
		116, 240, 128, -24,
		-430, -775, -236, 277,
		-42, -119, -28, 105,
		-278, -433, -160, 163,
	).Scale(1.0 / det)

	got, err := m.Inv()

	if err != nil {
		t.Error(err)
	} else if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestMat4x4_Inv2(t *testing.T) {
Outer:
	for i := 0; i < 1000; i++ {
		for {
			a := New4x4(randomElements16()...)
			b := New4x4(randomElements16()...)
			c := a.Dot(b)

			if inv, err := b.Inv(); err == nil {
				got := c.Dot(inv)

				if !a.Equal(got) {
					t.Error(expectedGotErrorString(a, got))
					break Outer
				}

				break
			}
		}
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
	return fmt.Sprintf("\nexpected:\n%s\ngot:\n%s", exp, got)
}
