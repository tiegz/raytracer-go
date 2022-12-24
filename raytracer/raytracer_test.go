package raytracer

import (
	"os"
	"reflect"
	"testing"
)

func TestAltMod(t *testing.T) {
	assertEqualFloat64(t, 0.75, altMod(-0.25, 1.0))
	assertEqualFloat64(t, 2.0, altMod(5.0, 3.0))
}

// Helpers

func expectationFailure(t *testing.T, expected interface{}, actual interface{}) {
	t.Errorf("\nExpected: %v\nActual:   %v\n", expected, actual)
}

func assert(t *testing.T, result bool) {
	if !result {
		expectationFailure(t, true, false)
	}
}

func assertEqualSliceOfStrings(t *testing.T, expected, actual []string) {
	if len(expected) != len(actual) {
		expectationFailure(t, len(expected), len(actual))
	} else {
		for i, v := range expected {
			if v != actual[i] {
				expectationFailure(t, expected, actual)
			}
		}
	}
}

func assertEqualBool(t *testing.T, expected, actual bool) {
	if expected != actual {
		expectationFailure(t, expected, actual)
	}
}

func assertEqualRay(t *testing.T, expected, actual Ray) {
	if !expected.IsEqualTo(&actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertEqualMatrix(t *testing.T, expected, actual Matrix) {
	if !expected.IsEqualTo(&actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertNotEqualMatrix(t *testing.T, expected, actual Matrix) {
	if expected.IsEqualTo(&actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertEqualTuple(t *testing.T, expected, actual Tuple) {
	if !expected.IsEqualTo(actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertEqualBoundingBox(t *testing.T, expected, actual BoundingBox) {
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

func assertEqualUInt(t *testing.T, expected uint, actual uint) {
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
	if !expected.IsEqualTo(&actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertEqualGroup(t *testing.T, expectedShapes []*Shape, actual Shape) {
	g := actual.LocalShape.(Group)
	assertEqualInt(t, len(expectedShapes), len(g.Children))
	for i, _ := range expectedShapes {
		assertEqualShape(t, *expectedShapes[i], *g.Children[i])
	}
}

func assertEqualIntersection(t *testing.T, expected Intersection, actual Intersection) {
	if !expected.IsEqualTo(&actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertEqualLight(t *testing.T, expected AreaLight, actual AreaLight) {
	if !expected.IsEqualTo(&actual) {
		expectationFailure(t, expected, actual)
	}
}

func assertEqualMaterial(t *testing.T, expected Material, actual Material) {
	if !expected.IsEqualTo(&actual) {
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

func assertEqualError(t *testing.T, expected error, actual error) {
	expectedStr := expected.Error()
	actualStr := actual.Error()
	if expectedStr != actualStr {
		expectationFailure(t, expectedStr, actualStr)
	}
}

func assertFileExists(t *testing.T, filepath string) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		t.Errorf("\nExpected %s to exist, but it did not.\n", filepath)
	}
}
