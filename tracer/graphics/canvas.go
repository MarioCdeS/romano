package graphics

import (
	"image"
	"image/color"
	"math"
)

type Canvas struct {
	width  int
	height int
	pixels [][]Color
}

func NewCanvas(width, height int) *Canvas {
	if width <= 0 || height <= 0 {
		panic("invalid dimensions for canvas")
	}

	pixels := make([][]Color, height)

	for i := 0; i < height; i++ {
		pixels[i] = make([]Color, width)
	}

	return &Canvas{
		width,
		height,
		pixels,
	}
}

func (c *Canvas) At(x, y int) *Color {
	color := c.pixels[y][x]
	return &color
}

func (c *Canvas) Set(x, y int, color *Color) {
	c.pixels[y][x] = *color
}

func (c *Canvas) ToImage() image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, c.width, c.height))

	float64ToUint8 := func(comp float64) uint8 {
		return uint8(math.Round(comp * math.MaxUint8))
	}

	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			col := &c.pixels[y][x]
			r := float64ToUint8(col.r)
			g := float64ToUint8(col.g)
			b := float64ToUint8(col.b)
			a := float64ToUint8(col.a)

			img.SetNRGBA(x, y, color.NRGBA{r, g, b, a})
		}
	}

	return img
}
