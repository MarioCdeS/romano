package main

import (
	"fmt"
	"image/png"
	"math"
	"os"

	"github.com/MarioCdeS/romano/tracer/geometry"
	"github.com/MarioCdeS/romano/tracer/graphics"
	"github.com/MarioCdeS/romano/tracer/linalg"
)

func main() {
	cnvs := graphics.NewCanvas(800, 600)
	black := graphics.Color{0, 0, 0, 1}

	cam := linalg.Point{4, 3, -5}
	orig := linalg.Point{0, 0, 0}
	xAxis := linalg.Vector{8, 0, 0}
	yAxis := linalg.Vector{0, 6, 0}
	xTick := 1 / float64(cnvs.Width())
	yTick := 1 / float64(cnvs.Height())

	var squaredDistToCanv float64

	{
		center := orig.
			AddVector(xAxis.Scale(0.5)).
			MutAddVector(yAxis.Scale(0.5))
		squaredDistToCanv = center.SubPoint(&cam).SquaredMagnitude()
	}

	sphere, err := geometry.NewTransformedSphere(linalg.NewTranslate(4, 3, 3).Dot(linalg.NewScale(3, 3, 3)))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for y := 0; y < cnvs.Height(); y++ {
		for x := 0; x < cnvs.Width(); x++ {
			rayDir := orig.
				AddVector(xAxis.Scale(float64(x) * xTick)).
				MutAddVector(yAxis.Scale(float64(y) * yTick)).
				SubPoint(&cam)
			ray := geometry.Ray{Origin: cam, Direction: *rayDir}
			var col graphics.Color

			if hit, ok := geometry.Hit(sphere.Intersections(&ray)); ok {
				squaredDistToHit := ray.Direction.Scale(hit.T).SquaredMagnitude()
				col = graphics.Color{
					R: 0,
					G: math.Min(math.Max(2-(squaredDistToHit/squaredDistToCanv), 0), 1),
					B: 0,
					A: 1,
				}
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

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	if err := png.Encode(f, cnvs.ToImage()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
