package raytracer

import "testing"

func TestCreatingAndQueryingARay(t *testing.T) {
	origin := NewPoint(1, 2, 3)
	direction := NewVector(4, 5, 6)
	ray := NewRay(origin, direction)

	assertEqualTuple(t, origin, ray.Origin)
	assertEqualTuple(t, direction, ray.Direction)
}

func TestComputingPointFromDistance(t *testing.T) {
	ray := NewRay(NewPoint(2, 3, 4), NewVector(1, 0, 0))

	assertEqualTuple(t, NewPoint(2, 3, 4), ray.Position(0))
	assertEqualTuple(t, NewPoint(3, 3, 4), ray.Position(1))
	assertEqualTuple(t, NewPoint(1, 3, 4), ray.Position(-1))
	assertEqualTuple(t, NewPoint(4.5, 3, 4), ray.Position(2.5))
}

func TestRayIntersectsSphereAtTwoPoints(t *testing.T) {
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	sphere := NewSphere()
	intersections := ray.Intersect(sphere)

	assertEqualInt(t, 2, len(intersections))
	assertEqualIntersection(t, NewIntersection(4.0, sphere), intersections[0])
	assertEqualIntersection(t, NewIntersection(6.0, sphere), intersections[1])
}

func TestRayIntersectsSphereAtTangent(t *testing.T) {
	ray := NewRay(NewPoint(0, 1, -5), NewVector(0, 0, 1))
	sphere := NewSphere()
	intersections := ray.Intersect(sphere)

	assertEqualInt(t, 2, len(intersections))
	assertEqualIntersection(t, NewIntersection(5.0, sphere), intersections[0])
	assertEqualIntersection(t, NewIntersection(5.0, sphere), intersections[1])
}

func TestRayMissesSphere(t *testing.T) {
	ray := NewRay(NewPoint(0, 2, -5), NewVector(0, 0, 1))
	sphere := NewSphere()
	intersections := ray.Intersect(sphere)

	assertEqualInt(t, 0, len(intersections))
}

func TestRayOriginatesInsideAndIntersectsSphere(t *testing.T) {
	ray := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	sphere := NewSphere()
	intersections := ray.Intersect(sphere)

	assertEqualInt(t, 2, len(intersections))
	assertEqualIntersection(t, NewIntersection(-1.0, sphere), intersections[0])
	assertEqualIntersection(t, NewIntersection(1.0, sphere), intersections[1])
}

func TestRayIsBehindAndIntersectsSphere(t *testing.T) {
	ray := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
	sphere := NewSphere()
	intersections := ray.Intersect(sphere)

	assertEqualInt(t, 2, len(intersections))
	assertEqualIntersection(t, NewIntersection(-6.0, sphere), intersections[0])
	assertEqualIntersection(t, NewIntersection(-4.0, sphere), intersections[1])
}
