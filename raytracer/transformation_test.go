package raytracer

import (
	"testing"
)

func TestNewTranslation(t *testing.T) {
	actual := NewTranslation(5, -3, 2)
	expected := NewMatrix(4, 4, []float64{
		1, 0, 0, 5,
		0, 1, 0, -3,
		0, 0, 1, 2,
		0, 0, 0, 1,
	})
	assertEqualMatrix(t, expected, actual)
}

func TestMultiplyingPointByTranslationMatrix(t *testing.T) {
	translation := NewTranslation(5, -3, 2)
	point := NewPoint(-3, 4, 5)
	expected := NewPoint(2, 1, 7)
	actual := translation.MultiplyByTuple(point)

	assertEqualTuple(t, expected, actual)
}

func TestMultiplyingPointByInvertedTranslationMatrix(t *testing.T) {
	translation := NewTranslation(5, -3, 2)
	translation_inverse := translation.Inverse()
	point := NewPoint(-3, 4, 5)

	expected := NewPoint(-8, 7, 3)
	actual := translation_inverse.MultiplyByTuple(point)

	assertEqualTuple(t, expected, actual)
}

func TestMultiplyingVectorByTranslationMatrixHasNoEffect(t *testing.T) {
	matrix := NewTranslation(5, -3, 2)
	vector := NewVector(-3, 4, 5)
	expected := vector
	actual := matrix.MultiplyByTuple(vector)

	assertEqualTuple(t, expected, actual)
}

// func Test(t *testing.T) {
// // TEMPLATE
// }
