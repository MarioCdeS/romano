package main

import (
	"fmt"
	"image/png"
	"math"
	"os"

	"github.com/MarioCdeS/romano/tracer/graphics"
	"github.com/MarioCdeS/romano/tracer/linalg"
)

func main() {
	orig := linalg.Point{400, 300, 0}
	rotZ := linalg.NewRotateZ(math.Pi / 6)
	white := graphics.Color{1, 1, 1, 1}
	tick := linalg.Vector{200, 0, 0}

	cnvs := graphics.NewCanvas(800, 600)
	cnvs.Clear(graphics.Color{0, 0, 0, 1})

	for i := 0; i < 12; i++ {
		tickPoint := orig.AddVector(tick)
		cnvs.Set(int(tickPoint.X), int(tickPoint.Y), white)
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
