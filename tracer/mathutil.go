package tracer

import (
	"math"
)

var epsilon float64

func init() {
	epsilon = math.Nextafter(1.0, 2.0) - 1.0
}

func ApproxEqual(a, b float64) bool {
	if a == b {
		return true
	}

	diff := math.Abs(a - b)
	norm := math.Min(math.Abs(a)+math.Abs(b), math.MaxFloat64)

	// TEST
	// return diff < math.Max(epsilon*norm, math.SmallestNonzeroFloat64)
	t := epsilon * norm
	s := math.SmallestNonzeroFloat64
	return diff < math.Max(t, s)
}
