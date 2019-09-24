package raytracer

import "math"

const EPSILON = 0.00001 // TODO: rename? this is the difference btwn floats less thanw which we'd conisder them the same

// TODO: wait wait wait this could be wrong
func equalFloat64s(x, y float64) bool {
	diff := math.Abs(x - y)
	return diff < EPSILON
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
