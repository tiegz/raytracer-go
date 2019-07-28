package raytracer

import (
	"fmt"
	"testing"
)

func TestARayIntersectsACube(t *testing.T) {
	testCases := []struct {
		Name      string
		Origin    Tuple
		Direction Tuple
		T1        float64
		T2        float64
	}{
		{"+x", NewPoint(5, 0.5, 0), NewVector(-1, 0, 0), 4, 6},
		{"-x", NewPoint(-5, 0.5, 0), NewVector(1, 0, 0), 4, 6},
		{"+y", NewPoint(0.5, 5, 0), NewVector(0, -1, 0), 4, 6},
		{"-y", NewPoint(0.5, -5, 0), NewVector(0, 1, 0), 4, 6},
		{"+z", NewPoint(0.5, 0, 5), NewVector(0, 0, -1), 4, 6},
		{"-z", NewPoint(0.5, 0, -5), NewVector(0, 0, 1), 4, 6},
		{"inside", NewPoint(0, 0.5, 0), NewVector(0, 0, 1), -1, 1},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Intersection on %s face", tc.Name), func(t *testing.T) {
			shape := NewCube()
			cube := shape.LocalShape.(*Cube)
			r := NewRay(tc.Origin, tc.Direction)
			xs := cube.LocalIntersect(r, &shape)

			assertEqualInt(t, 2, len(xs))
			assertEqualFloat64(t, tc.T1, xs[0].Time)
			assertEqualFloat64(t, tc.T2, xs[1].Time)
		})
	}
}

func TestARayMissesACube(t *testing.T) {
	testCases := []struct {
		Origin    Tuple
		Direction Tuple
	}{
		{NewPoint(-2, 0, 0), NewVector(0.2673, 0.5345, 0.8018)},
		{NewPoint(0, -2, 0), NewVector(1, 0, 0)},
		{NewPoint(0, 0, -2), NewVector(0, -1, 0)},
		{NewPoint(2, 0, 2), NewVector(0, 1, 0)},
		{NewPoint(0, 2, 2), NewVector(0, 0, -1)},
		{NewPoint(2, 2, 0), NewVector(0, 0, 1)},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Intersection that misses from %v", tc.Origin), func(t *testing.T) {
			shape := NewCube()
			cube := shape.LocalShape.(*Cube)
			r := NewRay(tc.Origin, tc.Direction)
			xs := cube.LocalIntersect(r, &shape)

			assertEqualInt(t, 0, len(xs))
		})
	}
}

func TestTheNormalOnTheSurfaceOfACube(t *testing.T) {
	testCases := []struct {
		Point          Tuple
		ExpectedNormal Tuple
	}{
		{NewPoint(1, 0.5, -0.8), NewVector(1, 0, 0)},
		{NewPoint(-1, -0.2, 0.9), NewVector(-1, 0, 0)},
		{NewPoint(-0.4, 1, -0.1), NewVector(0, 1, 0)},
		{NewPoint(0.3, -1, -0.7), NewVector(0, -1, 0)},
		{NewPoint(-0.6, 0.3, 1), NewVector(0, 0, 1)},
		{NewPoint(0.4, 0.4, -1), NewVector(0, 0, -1)},
		{NewPoint(1, 1, 1), NewVector(1, 0, 0)},
		{NewPoint(-1, -1, -1), NewVector(-1, 0, 0)},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Normal on the surface at %v", tc.Point), func(t *testing.T) {
			shape := NewCube()
			cube := shape.LocalShape.(*Cube)
			normal := cube.LocalNormalAt(tc.Point)

			assertEqualTuple(t, tc.ExpectedNormal, normal)
		})
	}
}
