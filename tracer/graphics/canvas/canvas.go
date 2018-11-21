package canvas

import (
	"image"
	"image/color"
	"math"

	col "github.com/MarioCdeS/romano/tracer/graphics/color"
)

type Canvas struct {
	width  int
	height int
	pixels [][]col.Color
}

func New(width, height int) *Canvas {
	if width <= 0 || height <= 0 {
		panic("invalid dimensions for canvas")
	}

	pixels := make([][]col.Color, height)

	for i := 0; i < height; i++ {
		pixels[i] = make([]col.Color, width)
	}

	return &Canvas{
		width,
		height,
		pixels,
	}
}

func (c *Canvas) Width() int {
	return c.width
}

func (c *Canvas) Height() int {
	return c.height
}

func (c *Canvas) Clear(col *col.Color) {
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			c.pixels[y][x] = *col
		}
	}
}

func (c *Canvas) At(x, y int) *col.Color {
	col := c.pixels[y][x]
	return &col
}

func (c *Canvas) Set(x, y int, col *col.Color) {
	if x < 0 || x >= c.width {
		return
	}

	if y < 0 || y >= c.height {
		return
	}

	c.pixels[y][x] = *col
}

func (c *Canvas) ToImage() image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, c.width, c.height))

	float64ToUint8 := func(comp float64) uint8 {
		return uint8(math.Round(comp * math.MaxUint8))
	}

	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			col := &c.pixels[y][x]
			r := float64ToUint8(col.R())
			g := float64ToUint8(col.G())
			b := float64ToUint8(col.B())
			a := float64ToUint8(col.A())

			img.SetNRGBA(x, y, color.NRGBA{r, g, b, a})
		}
	}

	return img
}
