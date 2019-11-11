package raytracer

import (
	"fmt"
	"math"
	"testing"
)

func TestIntersectingAConeWithARay(t *testing.T) {
	testCases := []struct {
		Origin    Tuple
		Direction Tuple
		T1        float64
		T2        float64
	}{
		{NewPoint(0, 0, -5), NewVector(0, 0, 1), 5, 5},                  //
		{NewPoint(0, 0, -5), NewVector(1, 1, 1), 8.66025, 8.66025},      //
		{NewPoint(1, 1, -5), NewVector(-0.5, -1, 1), 4.55006, 49.44994}, //
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Intersection missing from ray, test #%v", idx), func(t *testing.T) {
			shape := NewCone()
			cone := shape.LocalShape.(*Cone)
			r := NewRay(tc.Origin, tc.Direction.Normalized())
			xs := cone.LocalIntersect(r, &shape)

			assertEqualInt(t, 2, len(xs))
			assertEqualFloat64(t, tc.T1, xs[0].Time)
			assertEqualFloat64(t, tc.T2, xs[1].Time)
		})
	}
}

func TestIntersectingAConeWithARayParallelToOneOfItsHalves(t *testing.T) {
	shape := NewCone()
	cone := shape.LocalShape.(*Cone)
	direction := NewVector(0, 1, 1)
	r := NewRay(NewPoint(0, 0, -1), direction.Normalized())
	xs := cone.LocalIntersect(r, &shape)

	assertEqualInt(t, 1, len(xs))
	assertEqualFloat64(t, 0.35355, xs[0].Time)
}

func TestIntersectingAConesEndCaps(t *testing.T) {
	testCases := []struct {
		Origin    Tuple
		Direction Tuple
		Count     int
	}{
		{NewPoint(0, 0, -5), NewVector(0, 1, 0), 0},    //
		{NewPoint(0, 0, -0.25), NewVector(0, 1, 1), 2}, //
		{NewPoint(0, 0, -0.25), NewVector(0, 1, 0), 4}, //
	}

	shape := NewCone()
	cone := shape.LocalShape.(*Cone)
	cone.Minimum = -0.5
	cone.Maximum = 0.5
	cone.Closed = true

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Intersecta cone's end caps #%v", idx), func(t *testing.T) {
			r := NewRay(tc.Origin, tc.Direction.Normalized())
			xs := cone.LocalIntersect(r, &shape)

			assertEqualInt(t, tc.Count, len(xs))
		})
	}
}

func TestComputingTheNormalVectorOnACone(t *testing.T) {
	shape := NewCone()
	cone := shape.LocalShape.(*Cone)

	testCases := []struct {
		Point  Tuple
		Normal Tuple
	}{
		{NewPoint(0, 0, 0), NewVector(0, 0, 0)},
		{NewPoint(1, 1, 1), NewVector(1, -math.Sqrt(2), 1)},
		{NewPoint(-1, -1, 0), NewVector(-1, 1, 0)},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Normal at point %v", tc.Point), func(t *testing.T) {
			n := cone.LocalNormalAt(tc.Point, NewIntersection(0, shape))
			assertEqualTuple(t, tc.Normal, n)
		})
	}
}

func TestAnUnboundedConeHasABoundingBox(t *testing.T) {
	c := NewCone()
	b := c.Bounds()

	assertEqualTuple(t, NewPoint(math.Inf(-1), math.Inf(-1), math.Inf(-1)), b.MinPoint)
	assertEqualTuple(t, NewPoint(math.Inf(1), math.Inf(1), math.Inf(1)), b.MaxPoint)
}

func TestABoundedConeHasABoundingBox(t *testing.T) {
	c := NewCone()
	cc := c.LocalShape.(*Cone)
	cc.Minimum = -5
	cc.Maximum = 3

	b := c.Bounds()

	assertEqualTuple(t, NewPoint(-5, -5, -5), b.MinPoint)
	assertEqualTuple(t, NewPoint(5, 3, 5), b.MaxPoint)
}
