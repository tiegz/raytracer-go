package raytracer

import "math"

const EPSILON = 0.00001 // TODO: rename? this is the difference btwn floats less thanw which we'd conisder them the same

func equalFloat64s(x, y float64) bool {
	if math.IsInf(x, -1) && math.IsInf(y, -1) {
		return true
	} else if math.IsInf(x, 1) && math.IsInf(y, 1) {
		return true
	} else {
		// TODO: wait wait wait this could be wrong
		diff := math.Abs(x - y)
		return diff < EPSILON
	}
}

func equalFloat64Slices(x, y []float64) bool {
	if len(x) != len(y) {
		return false
	}
	for idx, item := range x {
		if item != y[idx] {
			return false
		}
	}
	return true
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
