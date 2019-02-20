package raytracer

import (
	"testing"
)

func TestNewMatrix(t *testing.T) {
	m := NewMatrix(4, 4, []float64{
		1, 2, 3, 4,
		5.5, 6.5, 7.5, 8.5,
		9, 10, 11, 12,
		13.5, 14.5, 15.5, 16.5,
	})

	assertEqualFloat64(t, 1, m.At(0, 0))
	assertEqualFloat64(t, 4, m.At(0, 3))
	assertEqualFloat64(t, 5.5, m.At(1, 0))
	assertEqualFloat64(t, 7.5, m.At(1, 2))
	assertEqualFloat64(t, 11, m.At(2, 2))
	assertEqualFloat64(t, 13.5, m.At(3, 0))
	assertEqualFloat64(t, 15.5, m.At(3, 2))
}

func TestNewTwoByTwoMatrix(t *testing.T) {
	m := NewMatrix(2, 2, []float64{
		-3, 5,
		1, -2,
	})
	assertEqualFloat64(t, -3, m.At(0, 0))
	assertEqualFloat64(t, 5, m.At(0, 1))
	assertEqualFloat64(t, 1, m.At(1, 0))
	assertEqualFloat64(t, -2, m.At(1, 1))
}

func TestNewThreeByThreeMatrix(t *testing.T) {
	m := NewMatrix(3, 3, []float64{
		-3, 5, 0,
		1, -2, -7,
		0, 1, 1,
	})
	assertEqualFloat64(t, -3, m.At(0, 0))
	assertEqualFloat64(t, -2, m.At(1, 1))
	assertEqualFloat64(t, 1, m.At(2, 2))
}
