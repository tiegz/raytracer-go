package raytracer

import (
	"testing"
)

// Helpers

func assertEqualTuple(t *testing.T, expected, actual Tuple) {
	if !expected.IsEqualTo(actual) {
		t.Errorf("Expected %v to be equal to %v, but was not", actual, expected)
	}
}

func assertEqualColor(t *testing.T, expected, actual Color) {
	if !expected.IsEqualTo(actual) {
		t.Errorf("Expected %v to be equal to %v, but was not", actual, expected)
	}
}

func assertNotEqualTuple(t *testing.T, expected, actual Tuple) {
	if expected.IsEqualTo(actual) {
		t.Errorf("Expected %v to not be equal to %v, but was equal", actual, expected)
	}
}

func assertEqualFloat64(t *testing.T, expected float64, actual float64) {
	if expected != actual {
		t.Errorf("Expected value to be %f, but was: %f\n", expected, actual)
	}
}

func assertEqualInt(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Expected %v to be equal to %v, but was not", expected, actual)
	}
}

func assertEqualString(t *testing.T, expected string, actual string) {
  if expected != actual {
    t.Errorf("\nExpected:\n---------\n%v\nActual:\n-------\n%v\n", expected, actual)
  }
}

