package raytracer

import (
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

func TestIsEqualTo(t *testing.T) {
  assertEqualTuple(t, Tuple{1, 1, 1, 1}, Tuple{1.0000001, 1.0000001, 1.0000001, 1.0000001})
  assertEqualTuple(t, Tuple{1, 1, 1, 1}, Tuple{1.000001, 1.000001, 1.000001, 1.000001})

  assertNotEqualTuple(t, Tuple{1, 1, 1, 1}, Tuple{1.00001, 1.00001, 1.00001, 1.00001})
  assertNotEqualTuple(t, Tuple{1, 1, 1, 1}, Tuple{1.0001, 1.0001, 1.0001, 1.0001})
  assertNotEqualTuple(t, Tuple{1, 1, 1, 1}, Tuple{1.001, 1.001, 1.001, 1.001})
}

// Helpers

func assertEqualTuple(t *testing.T, expected, actual Tuple) {
	if !expected.IsEqualTo(actual) {
		t.Errorf("Expected %v to be equal to %v, but was not", expected, actual)
	}
}

func assertNotEqualTuple(t *testing.T, expected, actual Tuple) {
  if expected.IsEqualTo(actual) {
    t.Errorf("Expected %v to not be equal to %v, but was equal", expected, actual)
  }
}

func assertEqualFloat64(t *testing.T, expected float64, actual float64) {
	if expected != actual {
		t.Errorf("Expected value to be %f, but was: %f\n", expected, actual)
	}
}
