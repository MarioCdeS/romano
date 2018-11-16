package graphics

import (
	"testing"
)

func TestColor_Add(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75, 0.5)
	c2 := NewColor(0.7, 0.1, 0.25, 0.5)
	expected := NewColor(1.6, 0.7, 1.0, 1.0)
	got := c1.Add(c2)

	assertColorEqual(t, expected, got)
}

func assertColorEqual(t *testing.T, expected *Color, got *Color) {
	if !expected.Equal(got) {
		t.Errorf("expected %s, got %s", expected, got)
	}
}
