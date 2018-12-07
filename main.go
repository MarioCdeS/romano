package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/MarioCdeS/romano/geometry"
	"github.com/MarioCdeS/romano/gfx"
	"github.com/MarioCdeS/romano/lighting"
	"github.com/MarioCdeS/romano/linalg"
)

func main() {
	cnvs := gfx.NewCanvas(800, 600)
	black := &gfx.Color{0, 0, 0, 1}

	xTick := 1 / float64(cnvs.Width())
	yTick := 1 / float64(cnvs.Height())

	var xAxis linalg.Vector
	var yAxis linalg.Vector

	if cnvs.Width() > cnvs.Height() {
		ratio := float64(cnvs.Width()) / float64(cnvs.Height())
		xAxis = linalg.Vector{ratio, 0, 0}
		yAxis = linalg.Vector{0, 1, 0}
	} else {
		ratio := float64(cnvs.Height()) / float64(cnvs.Width())
		xAxis = linalg.Vector{1, 0, 0}
		yAxis = linalg.Vector{0, ratio, 0}
	}

	center := geometry.Origin().MutAddVector(xAxis.Scale(0.5)).MutAddVector(yAxis.Scale(0.5))
	cam := center.AddVector(&linalg.Vector{0, 0, -1})
	light := lighting.PointLight{
		Position:  linalg.Point{-10, 10, -10},
		Intensity: gfx.Color{1, 1, 1, 1},
	}

	spherePos := center.AddVector(&linalg.Vector{0, 0, 2}).SubPoint(geometry.Origin())
	sphereMat := geometry.NewMaterial(&gfx.Color{1, 0.2, 1, 1}, 0.1, 0.9, 0.9, 200)
	sphere, err := geometry.NewTransformedSphere(
		linalg.NewTranslateFromVec(spherePos),
		sphereMat,
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for y := 0; y < cnvs.Height(); y++ {
		rayDirY := yAxis.Scale(float64(y) * yTick)

		for x := 0; x < cnvs.Width(); x++ {
			rayDir := geometry.Origin().
				MutAddVector(xAxis.Scale(float64(x) * xTick)).
				MutAddVector(rayDirY).
				SubPoint(cam)
			ray := geometry.Ray{Origin: *cam, Direction: *rayDir}
			var col *gfx.Color

			if hit, ok := geometry.Hit(sphere.Intersections(&ray)); ok {
				point := ray.PointAt(hit.T)
				camVec := ray.Direction.Neg().MutNormalized()
				col = lighting.At(point, hit.Object().NormalAt(point), camVec, &light, hit.Object().Material())
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
