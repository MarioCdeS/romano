package linalg

import (
	"math"
)

var epsilon float64

func init() {
	epsilon = math.Nextafter(1.0, 2.0) - 1.0
}

func Equalf64(x, y float64) bool {
	return math.Abs(x-y) < epsilon
}
