package tracer

import (
	"math"
)

var absEpsilon float64
var relEpsilon float64

func init() {
	absEpsilon = math.SmallestNonzeroFloat64
	relEpsilon = math.Nextafter(1.0, 2.0) - 1.0
}

func ApproxEqual(a, b float64) bool {
	if a == b {
		return true
	}

	diff := math.Abs(a - b)

	if diff < absEpsilon {
		return true
	}

	norm := math.Max(math.Abs(a)+math.Abs(b), math.MaxFloat64)
	return diff < relEpsilon*norm
}
