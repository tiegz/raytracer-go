package raytracer

import (
	"math"
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

func TestNewScale(t *testing.T) {
	actual := NewScale(2, 3, 4)
	expected := NewMatrix(4, 4, []float64{
		2, 0, 0, 0,
		0, 3, 0, 0,
		0, 0, 4, 0,
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

// Translation has no effect on vectors.
func TestMultiplyingVectorByTranslationMatrixHasNoEffect(t *testing.T) {
	translation := NewTranslation(5, -3, 2)
	vector := NewVector(-3, 4, 5)
	expected := vector
	actual := translation.MultiplyByTuple(vector)

	assertEqualTuple(t, expected, actual)
}

func TestScalingPointByScaleMatrix(t *testing.T) {
	scale := NewScale(2, 3, 4)
	point := NewPoint(-4, 6, 8)
	expected := NewPoint(-8, 18, 32)
	actual := scale.MultiplyByTuple(point)

	assertEqualTuple(t, expected, actual)
}

// Scaling applies to both vectors and points.
func TestScalingVectorByScaleMatrix(t *testing.T) {
	scale := NewScale(2, 3, 4)
	vector := NewVector(-4, 6, 8)
	expected := NewVector(-8, 18, 32)
	actual := scale.MultiplyByTuple(vector)

	assertEqualTuple(t, expected, actual)
}

func TestScalingVectorByInvertedScaleMatrix(t *testing.T) {
	scale := NewScale(2, 3, 4)
	scale_inverted := scale.Inverse()
	vector := NewVector(-4, 6, 8)
	expected := NewVector(-2, 2, 2)
	actual := scale_inverted.MultiplyByTuple(vector)

	assertEqualTuple(t, expected, actual)
}

func TestReflectingPointByNegativeScaleMatrix(t *testing.T) {
	scale := NewScale(-1, 1, 1) // reflect on X-axis
	point := NewPoint(2, 3, 4)
	expected := NewPoint(-2, 3, 4)
	actual := scale.MultiplyByTuple(point)

	assertEqualTuple(t, expected, actual)
}

func TestRotatingPointAroundXAxis(t *testing.T) {
	point := NewPoint(0, 1, 0)
	half_quarter_rotation := NewRotateX(math.Pi / 4)
	full_quarter_rotation := NewRotateX(math.Pi / 2)

	assertEqualTuple(t, NewPoint(0, math.Sqrt(2)/2, math.Sqrt(2)/2), half_quarter_rotation.MultiplyByTuple(point))
	assertEqualTuple(t, NewPoint(0, 0, 1), full_quarter_rotation.MultiplyByTuple(point))
}

func TestRotatingPointAroundXAxisInOppositeDirection(t *testing.T) {
	point := NewPoint(0, 1, 0)
	half_quarter_rotation := NewRotateX(math.Pi / 4)
	half_quarter_rotation_inverted := half_quarter_rotation.Inverse()

	assertEqualTuple(t, NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2), half_quarter_rotation_inverted.MultiplyByTuple(point))
}

func TestRotatingPointAroundYAxis(t *testing.T) {
	point := NewPoint(0, 0, 1)
	half_quarter_rotation := NewRotateY(math.Pi / 4)
	full_quarter_rotation := NewRotateY(math.Pi / 2)

	assertEqualTuple(t, NewPoint(math.Sqrt(2)/2, 0, math.Sqrt(2)/2), half_quarter_rotation.MultiplyByTuple(point))
	assertEqualTuple(t, NewPoint(1, 0, 0), full_quarter_rotation.MultiplyByTuple(point))
}

func TestRotatingPointAroundZAxis(t *testing.T) {
	point := NewPoint(0, 1, 0)
	half_quarter_rotation := NewRotateZ(math.Pi / 4)
	full_quarter_rotation := NewRotateZ(math.Pi / 2)

	assertEqualTuple(t, NewPoint(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0), half_quarter_rotation.MultiplyByTuple(point))
	assertEqualTuple(t, NewPoint(-1, 0, 0), full_quarter_rotation.MultiplyByTuple(point))
}

// TODO table-based tests for these 6 shearing examples
func TestShearingPointXByProportionToY(t *testing.T) {
	point := NewPoint(2, 3, 4)
	shearing := NewShear(1, 0, 0, 0, 0, 0)

	assertEqualTuple(t, NewPoint(5, 3, 4), shearing.MultiplyByTuple(point))
}

func TestShearingPointXByProportionToZ(t *testing.T) {
	point := NewPoint(2, 3, 4)
	shearing := NewShear(0, 1, 0, 0, 0, 0)

	assertEqualTuple(t, NewPoint(6, 3, 4), shearing.MultiplyByTuple(point))
}

func TestShearingPointYByProportionToX(t *testing.T) {
	point := NewPoint(2, 3, 4)
	shearing := NewShear(0, 0, 1, 0, 0, 0)

	assertEqualTuple(t, NewPoint(2, 5, 4), shearing.MultiplyByTuple(point))
}
func TestShearingPointYByProportionToZ(t *testing.T) {
	point := NewPoint(2, 3, 4)
	shearing := NewShear(0, 0, 0, 1, 0, 0)

	assertEqualTuple(t, NewPoint(2, 7, 4), shearing.MultiplyByTuple(point))
}

func TestShearingPointZByProportionToX(t *testing.T) {
	point := NewPoint(2, 3, 4)
	shearing := NewShear(0, 0, 0, 0, 1, 0)

	assertEqualTuple(t, NewPoint(2, 3, 6), shearing.MultiplyByTuple(point))
}

func TestShearingPointZByProportionToY(t *testing.T) {
	point := NewPoint(2, 3, 4)
	shearing := NewShear(0, 0, 0, 0, 0, 1)

	assertEqualTuple(t, NewPoint(2, 3, 7), shearing.MultiplyByTuple(point))
}

// TODO try to write "fluent api" for method chaining (pg 55)
func TestChainedTransformationAreAppliedInSequence(t *testing.T) {
	point := NewPoint(1, 0, 1)
	rotation := NewRotateX(math.Pi / 2) // 90 deg rotation
	scale := NewScale(5, 5, 5)
	translation := NewTranslation(10, 5, 7)

	chained_transformations := translation.Multiply(scale)
	chained_transformations = chained_transformations.Multiply(rotation)
	assertEqualTuple(t, NewPoint(15, 0, 7), chained_transformations.MultiplyByTuple(point))
}

func TestTransformationMatrixForDefaultOrientation(t *testing.T) {
	from := NewPoint(0, 0, 0)
	to := NewPoint(0, 0, -1)
	up := NewVector(0, 1, 0)

	actual := NewViewTransform(from, to, up)
	expected := IdentityMatrix()

	assertEqualMatrix(t, expected, actual)
}

// Aka Z-axis reflection
func TestViewTransformationMatrixLookingInPositiveZDirection(t *testing.T) {
	from := NewPoint(0, 0, 0)
	to := NewPoint(0, 0, 1) // world is behind us (?)
	up := NewVector(0, 1, 0)

	actual := NewViewTransform(from, to, up)
	expected := NewScale(-1, 1, -1)

	assertEqualMatrix(t, expected, actual)
}

func TestViewTransformationMovesTheWorld(t *testing.T) {
	from := NewPoint(0, 0, 8)
	to := NewPoint(0, 0, 0)
	up := NewVector(0, 1, 0)

	actual := NewViewTransform(from, to, up)
	expected := NewTranslation(0, 0, -8)

	assertEqualMatrix(t, expected, actual)
}

func TestArbitraryViewTransformation(t *testing.T) {
	from := NewPoint(1, 3, 2)
	to := NewPoint(4, -2, 8)
	up := NewVector(1, 1, 0)

	actual := NewViewTransform(from, to, up)
	expected := NewMatrix(4, 4, []float64{
		-0.50709, 0.50709, 0.67612, -2.36643,
		0.76772, 0.60609, 0.12122, -2.82843,
		-0.35857, 0.59761, -0.71714, 0.00000,
		0.00000, 0.00000, 0.00000, 1.00000,
	})

	assertEqualMatrix(t, expected, actual)
}
