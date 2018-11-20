package point

import (
	"fmt"
	"testing"

	"github.com/MarioCdeS/romano/tracer/linalg/vector"
)

func ExampleNew() {
	p := New(4.3, -4.2, 3.1)
	fmt.Println(p)

	// Output: P(4.3, -4.2, 3.1)
}

func ExamplePoint_Equal() {
	p1 := New(1, 2, 3)
	p2 := New(2, 3, 4)

	fmt.Println(p1.Equal(p1))
	fmt.Println(p1.Equal(p2))

	// Output:
	// true
	// false
}

func TestNew(t *testing.T) {
	p := New(4.3, -4.2, 3.1)
	assertPoint(t, p)
}

func TestPoint_AddVector(t *testing.T) {
	p := New(1, 2, 3)
	v := vector.New(1, 2, 3)
	expected := New(2, 4, 6)
	got := p.Add(v)

	assertPoint(t, got)
	assertPointsEqual(t, expected, got)
}

func TestPoint_SubVector(t *testing.T) {
	p := New(1, 2, 3)
	v := vector.New(3, 2, 1)
	expected := New(-2, 0, 2)
	got := p.SubVector(v)

	assertPoint(t, got)
	assertPointsEqual(t, expected, got)
}

func TestPoint_SubPoint(t *testing.T) {
	p1 := New(1, 2, 3)
	p2 := New(3, 2, 1)
	expected := vector.New(-2, 0, 2)
	got := p1.SubPoint(p2)

	if got[3] != 0 {
		t.Error("not a vector")
	}

	if !expected.Equal(got) {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func assertPoint(t *testing.T, p *Point) {
	if p[3] != 1 {
		t.Error("not a point")
	}
}

func assertPointsEqual(t *testing.T, expected *Point, got *Point) {
	if !expected.Equal(got) {
		t.Errorf("expected %s, got %s", expected, got)
	}
}
