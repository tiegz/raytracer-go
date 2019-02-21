package raytracer

import "math"

const equalFloat64sTolerance = 0.00001

func equalFloat64s(x, y float64) bool {
	diff := math.Abs(x - y)
	return diff < equalFloat64sTolerance
}
