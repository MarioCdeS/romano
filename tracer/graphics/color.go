package graphics

import (
	"fmt"

	"github.com/MarioCdeS/romano/tracer"
)

type Color struct {
	r float64
	g float64
	b float64
	a float64
}

func NewColor(r, g, b, a float64) *Color {
	return &Color{r, g, b, a}
}

func (c *Color) R() float64 {
	return c.r
}

func (c *Color) G() float64 {
	return c.g
}

func (c *Color) B() float64 {
	return c.b
}

func (c *Color) A() float64 {
	return c.a
}

func (c *Color) Add(other *Color) *Color {
	return &Color{
		c.r + other.r,
		c.g + other.g,
		c.b + other.b,
		c.a + other.a,
	}
}

func (c *Color) Sub(other *Color) *Color {
	return &Color{
		c.r - other.r,
		c.g - other.g,
		c.b - other.b,
		c.a - other.a,
	}
}

func (c *Color) MulScalar(scalar float64) *Color {
	return &Color{
		c.r * scalar,
		c.g * scalar,
		c.b * scalar,
		c.a * scalar,
	}
}

func (c *Color) MulColor(other *Color) *Color {
	// Hadamard product
	return &Color{
		c.r * other.r,
		c.g * other.g,
		c.b * other.b,
		c.a * other.a,
	}
}

func (c *Color) Equal(other *Color) bool {
	return tracer.Equalf64(c.r, other.r) &&
		tracer.Equalf64(c.g, other.g) &&
		tracer.Equalf64(c.b, other.b) &&
		tracer.Equalf64(c.a, other.a)
}

func (c *Color) String() string {
	return fmt.Sprintf("C(%g, %g, %g, %g)", c.r, c.g, c.b, c.a)
}
