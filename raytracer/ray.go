package raytracer

import (
	"fmt"
)

type Ray struct {
	Origin    Tuple
	Direction Tuple // aka velocity (i.e. how far it moves per time unit)
}

func NewRay(o, d Tuple) *Ray {
	return &Ray{o, d}
}

func (r *Ray) String() string {
	return fmt.Sprintf("Ray(\n  Origin: %v\n  Direction: %v\n)", r.Origin, r.Direction)
}

func (r *Ray) Position(time float64) Tuple {
	return r.Origin.Add(r.Direction.Multiply(time))
}

func (r *Ray) Transform(t Matrix) *Ray {
	return NewRay(
		t.MultiplyByTuple(r.Origin),
		t.MultiplyByTuple(r.Direction),
	)
}

func (r *Ray) IsEqualTo(r2 *Ray) bool {
	if !r.Origin.IsEqualTo(r2.Origin) {
		return false
	} else if !r.Direction.IsEqualTo(r2.Direction) {
		return false
	}
	return true
}
