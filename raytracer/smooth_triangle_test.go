package raytracer

import (
	"testing"
)

// TODO: can we somehow cajole the first test into using this too?
func getTestTriangle() *Shape {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	n1 := NewVector(0, 1, 0)
	n2 := NewVector(-1, 0, 0)
	n3 := NewVector(1, 0, 0)
	shape := NewSmoothTriangle(p1, p2, p3, n1, n2, n3)

	return shape
}

func TestConstructingASmoothTriangle(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	n1 := NewVector(0, 1, 0)
	n2 := NewVector(-1, 0, 0)
	n3 := NewVector(1, 0, 0)

	shape := NewSmoothTriangle(p1, p2, p3, n1, n2, n3)
	tri := shape.LocalShape.(*SmoothTriangle)

	assertEqualTuple(t, p1, tri.P1)
	assertEqualTuple(t, p2, tri.P2)
	assertEqualTuple(t, p3, tri.P3)
	assertEqualTuple(t, n1, tri.N1)
	assertEqualTuple(t, n2, tri.N2)
	assertEqualTuple(t, n3, tri.N3)
}

func TestAnIntersectionWithASmoothTriangleStoresUV(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	n1 := NewVector(0, 1, 0)
	n2 := NewVector(-1, 0, 0)
	n3 := NewVector(1, 0, 0)

	shape := NewSmoothTriangle(p1, p2, p3, n1, n2, n3)
	tri := shape.LocalShape.(*SmoothTriangle)

	r := NewRay(NewPoint(-0.2, 0.3, -2), NewVector(0, 0, 1))
	xs := tri.LocalIntersect(r, shape)

	assertEqualFloat64(t, 0.45, xs[0].U)
	assertEqualFloat64(t, 0.25, xs[0].V)
}

func TestASmoothTriangleUsesUVToInteprolateTheNormal(t *testing.T) {
	shape := getTestTriangle()

	i := NewIntersectionWithUV(1, shape, 0.45, 0.25)
	n := shape.NormalAt(NewPoint(0, 0, 0), i)

	assertEqualTuple(t, NewVector(-0.5547, 0.83205, 0), n)
}

func TestPreparingTheNormalOnASmoothTriangle(t *testing.T) {
	shape := getTestTriangle()

	i := NewIntersectionWithUV(1, shape, 0.45, 0.25)
	r := NewRay(NewPoint(-0.2, 0.3, -2), NewVector(0, 0, 1))
	xs := Intersections{i}
	c := i.PrepareComputations(r, xs...)

	assertEqualTuple(t, NewVector(-0.5547, 0.83205, 0), c.NormalV)
}
