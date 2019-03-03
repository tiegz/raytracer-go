package raytracer

import (
	"math"
	"testing"
)

func TestNewSphere(t *testing.T) {
	sphere := NewSphere()

	assertEqualTuple(t, NewPoint(0, 0, 0), sphere.Origin)
	assertEqualFloat64(t, 1.0, sphere.Radius)
	assertEqualMatrix(t, IdentityMatrix(), sphere.Transform)
}

func TestChangingSpheresTransformation(t *testing.T) {
	s1 := NewSphere()
	t1 := NewTranslation(2, 3, 4)
	s1.Transform = t1

	assertEqualMatrix(t, t1, s1.Transform)
}

func TestIntersectingScaledSphereWithRay(t *testing.T) {
	r1 := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s1 := NewSphere()
	s1.Transform = NewScale(2, 2, 2)
	intersections := r1.Intersect(s1)

	assertEqualInt(t, 2, len(intersections))
	assertEqualFloat64(t, 3, intersections[0].Time)
	assertEqualFloat64(t, 7, intersections[1].Time)
}

func TestIntersectingTranslatedSphereWithRay(t *testing.T) {
	r1 := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s1 := NewSphere()
	s1.Transform = NewTranslation(5, 0, 0)
	intersections := r1.Intersect(s1)

	assertEqualInt(t, 0, len(intersections))
}

func TestNormalAtOnXAxis(t *testing.T) {
	s1 := NewSphere()
	actual := s1.NormalAt(NewPoint(1, 0, 0))
	expected := NewVector(1, 0, 0)

	assertEqualTuple(t, expected, actual)
}

func TestNormalAtOnYAxis(t *testing.T) {
	s1 := NewSphere()
	actual := s1.NormalAt(NewPoint(0, 1, 0))
	expected := NewVector(0, 1, 0)

	assertEqualTuple(t, expected, actual)
}

func TestNormalAtOnZAxis(t *testing.T) {
	s1 := NewSphere()
	actual := s1.NormalAt(NewPoint(0, 0, 1))
	expected := NewVector(0, 0, 1)

	assertEqualTuple(t, expected, actual)
}

func TestNormalAtOnNonAxialPoint(t *testing.T) {
	s1 := NewSphere()
	actual := s1.NormalAt(NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	expected := NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)

	assertEqualTuple(t, expected, actual)
}

func TestNormalAtIsNormalized(t *testing.T) {
	s1 := NewSphere()
	actual := s1.NormalAt(NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	expected := actual.Normalized()

	assertEqualTuple(t, expected, actual)
}

func TestNormalAtOnTranslatedSphere(t *testing.T) {
	s1 := NewSphere()
	s1.Transform = NewTranslation(0, 1, 0)
	actual := s1.NormalAt(NewPoint(0, 1.70711, -0.70711))
	expected := NewVector(0, 0.70711, -0.70711)

	assertEqualTuple(t, expected, actual)
}

func TestNormalAtOnTransformedSphere(t *testing.T) {
	s1 := NewSphere()
	scale := NewScale(1, 0.5, 1)
	rotation := NewRotateZ(math.Pi / 5)
	transform := scale.Multiply(rotation)
	s1.Transform = s1.Transform.Multiply(transform)
	actual := s1.NormalAt(NewPoint(0, math.Sqrt(2)/2, -(math.Sqrt(2) / 2)))
	expected := NewVector(0, 0.97014, -0.24254)

	assertEqualTuple(t, expected, actual)
}
