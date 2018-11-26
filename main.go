package main

import (
	"fmt"
	"image/png"
	"math"
	"os"

	"github.com/MarioCdeS/romano/tracer/graphics/canvas"
	"github.com/MarioCdeS/romano/tracer/graphics/color"
	"github.com/MarioCdeS/romano/tracer/linalg/point"
	"github.com/MarioCdeS/romano/tracer/linalg/transform"
	"github.com/MarioCdeS/romano/tracer/linalg/vector"
)

func main() {
	orig := point.New(400, 300, 0)
	rotZ := transform.NewRotateZ(math.Pi / 6)
	white := color.New(1, 1, 1, 1)
	tick := vector.New(200, 0, 0)

	cnvs := canvas.New(800, 600)
	cnvs.Clear(color.New(0, 0, 0, 1))

	for i := 0; i < 12; i++ {
		tickDot := orig.AddVector(tick)
		cnvs.Set(int(tickDot.X), int(tickDot.Y), white)
		tick = rotZ.DotVector(tick)
	}

	f, err := os.Create("out.png")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer f.Close()

	if err := png.Encode(f, cnvs.ToImage()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
