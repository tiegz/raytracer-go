package raytracer

import (
	"fmt"
	"math"
	"testing"
)

func TestARayMissesACylinder(t *testing.T) {
	testCases := []struct {
		Origin    Tuple
		Direction Tuple
	}{
		{NewPoint(1, 0, 0), NewVector(0, 1, 0)},  // on surface of cylinder, pointed up the y axis
		{NewPoint(0, 0, 0), NewVector(0, 1, 0)},  // inside the cylinder, pointed up the y axis
		{NewPoint(0, 0, -5), NewVector(1, 1, 1)}, // outside the cylinder
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Intersection missing from ray at %v", tc.Origin), func(t *testing.T) {
			shape := NewCylinder()
			cyl := shape.LocalShape.(*Cylinder)
			r := NewRay(tc.Origin, tc.Direction.Normalized())
			xs := cyl.LocalIntersect(r, shape)

			assertEqualInt(t, 0, len(xs))
		})
	}
}

func TestARayStrikesACylnder(t *testing.T) {
	testCases := []struct {
		Origin    Tuple
		Direction Tuple
		T1        float64
		T2        float64
	}{
		{NewPoint(1, 0, -5), NewVector(0, 0, 1), 5, 5},                 // strikes the cylinder on a tangent (but will return 2 points anyway)
		{NewPoint(0, 0, -5), NewVector(0, 0, 1), 4, 6},                 //intersects cylnder perpendicularly thru middle
		{NewPoint(0.5, 0, -5), NewVector(0.1, 1, 1), 6.80798, 7.08872}, // skewed, and strikes the cylinder at an angle
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Intersection missing from ray at %v", tc.Origin), func(t *testing.T) {
			shape := NewCylinder()
			cyl := shape.LocalShape.(*Cylinder)
			r := NewRay(tc.Origin, tc.Direction.Normalized())
			xs := cyl.LocalIntersect(r, shape)

			assertEqualInt(t, 2, len(xs))
			assertEqualFloat64(t, tc.T1, xs[0].Time)
			assertEqualFloat64(t, tc.T2, xs[1].Time)
		})
	}
}

func TestNormalVectorOnACylinder(t *testing.T) {
	testCases := []struct {
		Point  Tuple
		Normal Tuple
	}{
		{NewPoint(1, 0, 0), NewVector(1, 0, 0)},
		{NewPoint(0, 5, -1), NewVector(0, 0, -1)},
		{NewPoint(0, -2, 1), NewVector(0, 0, 1)},
		{NewPoint(-1, 1, 0), NewVector(-1, 0, 0)},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Normal at point %v", tc.Point), func(t *testing.T) {
			shape := NewCylinder()
			cyl := shape.LocalShape.(*Cylinder)
			n := cyl.LocalNormalAt(tc.Point, NewIntersection(0, shape))
			assertEqualTuple(t, tc.Normal, n)
		})
	}
}

func TestTheDefaultMinimumAndMaximumForACylinder(t *testing.T) {
	shape := NewCylinder()
	cyl := shape.LocalShape.(*Cylinder)
	assert(t, math.IsInf(cyl.Minimum, -1))
	assert(t, math.IsInf(cyl.Maximum, 1))
}

func TestIntersectingAConstraintedCylinder(t *testing.T) {
	shape := NewCylinder()
	cyl := shape.LocalShape.(*Cylinder)
	cyl.Minimum = 1
	cyl.Maximum = 2
	testCases := []struct {
		Point     Tuple
		Direction Tuple
		Count     int
	}{
		{NewPoint(0, 1.5, 0), NewVector(0.1, 1, 0), 0}, // ray begins inside cylinder, cast diagonally, exits without intersection
		{NewPoint(0, 3, -5), NewVector(0, 0, 1), 0},    // ray perpendicular to y axis, below cylnder, exits without intersection
		{NewPoint(0, 0, -5), NewVector(0, 0, 1), 0},    // ray perpendicular to y axis, above cylinder, exits without intersection
		{NewPoint(0, 2, -5), NewVector(0, 0, 1), 0},    // edge case: min and max are exclusive
		{NewPoint(0, 1, -5), NewVector(0, 0, 1), 0},    // edge case: min and max are exclusive
		{NewPoint(0, 1.5, -2), NewVector(0, 0, 1), 2},  // ray perpendicular to y axis, rays goes thru middle of cylindar, exits with intersection
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Intersecting a constrained cylinder %v, %v", tc.Point, tc.Direction), func(t *testing.T) {
			r := NewRay(tc.Point, tc.Direction.Normalized())
			xs := cyl.LocalIntersect(r, shape)
			assertEqualInt(t, tc.Count, len(xs))
		})
	}
}

