package raytracer

import (
	"fmt"
	"math"
	"testing"
)

func TestCreatingAnEmptyBoundingBox(t *testing.T) {
	bb := NullBoundingBox()

	assertEqualTuple(t, NewPoint(math.Inf(1), math.Inf(1), math.Inf(1)), bb.MinPoint)
	assertEqualTuple(t, NewPoint(math.Inf(-1), math.Inf(-1), math.Inf(-1)), bb.MaxPoint)
}

func TestCreatingABoundingBoxWithVolume(t *testing.T) {
	bb := NewBoundingBox(NewPoint(-1, -2, -3), NewPoint(3, 2, 1))

	assertEqualTuple(t, NewPoint(-1, -2, -3), bb.MinPoint)
	assertEqualTuple(t, NewPoint(3, 2, 1), bb.MaxPoint)

}

func TestAddingPointsToAnEmptyBoundingBox(t *testing.T) {
	bb := NullBoundingBox()
	p1 := NewPoint(-5, 2, 0)
	p2 := NewPoint(7, 0, -3)

	bb.AddPoints(p1, p2)

	assertEqualTuple(t, NewPoint(-5, 0, -3), bb.MinPoint)
	assertEqualTuple(t, NewPoint(7, 2, 0), bb.MaxPoint)
}

func TestAddingOneBoundingBoxToAnother(t *testing.T) {
	bb1 := NewBoundingBox(NewPoint(-5, -2, 0), NewPoint(7, 4, 4))
	bb2 := NewBoundingBox(NewPoint(8, -7, -2), NewPoint(14, 2, 8))

	bb1.AddBoundingBoxes(bb2)

	assertEqualTuple(t, NewPoint(-5, -7, -2), bb1.MinPoint)
	assertEqualTuple(t, NewPoint(14, 4, 8), bb1.MaxPoint)
}

func TestCheckingToSeeIfABoxContainsAGivenPoint(t *testing.T) {
	b := NewBoundingBox(NewPoint(5, -2, 0), NewPoint(11, 4, 7))
	var testCases = []struct {
		Point  Tuple
		Result bool
	}{
		{NewPoint(5, -2, 0), true},
		{NewPoint(11, 4, 7), true},
		{NewPoint(8, 1, 3), true},
		{NewPoint(3, 0, 3), false},
		{NewPoint(8, -4, 3), false},
		{NewPoint(8, 1, -1), false},
		{NewPoint(13, 1, 3), false},
		{NewPoint(8, 5, 3), false},
		{NewPoint(8, 1, 8), false},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test #%v", idx), func(t *testing.T) {
			assertEqualBool(t, b.ContainsPoint(tc.Point), tc.Result)
		})
	}
}

func TestCheckingToSeeIfABoxContainsAGivenBox(t *testing.T) {
	b := NewBoundingBox(NewPoint(5, -2, 0), NewPoint(11, 4, 7))
	var testCases = []struct {
		Min    Tuple
		Max    Tuple
		Result bool
	}{
		{NewPoint(5, -2, 0), NewPoint(11, 4, 7), true},
		{NewPoint(6, -1, 1), NewPoint(10, 3, 6), true},
		{NewPoint(4, -3, -1), NewPoint(10, 3, 6), false},
		{NewPoint(6, -1, 1), NewPoint(12, 5, 8), false},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test #%v", idx), func(t *testing.T) {
			assertEqualBool(t, b.ContainsBox(NewBoundingBox(tc.Min, tc.Max)), tc.Result)
		})
	}
}

func TestTransformingABoundingBox(t *testing.T) {
	b := NewBoundingBox(NewPoint(-1, -1, -1), NewPoint(1, 1, 1))
	matrix := NewRotateX(math.Pi / 4)
	matrix = matrix.Multiply(NewRotateY(math.Pi / 4))
	b2 := b.Transform(matrix)

	// HACK: these are slightly different than the bonus chapter's:
	// NewPoint(-1.4142, -1.7071, -1.7071)
	// NewPoint(1.4142, 1.7071, 1.7071)
	assertEqualTuple(t, NewPoint(-1.41421, -1.70711, -1.70711), b2.MinPoint)
	assertEqualTuple(t, NewPoint(1.41421, 1.70711, 1.70711), b2.MaxPoint)
}

func TestIntersectingARayWithABoundingBoxAtTheOrigin(t *testing.T) {
	b := NewBoundingBox(NewPoint(-1, -1, -1), NewPoint(1, 1, 1))

	var testCases = []struct {
		Origin    Tuple
		Direction Tuple
		Result    bool
	}{
		{NewPoint(5, 0.5, 0), NewVector(-1, 0, 0), true},
		{NewPoint(-5, 0.5, 0), NewVector(1, 0, 0), true},
		{NewPoint(0.5, 5, 0), NewVector(0, -1, 0), true},
		{NewPoint(0.5, -5, 0), NewVector(0, 1, 0), true},
		{NewPoint(0.5, 0, 5), NewVector(0, 0, -1), true},
		{NewPoint(0.5, 0, -5), NewVector(0, 0, 1), true},
		{NewPoint(0, 0.5, 0), NewVector(0, 0, 1), true},
		{NewPoint(-2, 0, 0), NewVector(2, 4, 6), false},
		{NewPoint(0, -2, 0), NewVector(6, 2, 4), false},
		{NewPoint(0, 0, -2), NewVector(4, 6, 2), false},
		{NewPoint(2, 0, 2), NewVector(0, 0, -1), false},
		{NewPoint(0, 2, 2), NewVector(0, -1, 0), false},
		{NewPoint(2, 2, 0), NewVector(-1, 0, 0), false},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test #%v", idx), func(t *testing.T) {
			r := NewRay(tc.Origin, tc.Direction.Normalized())

			assertEqualBool(t, tc.Result, b.Intersects(r))
		})
	}
}
func TestIntersectingARayWithANonCubicBoundingBox(t *testing.T) {
	b := NewBoundingBox(NewPoint(5, -2, 0), NewPoint(11, 4, 7))

	var testCases = []struct {
		Origin    Tuple
		Direction Tuple
		Result    bool
	}{
		{NewPoint(15, 1, 2), NewVector(-1, 0, 0), true},
		{NewPoint(-5, -1, 4), NewVector(1, 0, 0), true},
		{NewPoint(7, 6, 5), NewVector(0, -1, 0), true},
		{NewPoint(9, -5, 6), NewVector(0, 1, 0), true},
		{NewPoint(8, 2, 12), NewVector(0, 0, -1), true},
		{NewPoint(6, 0, -5), NewVector(0, 0, 1), true},
		{NewPoint(8, 1, 3.5), NewVector(0, 0, 1), true},
		{NewPoint(9, -1, -8), NewVector(2, 4, 6), false},
		{NewPoint(8, 3, -4), NewVector(6, 2, 4), false},
		{NewPoint(9, -1, -2), NewVector(4, 6, 2), false},
		{NewPoint(4, 0, 9), NewVector(0, 0, -1), false},
		{NewPoint(8, 6, -1), NewVector(0, -1, 0), false},
		{NewPoint(12, 5, 4), NewVector(-1, 0, 0), false},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test #%v", idx), func(t *testing.T) {
			r := NewRay(tc.Origin, tc.Direction.Normalized())

			assertEqualBool(t, tc.Result, b.Intersects(r))
		})
	}
}
