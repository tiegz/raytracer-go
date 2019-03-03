package raytracer

import (
	"math"
	"testing"
)

func TestTuple(t *testing.T) {
	tuple := Tuple{4.3, -4.2, 3.1, 1.0}

	assertEqualFloat64(t, tuple.X, 4.3)
	assertEqualFloat64(t, tuple.Y, -4.2)
	assertEqualFloat64(t, tuple.Z, 3.1)
	assertEqualFloat64(t, tuple.W, 1.0)
}

func TestTupleIsVector(t *testing.T) {
	tuple := Tuple{4.3, -4.2, 3.1, 0.0}

	assertEqualFloat64(t, tuple.X, 4.3)
	assertEqualFloat64(t, tuple.Y, -4.2)
	assertEqualFloat64(t, tuple.Z, 3.1)
	assertEqualFloat64(t, tuple.W, 0.0)
}

func TestNewPointFunctionCreatesTuples(t *testing.T) {
	tuple := NewPoint(4, -4, 3)

	assertEqualTuple(t, tuple, Tuple{4, -4, 3, 1})
}

func TestNewVectorFunctionCreatesTuples(t *testing.T) {
	tuple := NewVector(4, -4, 3)

	assertEqualTuple(t, tuple, Tuple{4, -4, 3, 0})
}

func TestIsEqualTo(t *testing.T) {
	assertEqualTuple(t, Tuple{1, 1, 1, 1}, Tuple{1.0000001, 1.0000001, 1.0000001, 1.0000001})
	assertEqualTuple(t, Tuple{1, 1, 1, 1}, Tuple{1.000001, 1.000001, 1.000001, 1.000001})

	assertNotEqualTuple(t, Tuple{1, 1, 1, 1}, Tuple{1.00001, 1.00001, 1.00001, 1.00001})
	assertNotEqualTuple(t, Tuple{1, 1, 1, 1}, Tuple{1.0001, 1.0001, 1.0001, 1.0001})
	assertNotEqualTuple(t, Tuple{1, 1, 1, 1}, Tuple{1.001, 1.001, 1.001, 1.001})
}

func TestAddingTwoTuples(t *testing.T) {
	t1 := Tuple{3, -2, 5, 1}
	t2 := Tuple{-2, 3, 1, 0}
	expected := Tuple{1, 1, 6, 1}
	actual := t1.Add(t2)

	assertEqualTuple(t, expected, actual)
}

func TestSubtractingTwoPoints(t *testing.T) {
	t1 := NewPoint(3, 2, 1)
	t2 := NewPoint(5, 6, 7)
	expected := NewVector(-2, -4, -6)
	actual := t1.Subtract(t2)

	assertEqualTuple(t, expected, actual)
}

func TestSubtractingVectorFromPoint(t *testing.T) {
	t1 := NewPoint(3, 2, 1)
	t2 := NewVector(5, 6, 7)
	expected := NewPoint(-2, -4, -6)
	actual := t1.Subtract(t2)

	assertEqualTuple(t, expected, actual)
}

func TestSubtractingTwoVectors(t *testing.T) {
	t1 := NewVector(3, 2, 1)
	t2 := NewVector(5, 6, 7)
	expected := NewVector(-2, -4, -6)
	actual := t1.Subtract(t2)

	assertEqualTuple(t, expected, actual)
}

func TestSubtractingVectorFromZeroVector(t *testing.T) {
	t1 := NewVector(0, 0, 0)
	t2 := NewVector(1, -2, 3)
	expected := NewVector(-1, 2, -3)
	actual := t1.Subtract(t2)

	assertEqualTuple(t, expected, actual)
}

func TestNegatingTuple(t *testing.T) {
	t1 := Tuple{1, -2, 3, -4}
	expected := Tuple{-1, 2, -3, 4}
	actual := t1.Negate()

	assertEqualTuple(t, expected, actual)
}

func TestMultiplyingTupleByScalar(t *testing.T) {
	t1 := Tuple{1, -2, 3, -4}
	expected := Tuple{3.5, -7, 10.5, -14}
	actual := t1.Multiply(3.5)

	assertEqualTuple(t, expected, actual)
}

func TestMultiplyingTupleByFraction(t *testing.T) {
	t1 := Tuple{1, -2, 3, -4}
	expected := Tuple{0.5, -1, 1.5, -2}
	actual := t1.Multiply(0.5)

	assertEqualTuple(t, expected, actual)
}

func TestDividingTupleByScalar(t *testing.T) {
	t1 := Tuple{1, -2, 3, -4}
	expected := Tuple{0.5, -1, 1.5, -2}
	actual := t1.Divide(2)

	assertEqualTuple(t, expected, actual)
}

var magnitudeTests = []struct {
	tuple     Tuple
	magnitude float64
}{
	{NewVector(0, 0, 0), 0},
	{NewVector(1, 0, 0), 1},
	{NewVector(0, 1, 0), 1},
	{NewVector(1, 2, 3), math.Sqrt(14)},
	{NewVector(-1, -2, -3), math.Sqrt(14)},
}

func TestComputingMagnitudeOfVectors(t *testing.T) {
	for _, tc := range magnitudeTests {
		assertEqualFloat64(t, tc.magnitude, tc.tuple.Magnitude())
	}
}

var normalizingTests = []struct {
	tuple           Tuple
	normalizedTuple Tuple
}{
	{NewVector(4, 0, 0), NewVector(1, 0, 0)},
	{NewVector(1, 2, 3), NewVector(0.26726, 0.53452, 0.80178)},
}

func TestNormalizingVectors(t *testing.T) {
	for _, tc := range normalizingTests {
		assertEqualTuple(t, tc.normalizedTuple, tc.tuple.Normalized())
	}
}

func TestMagnitudeOfNormalizedVector(t *testing.T) {
	t1 := NewVector(1, 2, 3)
	normalizedT1 := t1.Normalized()

	assertEqualFloat64(t, 1, normalizedT1.Magnitude())
}

func TestDotProductOfTwoVectors(t *testing.T) {
	t1 := NewVector(1, 2, 3)
	t2 := NewVector(2, 3, 4)

	assertEqualFloat64(t, 20, t1.Dot(t2))
}

func TestCrossProductOfTwoVectors(t *testing.T) {
	t1 := NewVector(1, 2, 3)
	t2 := NewVector(2, 3, 4)

	assertEqualTuple(t, NewVector(-1, 2, -1), t1.Cross(t2))
	assertEqualTuple(t, NewVector(1, -2, 1), t2.Cross(t1))
}
func TestReflectVectorAt45Degrees(t *testing.T) {
	t1 := NewVector(1, -1, 0)
	normal := NewVector(0, 1, 0)
	actual := t1.Reflect(normal)
	expected := NewVector(1, 1, 0)

	assertEqualTuple(t, expected, actual)
}

func TestReflectVectorAtSlantedSurface(t *testing.T) {
	t1 := NewVector(0, -1, 0)
	normal := NewVector(math.Sqrt(2)/2, math.Sqrt(2)/2, 0)
	actual := t1.Reflect(normal)
	expected := NewVector(1, 0, 0)

	assertEqualTuple(t, expected, actual)
}