func TestTheDefaultClosedValueForACylinder(t *testing.T) {
	shape := NewCylinder()
	cyl := shape.LocalShape.(*Cylinder)
	assert(t, !cyl.Closed)
}

func TestIntersectingTheCapsOfAClosedCylinder(t *testing.T) {
	shape := NewCylinder()
	cyl := shape.LocalShape.(*Cylinder)
	cyl.Minimum = 1
	cyl.Maximum = 2
	cyl.Closed = true
	testCases := []struct {
		Point     Tuple
		Direction Tuple
		Count     int
	}{
		{NewPoint(0, 3, 0), NewVector(0, -1, 0), 2},  // starts above cylinder, points down thru middle along y, intersects both caps
		{NewPoint(0, 3, -2), NewVector(0, -1, 2), 2}, // starts above cylindar, points diagonally through it, intersects both caps
		{NewPoint(0, 4, -2), NewVector(0, -1, 1), 2}, // edge case
		{NewPoint(0, 0, -2), NewVector(0, 1, 2), 2},  // starts below cylindar, points diagonally through it, intersects both caps
		{NewPoint(0, -1, -2), NewVector(0, 1, 1), 2}, // edge case
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Intersecting the caps of a closed cylinder %v", tc.Point), func(t *testing.T) {
			r := NewRay(tc.Point, tc.Direction.Normalized())
			xs := cyl.LocalIntersect(r, shape)
			assertEqualInt(t, tc.Count, len(xs))
		})
	}
}

func TestTheNormalVectorOnACylindersEndCaps(t *testing.T) {
	shape := NewCylinder()
	cyl := shape.LocalShape.(*Cylinder)
	cyl.Minimum = 1
	cyl.Maximum = 2
	cyl.Closed = true
	testCases := []struct {
		Point  Tuple
		Normal Tuple
	}{
		{NewPoint(0, 1, 0), NewVector(0, -1, 0)},   //
		{NewPoint(0.5, 1, 0), NewVector(0, -1, 0)}, //
		{NewPoint(0, 1, 0.5), NewVector(0, -1, 0)}, //
		{NewPoint(0, 2, 0), NewVector(0, 1, 0)},    //
		{NewPoint(0.5, 2, 0), NewVector(0, 1, 0)},  //
		{NewPoint(0, 2, 0.5), NewVector(0, 1, 0)},  //
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Intersecting the caps of a closed cylinder %v", tc.Point), func(t *testing.T) {
			normal := cyl.LocalNormalAt(tc.Point, NewIntersection(0, shape))
			assertEqualTuple(t, tc.Normal, normal)
		})
	}
}

func TestAnUnboundedCylinderHasABoundingBox(t *testing.T) {
	c := NewCylinder()
	b := c.Bounds()

	assertEqualTuple(t, NewPoint(-1, math.Inf(-1), -1), b.MinPoint)
	assertEqualTuple(t, NewPoint(1, math.Inf(1), 1), b.MaxPoint)
}

func TestABoundedCylinderHasABoundingBox(t *testing.T) {
	c := NewCylinder()
	cc := c.LocalShape.(*Cylinder)
	cc.Minimum = -5
	cc.Maximum = 3
	b := c.Bounds()

	assertEqualTuple(t, NewPoint(-1, -5, -1), b.MinPoint)
	assertEqualTuple(t, NewPoint(1, 3, 1), b.MaxPoint)
}
