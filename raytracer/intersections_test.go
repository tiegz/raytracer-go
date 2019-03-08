package raytracer

import "testing"

func TestNewIntersection(t *testing.T) {
	sphere := NewSphere()
	i := NewIntersection(1.23, sphere)

	assertEqualFloat64(t, 1.23, i.Time)
	assertEqualObject(t, sphere, i.Object)
}

func TestAggregatingIntersections(t *testing.T) {
	sphere := NewSphere()
	i1 := NewIntersection(1, sphere)
	i2 := NewIntersection(2, sphere)
	intersections := Intersections{i1, i2}

	assertEqualInt(t, 2, len(intersections))
	assertEqualFloat64(t, 1, intersections[0].Time)
	assertEqualFloat64(t, 2, intersections[1].Time)
}

func TestHitWithAllIntersectionsHavingPositiveT(t *testing.T) {
	sphere := NewSphere()
	i1 := NewIntersection(1, sphere)
	i2 := NewIntersection(2, sphere)
	intersections := Intersections{i1, i2}

	hit := intersections.Hit()
	assertEqualIntersection(t, i1, hit)
}

func TestHitWithSomeIntersectionsHavingNegativeT(t *testing.T) {
	sphere := NewSphere()
	i1 := NewIntersection(-1, sphere)
	i2 := NewIntersection(1, sphere)
	intersections := Intersections{i1, i2}

	hit := intersections.Hit()
	assertEqualIntersection(t, i2, hit)
}

func TestHitWhenAllIntersectionsHaveNegativeT(t *testing.T) {
	sphere := NewSphere()
	i1 := NewIntersection(-2, sphere)
	i2 := NewIntersection(-1, sphere)
	intersections := Intersections{i1, i2}

	hit := intersections.Hit()
	assertEqualIntersection(t, NullIntersection(), hit)
}

func TestHitIsAlwaysLowestNonNegativeIntersection(t *testing.T) {
	sphere := NewSphere()
	i1 := NewIntersection(5, sphere)
	i2 := NewIntersection(7, sphere)
	i3 := NewIntersection(-3, sphere)
	i4 := NewIntersection(2, sphere)
	intersections := Intersections{i1, i2, i3, i4}

	hit := intersections.Hit()
	assertEqualIntersection(t, i4, hit)
}

func TestPrecomputingStateOfIntersection(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewSphere()
	i := NewIntersection(4, s)
	c := i.PrepareComputations(r)

	assertEqualFloat64(t, i.Time, c.Time)
	assertEqualObject(t, s, c.Object)
	assertEqualTuple(t, NewPoint(0, 0, -1), c.Point)
	assertEqualTuple(t, NewVector(0, 0, -1), c.EyeV)
	assertEqualTuple(t, NewVector(0, 0, -1), c.NormalV)
}

func TestHitWhenIntersectionOccursOnOutside(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewSphere()
	i := NewIntersection(4, s)
	c := i.PrepareComputations(r)

	assert(t, !c.Inside)
}

func TestHitWhenIntersectionOccursOnInside(t *testing.T) {
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	s := NewSphere()
	i := NewIntersection(1, s)
	c := i.PrepareComputations(r)

	assertEqualObject(t, s, c.Object)
	assertEqualTuple(t, NewPoint(0, 0, 1), c.Point)
	assertEqualTuple(t, NewVector(0, 0, -1), c.EyeV)
	assertEqualTuple(t, NewVector(0, 0, -1), c.NormalV)
	assert(t, c.Inside)
}
