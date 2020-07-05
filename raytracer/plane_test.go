package raytracer

import (
	"math"
	"testing"
)

func TestNewPlane(t *testing.T) {
	plane := NewPlane()
	localPlane := plane.LocalShape.(*Plane)

	n1 := localPlane.LocalNormalAt(NewPoint(0, 0, 0), NewIntersection(0, plane))
	n2 := localPlane.LocalNormalAt(NewPoint(10, 0, -10), NewIntersection(0, plane))
	n3 := localPlane.LocalNormalAt(NewPoint(-5, 0, 150), NewIntersection(0, plane))

	assertEqualTuple(t, NewVector(0, 1, 0), n1)
	assertEqualTuple(t, NewVector(0, 1, 0), n2)
	assertEqualTuple(t, NewVector(0, 1, 0), n3)
}

func TestIntersectWithARayParallelToThePlane(t *testing.T) {
	plane := NewPlane()
	localPlane := plane.LocalShape.(*Plane)
	r := NewRay(NewPoint(0, 10, 0), NewVector(0, 0, 1))
	xs := localPlane.LocalIntersect(r, &plane)

	// TODO assertEmpty function
	assertEqualInt(t, 0, len(xs))
}

func TestIntersectWithACoplanarRay(t *testing.T) {
	plane := NewPlane()
	localPlane := plane.LocalShape.(*Plane)
	r := NewRay(NewPoint(0, 10, 0), NewVector(0, 0, 1))
	xs := localPlane.LocalIntersect(r, &plane)

	assertEqualInt(t, 0, len(xs))
}

func TestARayIntersectingAPlaneFromAbove(t *testing.T) {
	plane := NewPlane()
	localPlane := plane.LocalShape.(*Plane)
	r := NewRay(NewPoint(0, 1, 0), NewVector(0, -1, 0))
	xs := localPlane.LocalIntersect(r, &plane)

	assertEqualInt(t, 1, len(xs))
	assertEqualFloat64(t, 1, xs[0].Time)
	assertEqualShape(t, plane, xs[0].Object)
}

func TestARayIntersectingAPlaneFromBelow(t *testing.T) {
	plane := NewPlane()
	localPlane := plane.LocalShape.(*Plane)
	r := NewRay(NewPoint(0, -1, 0), NewVector(0, 1, 0))
	xs := localPlane.LocalIntersect(r, &plane)

	assertEqualInt(t, 1, len(xs))
	assertEqualFloat64(t, 1, xs[0].Time)
	assertEqualShape(t, plane, xs[0].Object)
}

func TestAPlaneHasABoundingBox(t *testing.T) {
	p := NewPlane()
	b := p.Bounds()

	assertEqualTuple(t, NewPoint(math.Inf(-1), 0, math.Inf(-1)), b.MinPoint)
	assertEqualTuple(t, NewPoint(math.Inf(1), 0, math.Inf(1)), b.MaxPoint)
}