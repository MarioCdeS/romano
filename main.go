package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/MarioCdeS/romano/tracer/geometry"
	"github.com/MarioCdeS/romano/tracer/graphics"
	"github.com/MarioCdeS/romano/tracer/linalg"
)

func main() {
	cnvs := graphics.NewCanvas(800, 600)
	red := graphics.Color{1, 0, 0, 1}
	black := graphics.Color{0, 0, 0, 1}

	cam := linalg.Point{4, 3, -5}
	orig := linalg.Point{0, 0, 0}
	xAxis := linalg.Vector{8, 0, 0}
	yAxis := linalg.Vector{0, 6, 0}
	xTick := 1 / float64(cnvs.Width())
	yTick := 1 / float64(cnvs.Height())

	sphere, err := geometry.NewTransformedSphere(linalg.NewTranslate(4, 3, 0).Dot(linalg.NewScale(3, 3, 3)))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for y := 0; y < cnvs.Height(); y++ {
		for x := 0; x < cnvs.Width(); x++ {
			// We are not worried about the distance to the hit, just that a hit occurred, so rayDir doesn't need to be
			// normalised.
			origCopy := orig
			rayDir := origCopy.
				MutAddVector(xAxis.Scale(float64(x) * xTick)).
				MutAddVector(yAxis.Scale(float64(y) * yTick)).
				SubPoint(cam)
			ray := geometry.Ray{Origin: cam, Direction: rayDir}
			var col graphics.Color

			if _, ok := geometry.Hit(sphere.Intersections(&ray)); ok {
				col = red
			} else {
				col = black
			}

			cnvs.Set(x, cnvs.Height()-y-1, col)
		}
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
