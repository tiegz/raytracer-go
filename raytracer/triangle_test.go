package raytracer

import (
	"testing"
)

func TestConstructingATriangle(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	shape := NewTriangle(p1, p2, p3)
	tri := shape.LocalShape.(*Triangle)

	assertEqualTuple(t, p1, tri.P1)
	assertEqualTuple(t, p2, tri.P2)
	assertEqualTuple(t, p3, tri.P3)
	assertEqualTuple(t, NewVector(-1, -1, 0), tri.E1)
	assertEqualTuple(t, NewVector(1, -1, 0), tri.E2)
	assertEqualTuple(t, NewVector(0, 0, -1), tri.Normal)
}

func TestFindingTheNormalOnATriangle(t *testing.T) {
	shape := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))
	tri := shape.LocalShape.(*Triangle)
	n1 := tri.LocalNormalAt(NewPoint(0, 0.5, 0))
	n2 := tri.LocalNormalAt(NewPoint(-0.5, 0.75, 0))
	n3 := tri.LocalNormalAt(NewPoint(0.5, 0.25, 0))

	assertEqualTuple(t, tri.Normal, n1)
	assertEqualTuple(t, tri.Normal, n2)
	assertEqualTuple(t, tri.Normal, n3)
}

func TestIntersectingARayParallelToTheTriangle(t *testing.T) {
	shape := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))
	tri := shape.LocalShape.(*Triangle)
	r := NewRay(NewPoint(0, -1, -2), NewVector(0, 1, 0))
	xs := tri.LocalIntersect(r, &shape)

	assertEqualInt(t, 0, len(xs))
}

func TestARayMissesTheP1ToP3Edge(t *testing.T) {
	shape := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))
	tri := shape.LocalShape.(*Triangle)
	r := NewRay(NewPoint(1, 1, -2), NewVector(0, 0, 1))
	xs := tri.LocalIntersect(r, &shape)

	assertEqualInt(t, 0, len(xs))
}

func TestARayMissesThteP1ToP2Edge(t *testing.T) {
	shape := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))
	tri := shape.LocalShape.(*Triangle)
	r := NewRay(NewPoint(-1, 1, -2), NewVector(0, 0, 1))
	xs := tri.LocalIntersect(r, &shape)

	assertEqualInt(t, 0, len(xs))
}

func TestARayMissesThteP2ToP3Edge(t *testing.T) {
	shape := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))
	tri := shape.LocalShape.(*Triangle)
	r := NewRay(NewPoint(0, -1, -2), NewVector(0, 0, 1))
	xs := tri.LocalIntersect(r, &shape)

	assertEqualInt(t, 0, len(xs))
}

func TestARayStrikesATriangle(t *testing.T) {
	shape := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))
	tri := shape.LocalShape.(*Triangle)
	r := NewRay(NewPoint(0, 0.5, -2), NewVector(0, 0, 1))
	xs := tri.LocalIntersect(r, &shape)

	assertEqualInt(t, 1, len(xs))
	assertEqualFloat64(t, 2, xs[0].Time)
}
