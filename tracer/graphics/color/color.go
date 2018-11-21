package color

import (
	"fmt"

	"github.com/MarioCdeS/romano/tracer/linalg/matrix"
)

type Color matrix.Mat4x1

func New(r, g, b, a float64) *Color {
	return &Color{r, g, b, a}
}

func (c *Color) Copy() *Color {
	res := *c
	return &res
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
	return (*Color)((*matrix.Mat4x1)(c).Add((*matrix.Mat4x1)(oth)))
}

func (c *Color) Sub(oth *Color) *Color {
	return (*Color)((*matrix.Mat4x1)(c).Sub((*matrix.Mat4x1)(oth)))
}

func (c *Color) MulScalar(scalar float64) *Color {
	return (*Color)((*matrix.Mat4x1)(c).Scale(scalar))
}

func (c *Color) MulColor(oth *Color) *Color {
	return (*Color)((*matrix.Mat4x1)(c).Hadamard((*matrix.Mat4x1)(oth)))
}

func (c *Color) Equal(oth *Color) bool {
	return (*matrix.Mat4x1)(c).Equal((*matrix.Mat4x1)(oth))
}

func (c *Color) String() string {
	return fmt.Sprintf("C(%g, %g, %g, %g)", c[0], c[1], c[2], c[3])
}
