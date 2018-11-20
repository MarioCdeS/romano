package vector

import (
	"fmt"
	"testing"

	"github.com/MarioCdeS/romano/tracer"
)

func ExampleNew() {
	v := New(4.3, -4.2, 3.1)
	fmt.Println(v)

	// Output: V(4.3, -4.2, 3.1)
}

func ExampleVector_Equal() {
	v1 := New(1, 2, 3)
	v2 := New(2, 2, 3)

	fmt.Println(v1.Equal(v1))
	fmt.Println(v1.Equal(v2))

	// Output:
	// true
	// false
}

func TestNew(t *testing.T) {
	v := New(4.3, -4.2, 3.1)
	assertVector(t, v)
}

func TestVector_AddVector(t *testing.T) {
	v1 := New(3, -2, 5)
	v2 := New(-2, 3, 1)
	expected := New(1, 1, 6)
	got := v1.Add(v2)

	assertVector(t, got)
	assertVectorsEqual(t, expected, got)
}

func TestVector_SubVector(t *testing.T) {
	v1 := New(3, -2, 5)
	v2 := New(-2, 3, 1)
	expected := New(5, -5, 4)
	got := v1.Sub(v2)

	assertVector(t, got)
	assertVectorsEqual(t, expected, got)
}

func TestVector_Mul(t *testing.T) {
	v := New(1, -2, 3)
	expected := New(3.5, -7, 10.5)
	got := v.Mul(3.5)

	assertVector(t, got)
	assertVectorsEqual(t, expected, got)

	v = New(1, -2, 3)
	expected = New(0.5, -1, 1.5)
	got = v.Mul(0.5)

	assertVectorsEqual(t, expected, got)
}

func TestVector_Div(t *testing.T) {
	v := New(1, -2, 3)
	expected := New(0.5, -1, 1.5)
	got := v.Div(2)

	assertVector(t, got)
	assertVectorsEqual(t, expected, got)
}

func TestVector_Neg(t *testing.T) {
	v := New(1, 2, 3)
	expected := New(-1, -2, -3)
	got := v.Neg()

	assertVector(t, got)
	assertVectorsEqual(t, expected, got)
}

func TestVector_Dot(t *testing.T) {
	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)
	expected := 20.0
	got := v1.Dot(v2)

	if !tracer.Equalf64(expected, got) {
		t.Errorf("expected %g, got %g", expected, got)
	}
}

func TestVector_Cross(t *testing.T) {
	v1 := New(1, 2, 3)
	v2 := New(2, 3, 4)
	expected := New(-1, 2, -1)
	got := v1.Cross(v2)

	assertVector(t, got)
	assertVectorsEqual(t, expected, got)

	expected = expected.Neg()
	got = v2.Cross(v1)

	assertVectorsEqual(t, expected, got)
}

func TestVector_Magnitude(t *testing.T) {
	v := New(0, 3, 4)
	expected := 5.0
	got := v.Magnitude()

	if !tracer.Equalf64(expected, got) {
		t.Errorf("expected %g, got %g", expected, got)
	}
}

func TestVector_Normalized(t *testing.T) {
	v := New(0, 3, 4)
	expected := New(0, 3.0/5.0, 4.0/5.0)
	got := v.Normalized()

	if !tracer.Equalf64(got.Magnitude(), 1) {
		t.Error("not a unit vector")
	}

	assertVector(t, got)
	assertVectorsEqual(t, expected, got)
}

func BenchmarkVector_Equals(b *testing.B) {
	v := New(1, 2, 3)

	for i := 0; i < b.N; i++ {
		v.Equal(v)
	}
}

func assertVector(t *testing.T, v *Vector) {
	if v[3] != 0 {
		t.Error("not a vector")
	}
}

func assertVectorsEqual(t *testing.T, expected *Vector, got *Vector) {
	if !expected.Equal(got) {
		t.Errorf("expected %s, got %s", expected, got)
	}
}
