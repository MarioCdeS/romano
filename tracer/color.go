package tracer

import "fmt"

type Color struct {
	R, G, B, A float64
}

func (c Color) Add(oth *Color) *Color {
	return c.MutAdd(oth)
}

func (c *Color) MutAdd(oth *Color) *Color {
	c.R += oth.R
	c.G += oth.G
	c.B += oth.B
	c.A += oth.A

	return c
}

func (c Color) Sub(oth *Color) *Color {
	return c.MutSub(oth)
}

func (c *Color) MutSub(oth *Color) *Color {
	c.R -= oth.R
	c.G -= oth.G
	c.B -= oth.B
	c.A -= oth.A

	return c
}

func (c Color) Scale(scalar float64) *Color {
	return c.MutScale(scalar)
}

func (c *Color) MutScale(scalar float64) *Color {
	c.R *= scalar
	c.G *= scalar
	c.B *= scalar
	c.A *= scalar

	return c
}

func (c Color) Hadamard(oth *Color) *Color {
	return c.MutHadamard(oth)
}

func (c *Color) MutHadamard(oth *Color) *Color {
	c.R *= oth.R
	c.G *= oth.G
	c.B *= oth.B
	c.A *= oth.A

	return c
}

func (c *Color) Equal(oth *Color) bool {
	return ApproxEqual(c.R, oth.R) &&
		ApproxEqual(c.G, oth.G) &&
		ApproxEqual(c.B, oth.B) &&
		ApproxEqual(c.A, oth.A)
}

func (c *Color) String() string {
	return fmt.Sprintf("C(%g, %g, %g, %g)", c.R, c.G, c.B, c.A)
}
