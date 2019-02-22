package raytracer

import "testing"

func TestNewIntersection(t *testing.T) {
	sphere := NewSphere()
	i := NewIntersection(1.23, sphere)

	assertEqualFloat64(t, 1.23, i.T)
	assertEqualObject(t, sphere, i.Object)
}

func TestAggregatingIntersections(t *testing.T) {
	sphere := NewSphere()
	i1 := NewIntersection(1, sphere)
	i2 := NewIntersection(2, sphere)
	intersections := Intersections{i1, i2}

	assertEqualInt(t, 2, len(intersections))
	assertEqualFloat64(t, 1, intersections[0].T)
	assertEqualFloat64(t, 2, intersections[1].T)
}
