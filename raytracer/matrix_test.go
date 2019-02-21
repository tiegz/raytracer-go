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
	actual := m1.Multiply(IdentityMatrix())
	expected := m1

	assertEqualMatrix(t, expected, actual)
}
func TestTransposingMatrix(t *testing.T) {
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

func TestMatrixTransposingIdentityMatrix(t *testing.T) {
	m1 := IdentityMatrix()

	assertEqualMatrix(t, m1, m1.Transpose())
}

func TestMatrixDeterminant(t *testing.T) {
	m1 := NewMatrix(2, 2, []float64{
		1, 5,
		-3, 2,
	})

	assertEqualFloat64(t, 17, m1.Determinant())
}

func TestSubmatrixOfThreeByThreeMatrix(t *testing.T) {
	m1 := NewMatrix(3, 3, []float64{
		1, 5, 0,
		-3, 2, 7,
		0, 6, -3,
	})

	expected := NewMatrix(2, 2, []float64{
		-3, 2,
		0, 6,
	})
	actual := m1.Submatrix(0, 2)

	assertEqualMatrix(t, expected, actual)
}

func TestSubmatrixOfFourByFourMatrix(t *testing.T) {
	m1 := NewMatrix(4, 4, []float64{
		-6, 1, 1, 6,
		-8, 5, 8, 6,
		-1, 0, 8, 2,
		-7, 1, -1, 1,
	})
	expected := NewMatrix(3, 3, []float64{
		-6, 1, 6,
		-8, 8, 6,
		-7, -1, 1,
	})
	actual := m1.Submatrix(2, 1)

	assertEqualMatrix(t, expected, actual)
}

func TestCalculatingMinorOfThreeByThreeMatrix(t *testing.T) {
	m1 := NewMatrix(3, 3, []float64{
		3, 5, 0,
		2, -1, -7,
		6, -1, 5,
	})
	sm1 := m1.Submatrix(1, 0)

	assertEqualFloat64(t, 25, sm1.Determinant())
	assertEqualFloat64(t, 25, m1.Minor(1, 0))
}

func TestCalculatingCofactorOfThreeByThreeMatrix(t *testing.T) {
	m1 := NewMatrix(3, 3, []float64{
		3, 5, 0,
		2, -1, -7,
		6, -1, 5,
	})

	assertEqualFloat64(t, -12, m1.Minor(0, 0))
	assertEqualFloat64(t, -12, m1.Cofactor(0, 0))
	assertEqualFloat64(t, 25, m1.Minor(1, 0))
	assertEqualFloat64(t, -25, m1.Cofactor(1, 0))
}

func TestCalculatingDeterminantOfThreeByThreeMatrix(t *testing.T) {
	m := NewMatrix(3, 3, []float64{
		1, 2, 6,
		-5, 8, -4,
		2, 6, 4,
	})

	assertEqualFloat64(t, 56, m.Cofactor(0, 0))
	assertEqualFloat64(t, 12, m.Cofactor(0, 1))
	assertEqualFloat64(t, -46, m.Cofactor(0, 2))
	assertEqualFloat64(t, -196, m.Determinant())
}

func TestCalculatingDeterminantOfFourByFourMatrix(t *testing.T) {
	m := NewMatrix(4, 4, []float64{
		-2, -8, 3, 5,
		-3, 1, 7, 3,
		1, 2, -9, 6,
		-6, 7, 7, -9,
	})

	assertEqualFloat64(t, 690, m.Cofactor(0, 0))
	assertEqualFloat64(t, 447, m.Cofactor(0, 1))
	assertEqualFloat64(t, 210, m.Cofactor(0, 2))
	assertEqualFloat64(t, 51, m.Cofactor(0, 3))
	assertEqualFloat64(t, -4071, m.Determinant())
}
