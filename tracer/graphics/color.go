package graphics

import (
	"fmt"
	"github.com/MarioCdeS/romano/tracer/linalg"
)

type Color linalg.Quadruple

func NewColor(r, g, b, a float64) *Color {
	return &Color{r, g, b, a}
}

func (c *Color) R() float64 {
	return c[0]
}

func (c *Color) G() float64 {
	return c[1]
}

func (c *Color) B() float64 {
	return c[2]
}

func (c *Color) A() float64 {
	return c[3]
}

func (c *Color) Add(oth *Color) *Color {
	return (*Color)((*linalg.Quadruple)(c).Add((*linalg.Quadruple)(oth)))
}

func (c *Color) Sub(oth *Color) *Color {
	return (*Color)((*linalg.Quadruple)(c).Sub((*linalg.Quadruple)(oth)))
}

func (c *Color) MulScalar(scalar float64) *Color {
	return (*Color)((*linalg.Quadruple)(c).Mul(scalar))
}

func (c *Color) MulColor(oth *Color) *Color {
	return (*Color)((*linalg.Quadruple)(c).Hadamard((*linalg.Quadruple)(oth)))
}

func (c *Color) Equal(oth *Color) bool {
	return (*linalg.Quadruple)(c).Equal((*linalg.Quadruple)(oth))
}

func (c *Color) String() string {
	return fmt.Sprintf("C(%g, %g, %g, %g)", c[0], c[1], c[2], c[3])
}
