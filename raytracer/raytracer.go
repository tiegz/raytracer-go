package raytracer

import "math"

// TODO are these both the same semantically?
const equalFloat64sTolerance = 0.00001
const EPSILON = 0.00001

func equalFloat64s(x, y float64) bool {
	diff := math.Abs(x - y)
	return diff < equalFloat64sTolerance
}

func minFloat64(floats ...float64) float64 {
	min := floats[0]
	for _, f := range floats {
		if f < min {
			min = f
		}
	}
	return min
}

func maxFloat64(floats ...float64) float64 {
	max := floats[0]
	for _, f := range floats {
		if f > max {
			max = f
		}
	}
	return max
}
