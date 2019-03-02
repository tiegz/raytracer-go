package raytracer

import (
	"fmt"
	"math"
)

type Ray struct {
	Origin    Tuple
	Direction Tuple // aka velocity (i.e. how far it moves per time unit)
}

func NewRay(o, d Tuple) Ray {
	return Ray{o, d}
}

func (r Ray) String() string {
	return fmt.Sprintf("Ray( %v, %v)", r.Origin, r.Direction)
}

func (r *Ray) Position(time float64) Tuple {
	return r.Origin.Add(r.Direction.Multiply(time))
}

func (r *Ray) Intersect(s Sphere) Intersections {
	// Instead of applying object's transformation to object, we can just apply
	// the inverse of the transformation to the ray.
	r2 := r.Transform(s.Transform.Inverse())

	i := make(Intersections, 0, 2)

	sphereToRay := r2.Origin.Subtract(s.Origin)
	a := r2.Direction.Dot(r2.Direction)
	b := 2 * r2.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1.0
	discriminant := (b * b) - 4*a*c

	if discriminant < 0 {
		return i
	}

	i1 := NewIntersection((-b-math.Sqrt(discriminant))/(2*a), s)
	i2 := NewIntersection((-b+math.Sqrt(discriminant))/(2*a), s)

	i = append(i, i1)
	i = append(i, i2)

	return i
}

func (r *Ray) Transform(t Matrix) Ray {
	r2 := Ray{}

	r2.Origin = t.MultiplyByTuple(r.Origin)
	r2.Direction = t.MultiplyByTuple(r.Direction)

	return r2
}
