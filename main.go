package main

import (
	"fmt"
	"github.com/MarioCdeS/romano/tracer/linalg"
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
		linalg.NewVector(1, 1, 0).Normalized(),
	}

	w := world{
		linalg.NewVector(0, -0.1, 0),
		linalg.NewVector(-0.01, 0, 0),
	}

	var count int

	for p.position.Y > 0 {
		pos := p.position.AddVector(p.velocity)
		vel := p.velocity.Add(w.gravity).Add(w.wind)
		p = projectile{pos, vel}
		count++

		fmt.Printf("%d -> Pos: %s\tVel: %s\n", count, pos, vel)
	}
}
