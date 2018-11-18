package linalg

import (
	"fmt"
	"testing"
)

func ExampleNewPoint() {
	p := NewPoint(4.3, -4.2, 3.1)
	fmt.Println(p)

	// Output: P(4.3, -4.2, 3.1)
}

func ExamplePoint_Equal() {
	p1 := NewPoint(1, 2, 3)
	p2 := NewPoint(2, 3, 4)

	fmt.Println(p1.Equal(p1))
	fmt.Println(p1.Equal(p2))

	// Output:
	// true
	// false
}

func TestNewPoint(t *testing.T) {
	p := NewPoint(4.3, -4.2, 3.1)
	assertPoint(t, p)
}

func TestPoint_AddVector(t *testing.T) {
	p := NewPoint(1, 2, 3)
	v := NewVector(1, 2, 3)
	expected := NewPoint(2, 4, 6)
	got := p.Add(v)

	assertPoint(t, got)
	assertPointsEqual(t, expected, got)
}

func TestPoint_SubVector(t *testing.T) {
	p := NewPoint(1, 2, 3)
	v := NewVector(3, 2, 1)
	expected := NewPoint(-2, 0, 2)
	got := p.SubVector(v)

	assertPoint(t, got)
	assertPointsEqual(t, expected, got)
}

func TestPoint_SubPoint(t *testing.T) {
	p1 := NewPoint(1, 2, 3)
	p2 := NewPoint(3, 2, 1)
	expected := NewVector(-2, 0, 2)
	got := p1.SubPoint(p2)

	assertVector(t, got)
	assertVectorsEqual(t, expected, got)
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
