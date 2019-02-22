package raytracer

import "testing"

func TestNewSphere(t *testing.T) {
	sphere := NewSphere()

	assertEqualTuple(t, NewPoint(0, 0, 0), sphere.Origin)
	assertEqualFloat64(t, 1.0, sphere.Radius)
}
