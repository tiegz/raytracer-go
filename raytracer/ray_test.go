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

func TestTranslatingRay(t *testing.T) {
	r := NewRay(NewPoint(1, 2, 3), NewVector(0, 1, 0))
	translation := NewTranslation(3, 4, 5)
	r2 := r.Transform(translation)

	assertEqualTuple(t, NewPoint(4, 6, 8), r2.Origin)
	assertEqualTuple(t, NewVector(0, 1, 0), r2.Direction)
}

func TestScalingRay(t *testing.T) {
	r := NewRay(NewPoint(1, 2, 3), NewVector(0, 1, 0))
	translation := NewScale(2, 3, 4)
	r2 := r.Transform(translation)

	assertEqualTuple(t, NewPoint(2, 6, 12), r2.Origin)
	assertEqualTuple(t, NewVector(0, 3, 0), r2.Direction)
}

/////////////
// Benchmarks
/////////////

func BenchmarkRayMethodIsEqualTo(b *testing.B) {
	ray := NewRay(NewPoint(2, 3, 4), NewVector(1, 0, 0))
	for i := 0; i < b.N; i++ {
		ray.IsEqualTo(ray)
	}
}

func BenchmarkRayMethodPosition(b *testing.B) {
	ray := NewRay(NewPoint(2, 3, 4), NewVector(1, 0, 0))
	for i := 0; i < b.N; i++ {
		ray.Position(2.5)
	}
}

func BenchmarkRayMethodTransform(b *testing.B) {
	r := NewRay(NewPoint(1, 2, 3), NewVector(0, 1, 0))
	translation := NewTranslation(3, 4, 5)
	for i := 0; i < b.N; i++ {
		r.Transform(translation)
	}
}
