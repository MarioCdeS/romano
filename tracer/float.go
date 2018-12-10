package tracer

import "math"

var epsilon = (math.Nextafter(1.0, 2.0) - 1.0) * 10e6

func ApproxEqual(a, b float64) bool {
	if a == b {
		return true
	}

	diff := math.Abs(a - b)
	return diff < epsilon || diff < epsilon*(math.Max(math.Abs(a), math.Abs(b)))
}

func Clamp(x, min, max float64) float64 {
	return math.Min(math.Max(x, min), max)
}
