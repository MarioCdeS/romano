package geometry

import "github.com/MarioCdeS/romano/tracer/linalg"

var origin = linalg.Point{0, 0, 0}

func Origin() *linalg.Point {
	res := origin
	return &res
}
