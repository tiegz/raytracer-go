package raytracer

import (
	"fmt"
)

type Ray struct {
	Origin    Tuple
	Direction Tuple // aka velocity (i.e. how far it moves per time unit)
}

func NewRay(o, d Tuple) Ray {
	return Ray{o, d}
}

func NullRay() Ray {
	return NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 0))
}

func (r Ray) String() string {
	return fmt.Sprintf("Ray(\n  Origin: %v\n  Direction: %v\n)", r.Origin, r.Direction)
}

func (r Ray) Position(time float64) Tuple {
	return r.Origin.Add(r.Direction.Multiply(time))
}

func (r Ray) Transform(t Matrix) Ray {
	r2 := Ray{}

	r2.Origin = t.MultiplyByTuple(r.Origin)
	r2.Direction = t.MultiplyByTuple(r.Direction)

	return r2
}

func (r Ray) IsEqualTo(r2 Ray) bool {
	if !r.Origin.IsEqualTo(r2.Origin) {
		return false
	} else if !r.Direction.IsEqualTo(r2.Direction) {
		return false
	}
	return true
}