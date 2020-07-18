package raytracer

import (
	"fmt"
	"testing"
)

func TestCSVGIsCreatedWithAnOperationAndTwoShapes(t *testing.T) {
	s1 := NewSphere()
	s2 := NewCube()

	shape := NewCsg("union", s1, s2)
	c := shape.LocalShape.(Csg)

	assertEqualString(t, "union", c.Operation)
	assertEqualShape(t, *s1, *c.Left)
	assertEqualShape(t, *s2, *c.Right)
	assertEqualShape(t, *shape, *s1.Parent)
	assertEqualShape(t, *shape, *s2.Parent)
}

func TestEvaluatingTheRuleForACsgOperation(t *testing.T) {
	testCases := []struct {
		op     string
		lhit   bool
		inl    bool
		inr    bool
		result bool
	}{
		{"union", true, true, true, false},
		{"union", true, true, false, true},
		{"union", true, false, true, false},
		{"union", true, false, false, true},
		{"union", false, true, true, false},
		{"union", false, true, false, false},
		{"union", false, false, true, true},
		{"union", false, false, false, true},
		{"intersection", true, true, true, true},
		{"intersection", true, true, false, false},
		{"intersection", true, false, true, true},
		{"intersection", true, false, false, false},
		{"intersection", false, true, true, true},
		{"intersection", false, true, false, true},
		{"intersection", false, false, true, false},
		{"intersection", false, false, false, false},
		{"difference", true, true, true, false},
		{"difference", true, true, false, true},
		{"difference", true, false, true, false},
		{"difference", true, false, false, true},
		{"difference", false, true, true, true},
		{"difference", false, true, false, true},
		{"difference", false, false, true, false},
		{"difference", false, false, false, false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Csg rule truth table %v", tc), func(t *testing.T) {
			result := IntersectionAllowed(tc.op, tc.lhit, tc.inl, tc.inr)
			assert(t, tc.result == result)
		})
	}
}

func TestFilteringAListOfIntersections(t *testing.T) {
	testCases := []struct {
		operation string
		x0        int
		x1        int
	}{
		{"union", 0, 3},
		{"intersection", 1, 2},
		{"difference", 0, 1},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Csg rule truth table %v", tc), func(t *testing.T) {
			s1 := NewSphere()
			s2 := NewCube()
			c := NewCsg(tc.operation, s1, s2)
			csg := c.LocalShape.(Csg)
			xs := Intersections{
				NewIntersection(1, s1),
				NewIntersection(2, s2),
				NewIntersection(3, s1),
				NewIntersection(4, s2),
			}
			result := csg.FilterIntersections(xs)

			assertEqualInt(t, 2, len(result))
			assertEqualIntersection(t, *xs[tc.x0], *result[0])
			assertEqualIntersection(t, *xs[tc.x1], *result[1])
		})
	}
}

func TestARayMissesACsgObject(t *testing.T) {
	sphere := NewSphere()
	cube := NewCube()
	c := NewCsg("union", sphere, cube)
	r := NewRay(NewPoint(0, 2, -5), NewVector(0, 0, 1))

	xs := c.LocalShape.LocalIntersect(r, c)
	assertEqualInt(t, 0, len(xs))
}

func TestARayHitsACsgObject(t *testing.T) {
	s1 := NewSphere()
	s1.Label = "s1"
	s2 := NewSphere()
	s2.Label = "s2"
	s2.SetTransform(NewTranslation(0, 0, 0.5))
	c := NewCsg("union", s1, s2)
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	xs := c.LocalShape.LocalIntersect(r, c)

	assertEqualInt(t, 2, len(xs))
	assertEqualFloat64(t, 4, xs[0].Time)
	assertEqualShape(t, *s1, *xs[0].Object)
	assertEqualFloat64(t, 6.5, xs[1].Time)
	assertEqualShape(t, *s2, *xs[1].Object)
}

func TestCreatingANewCsg(t *testing.T) {
	shape1 := NewCube()
	shape2 := NewSphere()
	s := NewCsg("difference", shape1, shape2)
	g := s.LocalShape.(Csg)

	assertEqualMatrix(t, IdentityMatrix(), s.Transform)
	assertEqualString(t, "difference", g.Operation)

}

// TODO: add more Csg tests similar to Group tests?

func TestACsgShapeHasABoundingBoxThatContainsItseChildren(t *testing.T) {
	left := NewSphere()
	right := NewSphere()
	right.SetTransform(NewTranslation(2, 3, 4))

	csg := NewCsg("difference", left, right)
	box := csg.Bounds()

	assertEqualTuple(t, NewPoint(-1, -1, -1), box.MinPoint)
	assertEqualTuple(t, NewPoint(3, 4, 5), box.MaxPoint)
}

func TestIntersectingRayAndCsgDoesntTestChildrenIfBoxIsMissed(t *testing.T) {
	left := NewTestShape()
	right := NewTestShape()
	s := NewCsg("difference", left, right)
	csg := s.LocalShape.(Csg)
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 2, 0))

	s.Intersect(r)

	assertNil(t, csg.Left.SavedRay)
	assertNil(t, csg.Right.SavedRay)
}

func TestIntersectingRayAndCsgTestsChildrenIfBoxIsHit(t *testing.T) {
	left := NewTestShape()
	right := NewTestShape()
	s := NewCsg("difference", left, right)
	csg := s.LocalShape.(Csg)
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))

	s.Intersect(r)

	assertEqualRay(t, r, *csg.Left.SavedRay)
	assertEqualRay(t, r, *csg.Right.SavedRay)
}
