package tracer

import (
	"testing"
)

var (
	black *Color
	red   *Color
)

func init() {
	black = &Color{0, 0, 0, 0}
	red = &Color{1, 0, 0, 1}
}

func TestNew(t *testing.T) {
	c := NewCanvas(10, 10)

Outer:
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			if !c.pixels[y][x].Equal(black) {
				t.Error("canvas not initialized to black")
				break Outer
			}
		}
	}
}

func TestCanvas_At(t *testing.T) {
	c := NewCanvas(5, 5)
	at := c.At(0, 0)

	if !at.Equal(black) {
		t.Errorf("expected %s, got %s", black, at)
	}
}

func TestCanvas_Set(t *testing.T) {
	c := NewCanvas(5, 5)
	c.Set(0, 0, red)

	at := c.At(0, 0)

	if !at.Equal(red) {
		t.Errorf("expected %s, got %s", red, at)
	}
}
