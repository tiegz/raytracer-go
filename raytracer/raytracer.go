package raytracer

import "math"

// TODO are these both the same semantically?
const equalFloat64sTolerance = 0.00001
const EPSILON = 0.00001

func equalFloat64s(x, y float64) bool {
	diff := math.Abs(x - y)
	return diff < equalFloat64sTolerance
}
