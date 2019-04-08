package raytracer

import (
	"math"
	"testing"
)

func TestDefaultTransformation(t *testing.T) {
	s := NewTestShape()

	assertEqualMatrix(t, s.Transform, IdentityMatrix())
}

func TestAssigningATransformation(t *testing.T) {
	s := NewTestShape()
	s.Transform = NewTranslation(2, 3, 4)

	assertEqualMatrix(t, NewTranslation(2, 3, 4), s.Transform)
}

// These replace the tests named “A sphere has a default material” and “A sphere may be assigned a material” (from the sphere scenarios on page 85).
func TestDefaultMatrial(t *testing.T) {
	s := NewTestShape()

	assertEqualMaterial(t, s.Material, DefaultMaterial())
}

func TestAssigningAMaterial(t *testing.T) {
	s := NewTestShape()
	m := DefaultMaterial()
	m.Ambient = 1
	s.Material = m

	assertEqualMaterial(t, s.Material, m)
}

//  These tests are both based on (and replace) the tests called “Intersecting a scaled sphere with a ray” and “Intersecting a translated sphere with a ray”
func TestIntersectingAScaledShapeWithARay(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewTestShape()
	s.Transform = NewScale(2, 2, 2)
	s.Intersect(r)

	assertEqualTuple(t, NewPoint(0, 0, -2.5), s.SavedRay.Origin)
	assertEqualTuple(t, NewVector(0, 0, 0.5), s.SavedRay.Direction)
}

func TestIntersectingATranslatedShapeWithARay(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewTestShape()
	s.Transform = NewTranslation(5, 0, 0)
	s.Intersect(r)

	assertEqualTuple(t, NewPoint(-5, 0, -5), s.SavedRay.Origin)
	assertEqualTuple(t, NewVector(0, 0, 1), s.SavedRay.Direction)
}

// The following two tests replace the ones called “Computing the normal on a translated sphere” and “Computing the normal on a transformed sphere”
func TestComputingTheNormalOnATranslatedShape(t *testing.T) {
	s := NewTestShape()
	s.Transform = NewTranslation(0, 1, 0)
	n := s.NormalAt(NewPoint(0, 1.70711, -0.70711))

	assertEqualTuple(t, NewVector(0, 0.70711, -0.70711), n)
}

func TestComputingTheNormalOnATransformedShape(t *testing.T) {
	s := NewTestShape()
	transform := NewScale(1, 0.5, 1)
	transform = transform.Multiply(NewRotateZ(math.Pi / 5))
	s.Transform = transform
	n := s.NormalAt(NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2))

	assertEqualTuple(t, NewVector(0, 0.97014, -0.24254), n)
}