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

func TestMatrixSet(t *testing.T) {
	actual := NewMatrix(2, 2, make([]float64, 4, 4))
	actual.Set(0, 0, 1)
	actual.Set(0, 1, 2)
	actual.Set(1, 0, 3)
	actual.Set(1, 1, 4)
	expected := NewMatrix(2, 2, []float64{
		1, 2,
		3, 4,
	})

	assertEqualMatrix(t, actual, expected)
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

func TestMatrixEquality(t *testing.T) {
	m1 := NewMatrix(4, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	})
	m2 := NewMatrix(4, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	})

	assertEqualMatrix(t, m1, m2)
}
func TestMatrixInequality(t *testing.T) {
	m1 := NewMatrix(4, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	})
	m2 := NewMatrix(4, 4, []float64{
		2, 3, 4, 5,
		6, 7, 8, 9,
		8, 7, 6, 5,
		4, 3, 2, 1,
	})

	assertNotEqualMatrix(t, m1, m2)
}

func TestMultiplyingTwoMatrices(t *testing.T) {
	m1 := NewMatrix(4, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	})
	m2 := NewMatrix(4, 4, []float64{
		-2, 1, 2, 3,
		3, 2, 1, -1,
		4, 3, 6, 5,
		1, 2, 7, 8,
	})
	expected := NewMatrix(4, 4, []float64{
		20, 22, 50, 48,
		44, 54, 114, 108,
		40, 58, 110, 102,
		16, 26, 46, 42,
	})
	actual := m1.Multiply(m2)

	assertEqualMatrix(t, expected, actual)
}

func TestMultiplyingMatrixByTuple(t *testing.T) {
	m1 := NewMatrix(4, 4, []float64{
		1, 2, 3, 4,
		2, 4, 4, 2,
		8, 6, 4, 1,
		0, 0, 0, 1,
	})
	t1 := Tuple{1, 2, 3, 1}
	actual := m1.MutiplyByTuple(t1)
	expected := Tuple{18, 24, 33, 1}

	assertEqualTuple(t, expected, actual)
}

func TestMultiplyingMatrixByIdentityMatrix(t *testing.T) {
	m1 := NewMatrix(4, 4, []float64{
		0, 1, 2, 4,
		1, 2, 4, 8,
		2, 4, 8, 16,
		4, 8, 16, 32,
	})
	actual := m1.Multiply(m1.Identity())
	expected := m1

	assertEqualMatrix(t, expected, actual)
}

// Scenario: Transposing a matrix Given the following matrix A:
// |0|9|3|0| |9|8|0|8| |1|8|5|3| |0|0|5|8|
// Then transpose(A) is the following matrix: |0|9|1|0|
// |9|8|8|0|
// |3|0|5|5|
// |0|8|3|8|

func TestMatrixTransposition(t *testing.T) {
	m1 := NewMatrix(4, 4, []float64{
		0, 9, 3, 0,
		9, 8, 0, 8,
		1, 8, 5, 3,
		0, 0, 5, 8,
	})

	actual := m1.Transpose()
	expected := NewMatrix(4, 4, []float64{
		0, 9, 1, 0,
		9, 8, 8, 0,
		3, 0, 5, 5,
		0, 8, 3, 8,
	})

	assertEqualMatrix(t, expected, actual)
}
