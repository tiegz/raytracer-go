package raytracer

import (
	"reflect"
	"testing"
)

// Helpers

func expectationFailure(t *testing.T, expected interface{}, actual interface{}) {
	t.Errorf("\nExpected: %v\nActual:   %v\n", expected, actual)
}

func assert(t *testing.T, result bool) {
	if !result {
		expectationFailure(t, true, false)
	}
}

func assertEqualMatrix(t *testing.T, expected, actual Matrix) {
	if !expected.IsEqualTo(actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertNotEqualMatrix(t *testing.T, expected, actual Matrix) {
	if expected.IsEqualTo(actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertEqualTuple(t *testing.T, expected, actual Tuple) {
	if !expected.IsEqualTo(actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertEqualColor(t *testing.T, expected, actual Color) {
	if !expected.IsEqualTo(actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertNotEqualTuple(t *testing.T, expected, actual Tuple) {
	if expected.IsEqualTo(actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertEqualFloat64(t *testing.T, expected float64, actual float64) {
	if !equalFloat64s(expected, actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertEqualInt(t *testing.T, expected int, actual int) {
	if expected != actual {
		expectationFailure(t, expected, actual)
	}
}

func assertEqualString(t *testing.T, expected string, actual string) {
	if expected != actual {
		expectationFailure(t, expected, actual)
	}
}

func assertEqualShape(t *testing.T, expected Shape, actual Shape) {
	if !expected.IsEqualTo(actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertEqualIntersection(t *testing.T, expected Intersection, actual Intersection) {
	if !expected.IsEqualTo(actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertEqualPointLight(t *testing.T, expected PointLight, actual PointLight) {
	if !expected.IsEqualTo(actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertEqualMaterial(t *testing.T, expected Material, actual Material) {
	if !expected.IsEqualTo(actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertNil(t *testing.T, object interface{}) {
	if object != nil {
		val := reflect.ValueOf(object)
		if !val.IsNil() {
			expectationFailure(t, nil, object)
		}
	}
}
