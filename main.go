package main

import (
	"fmt"
	"image/png"
	"os"

	geom "github.com/MarioCdeS/romano/tracer/geometry"
	gfx "github.com/MarioCdeS/romano/tracer/graphics"
	lin "github.com/MarioCdeS/romano/tracer/linalg"
)

func main() {
	cnvs := gfx.NewCanvas(640, 480)
	red := gfx.Color{1, 0, 0, 1}
	black := gfx.Color{0, 0, 0, 1}

	xTick := 1 / float64(cnvs.Width())
	yTick := 1 / float64(cnvs.Height())

	var xAxis lin.Vector
	var yAxis lin.Vector

	if cnvs.Width() > cnvs.Height() {
		ratio := float64(cnvs.Width()) / float64(cnvs.Height())
		xAxis = lin.Vector{ratio, 0, 0}
		yAxis = lin.Vector{0, 1, 0}
	} else {
		ratio := float64(cnvs.Height()) / float64(cnvs.Width())
		xAxis = lin.Vector{1, 0, 0}
		yAxis = lin.Vector{0, ratio, 0}
	}

	center := geom.Origin().MutAddVector(xAxis.Scale(0.5)).MutAddVector(yAxis.Scale(0.5))
	cam := center.AddVector(&lin.Vector{0, 0, -1})

	spherePos := center.AddVector(&lin.Vector{0, 0, 2}).SubPoint(geom.Origin())
	sphere, err := geom.NewTransformedSphere(
		lin.NewTranslateFromVec(spherePos).Dot(lin.NewUniformScale(1.1)),
	)
	/*
		sphere, err := geometry.NewTransformedSphere(
			lin.NewTranslateFromVec(spherePos).Dot(lin.NewShear(1, 0, 0, 0, 0, 0)),
		)
	*/

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for y := 0; y < cnvs.Height(); y++ {
		rayDirY := yAxis.Scale(float64(y) * yTick)

		for x := 0; x < cnvs.Width(); x++ {
			rayDir := geom.Origin().
				MutAddVector(xAxis.Scale(float64(x) * xTick)).
				MutAddVector(rayDirY).
				SubPoint(cam)
			ray := geom.Ray{Origin: *cam, Direction: *rayDir}
			var col gfx.Color

			if _, ok := geom.Hit(sphere.Intersections(&ray)); ok {
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
