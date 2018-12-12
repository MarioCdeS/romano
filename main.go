package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/MarioCdeS/romano/tracer"
)

func main() {
	cnvs := tracer.NewCanvas(800, 600)
	black := &tracer.Color{0, 0, 0, 1}

	xTick := 1 / float64(cnvs.Width())
	yTick := 1 / float64(cnvs.Height())

	var xAxis tracer.Vector
	var yAxis tracer.Vector

	if cnvs.Width() > cnvs.Height() {
		ratio := float64(cnvs.Width()) / float64(cnvs.Height())
		xAxis = tracer.Vector{ratio, 0, 0}
		yAxis = tracer.Vector{0, 1, 0}
	} else {
		ratio := float64(cnvs.Height()) / float64(cnvs.Width())
		xAxis = tracer.Vector{1, 0, 0}
		yAxis = tracer.Vector{0, ratio, 0}
	}

	center := tracer.Origin().MutAddVector(xAxis.Scale(0.5)).MutAddVector(yAxis.Scale(0.5))
	cam := center.AddVector(&tracer.Vector{0, 0, -1})
	light := tracer.PointLight{
		Position:  tracer.Point{-10, 10, -10},
		Intensity: tracer.Color{1, 1, 1, 1},
	}

	spherePos := center.AddVector(&tracer.Vector{0, 0, 2}).SubPoint(tracer.Origin())
	sphereMat := tracer.NewMaterial(&tracer.Color{0, 1, 0, 1}, 0.1, 0.9, 0.9, 200)
	sphere, err := tracer.NewTransformedSphere(
		tracer.NewTranslateFromVec(spherePos),
		sphereMat,
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for y := 0; y < cnvs.Height(); y++ {
		rayDirY := yAxis.Scale(float64(y) * yTick)

		for x := 0; x < cnvs.Width(); x++ {
			rayDir := tracer.Origin().
				MutAddVector(xAxis.Scale(float64(x) * xTick)).
				MutAddVector(rayDirY).
				SubPoint(cam)
			ray := tracer.Ray{Origin: *cam, Direction: *rayDir}
			var col *tracer.Color

			if hit, ok := tracer.FindHit(&ray, sphere.Intersections(&ray)); ok {
				col = tracer.At(hit.Point, hit.Normal, hit.Camera, &light, hit.Object().Material())
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
