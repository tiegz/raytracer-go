package raytracer

import (
	"testing"
)

// Helpers

func assert(t *testing.T, result bool) {
	if !result {
		t.Errorf("\nExpected %v to be true, but was false.", result)
	}
}

func assertEqualMatrix(t *testing.T, expected, actual Matrix) {
	if !expected.IsEqualTo(actual) {
		t.Errorf("\nExpected\n---------\n%v\nTo equal\n-------\n%v\n", expected, actual)
	}
}

func assertNotEqualMatrix(t *testing.T, expected, actual Matrix) {
	if expected.IsEqualTo(actual) {
		t.Errorf("\nExpected\n---------\n%v\nTo not equal\n-------\n%v\n", expected, actual)
	}
}

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
	if !equalFloat64s(expected, actual) {
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

func assertEqualObject(t *testing.T, expected Sphere, actual Sphere) {
	if !expected.IsEqualTo(actual) {
		t.Errorf("\nExpected:\n---------\n%v\nActual:\n-------\n%v\n", expected, actual)
	}
}

func assertEqualIntersection(t *testing.T, expected Intersection, actual Intersection) {
	if !expected.IsEqualTo(actual) {
		t.Errorf("\nExpected:\n---------\n%v\nActual:\n-------\n%v\n", expected, actual)
	}
}

func assertEqualMaterial(t *testing.T, expected Material, actual Material) {
	if !expected.IsEqualTo(actual) {
		t.Errorf("\nExpected:\n---------\n%v\nActual:\n-------\n%v\n", expected, actual)
	}
}

func assertNil(t *testing.T, object interface{}) {
	if object != nil {
		t.Errorf("\nExpected object to be nil, but wasn't:\n%v\n", object)
	}
}
