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

//   Scenario: point() creates tuples with w=1
//     Given p point(4, -4, 3)
//     Then p = tuple(4, -4, 3, 1)

//   Scenario: vector() creates tuples with w=0
//     Given v vector(4, -4, 3)
//     Then v = tuple(4, -4, 3, 0)

//   Scenario: Adding two tuples
//     Given a1 tuple(3, -2, 5, 1)
//     And a2 tuple(-2, 3, 1, 0)
//     Then a1 + a2 = tuple(1, 1, 6, 1)
