package float

import (
	"math"
)

var absEpsilon float64
var relEpsilon float64

func init() {
	absEpsilon = 10e6 * math.SmallestNonzeroFloat64
	relEpsilon = 10e6 * (math.Nextafter(1.0, 2.0) - 1.0)
}

func ApproxEqual(a, b float64) bool {
	if a == b {
		return true
	}

	diff := math.Abs(a - b)

	if diff < absEpsilon {
		return true
	}

	norm := math.Abs(a) + math.Abs(b)
	return diff < relEpsilon*norm
}
