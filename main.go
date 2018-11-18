package main

import (
	"fmt"
	"github.com/MarioCdeS/romano/tracer/graphics"
	"github.com/MarioCdeS/romano/tracer/linalg"
	"image/png"
	"math"
	"os"
)

type projectile struct {
	position *linalg.Point
	velocity *linalg.Vector
}

type world struct {
	gravity *linalg.Vector
	wind    *linalg.Vector
}

func main() {
	p := projectile{
		linalg.NewPoint(0, 1, 0),
		linalg.NewVector(0.5, 1, 0).Normalized().Mul(10),
	}

	w := world{
		linalg.NewVector(0, -0.1, 0),
		linalg.NewVector(-0.01, 0, 0),
	}

	canvas := graphics.NewCanvas(800, 600)
	canvas.Clear(graphics.NewColor(0, 0, 0, 1))

	white := graphics.NewColor(1, 1, 1, 1)

	for p.position.X() < float64(canvas.Width()) && p.position.Y() > 0 {
		p.position = p.position.Add(p.velocity)
		p.velocity = p.velocity.Add(w.gravity).Add(w.wind)

		x := int(math.Floor(p.position.X()))
		y := canvas.Height() - int(math.Floor(p.position.Y())) - 1
		canvas.Set(x, y, white)
	}

	f, err := os.Create("out.png")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer f.Close()

	if err := png.Encode(f, canvas.ToImage()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
