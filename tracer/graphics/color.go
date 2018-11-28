package graphics

import (
	"fmt"

	"github.com/MarioCdeS/romano/tracer/float"
)

type Color struct {
	R, G, B, A float64
}

func (c Color) Add(oth Color) Color {
	c.MutAdd(oth)
	return c
}

func (c *Color) MutAdd(oth Color) *Color {
	c.R += oth.R
	c.G += oth.G
	c.B += oth.B
	c.A += oth.A

	return c
}

func (c Color) Sub(oth Color) Color {
	c.MutSub(oth)
	return c
}

func (c *Color) MutSub(oth Color) *Color {
	c.R -= oth.R
	c.G -= oth.G
	c.B -= oth.B
	c.A -= oth.A

	return c
}

func (c Color) Scale(scalar float64) Color {
	c.MutScale(scalar)
	return c
}

func (c *Color) MutScale(scalar float64) *Color {
	c.R *= scalar
	c.G *= scalar
	c.B *= scalar
	c.A *= scalar

	return c
}

func (c Color) Hadamard(oth Color) Color {
	c.MutHadamard(oth)
	return c
}

func (c *Color) MutHadamard(oth Color) *Color {
	c.R *= oth.R
	c.G *= oth.G
	c.B *= oth.B
	c.A *= oth.A

	return c
}

func (c Color) Equal(oth Color) bool {
	return float.ApproxEqual(c.R, oth.R) &&
		float.ApproxEqual(c.G, oth.G) &&
		float.ApproxEqual(c.B, oth.B) &&
		float.ApproxEqual(c.A, oth.A)
}

func (c Color) String() string {
	return fmt.Sprintf("C(%g, %g, %g, %g)", c.R, c.G, c.B, c.A)
}
