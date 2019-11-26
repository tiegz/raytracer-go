package raytracer

import (
	"math"
	"testing"
)

func TestNewSphere(t *testing.T) {
	shape := NewSphere()
	sphere := shape.LocalShape.(*Sphere)

	assertEqualTuple(t, NewPoint(0, 0, 0), sphere.Origin)
	assertEqualFloat64(t, 1.0, sphere.Radius)
}

func TestNormalAtOnXAxis(t *testing.T) {
	s1 := NewSphere()
	actual := s1.NormalAt(NewPoint(1, 0, 0), NewIntersection(0, s1))
	expected := NewVector(1, 0, 0)

	assertEqualTuple(t, expected, actual)
}

func TestNormalAtOnYAxis(t *testing.T) {
	s1 := NewSphere()
	actual := s1.NormalAt(NewPoint(0, 1, 0), NewIntersection(0, s1))
	expected := NewVector(0, 1, 0)

	assertEqualTuple(t, expected, actual)
}

func TestNormalAtOnZAxis(t *testing.T) {
	s1 := NewSphere()
	actual := s1.NormalAt(NewPoint(0, 0, 1), NewIntersection(0, s1))
	expected := NewVector(0, 0, 1)

	assertEqualTuple(t, expected, actual)
}

func TestNormalAtOnNonAxialPoint(t *testing.T) {
	s1 := NewSphere()
	actual := s1.NormalAt(NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3), NewIntersection(0, s1))
	expected := NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)

	assertEqualTuple(t, expected, actual)
}

func TestNormalAtIsNormalized(t *testing.T) {
	s1 := NewSphere()
	actual := s1.NormalAt(NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3), NewIntersection(0, s1))
	expected := actual.Normalized()

	assertEqualTuple(t, expected, actual)
}

func TestRayIntersectsSphereAtTwoPoints(t *testing.T) {
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	sphere := NewSphere()
	intersections := sphere.Intersect(ray)

	assertEqualInt(t, 2, len(intersections))
	assertEqualIntersection(t, NewIntersection(4.0, sphere), intersections[0])
	assertEqualIntersection(t, NewIntersection(6.0, sphere), intersections[1])
}

func TestRayIntersectsSphereAtTangent(t *testing.T) {
	ray := NewRay(NewPoint(0, 1, -5), NewVector(0, 0, 1))
	sphere := NewSphere()
	intersections := sphere.Intersect(ray)

	assertEqualInt(t, 2, len(intersections))
	assertEqualIntersection(t, NewIntersection(5.0, sphere), intersections[0])
	assertEqualIntersection(t, NewIntersection(5.0, sphere), intersections[1])
}

func TestRayMissesSphere(t *testing.T) {
	ray := NewRay(NewPoint(0, 2, -5), NewVector(0, 0, 1))
	sphere := NewSphere()
	intersections := sphere.Intersect(ray)

	assertEqualInt(t, 0, len(intersections))
}

func TestRayOriginatesInsideAndIntersectsSphere(t *testing.T) {
	ray := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	sphere := NewSphere()
	intersections := sphere.Intersect(ray)

	assertEqualInt(t, 2, len(intersections))
	assertEqualIntersection(t, NewIntersection(-1.0, sphere), intersections[0])
	assertEqualIntersection(t, NewIntersection(1.0, sphere), intersections[1])
}

func TestRayIsBehindAndIntersectsSphere(t *testing.T) {
	ray := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
	sphere := NewSphere()
	intersections := sphere.Intersect(ray)

	assertEqualInt(t, 2, len(intersections))
	assertEqualIntersection(t, NewIntersection(-6.0, sphere), intersections[0])
	assertEqualIntersection(t, NewIntersection(-4.0, sphere), intersections[1])
}

func TestAHelperForProducingASphereWithAGlassyMaterial(t *testing.T) {
	s := NewGlassSphere()

	assertEqualMatrix(t, IdentityMatrix(), s.Transform)
	assertEqualFloat64(t, 1.0, s.Material.Transparency)
	assertEqualFloat64(t, 1.5, s.Material.RefractiveIndex)

}
