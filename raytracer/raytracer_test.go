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



func assertEqualFloat64(t *testing.T, expected float64, actual float64) {
  if (expected != actual) {
    t.Errorf("Expected value to be %f, but was: %f\n", expected, actual)
  }
}


//   Scenario: A tuple with w=0 is a vector
//     Given a tuple (4.3, -4.2, 3.1, 0.0)
//     Then a.x = 4.3
//     And a.y = -4.2
//     And a.z = 3.1
//     And a.w = 0.0
//     And a is a vector
//     And a is not a point

func assertEqualTuple(t *testing.T, expected, actual Tuple) {
  if (!expected.IsEqualTo(actual)) {
    t.Errorf("Expected %v to be equal to %v, but was not", expected, actual)
  }
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
