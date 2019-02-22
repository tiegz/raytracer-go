package raytracer

import "testing"

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
