package raytracer

import (
	"math"
)

type Ray struct {
	Origin    Tuple
	Direction Tuple // aka velocity (i.e. how far it moves per time unit)
}

func NewRay(o, d Tuple) Ray {
	return Ray{o, d}
}

func (r *Ray) Position(time float64) Tuple {
	return r.Origin.Add(r.Direction.Multiply(time))
}

func (r *Ray) Intersect(s Sphere) Intersections {
	i := make(Intersections, 0, 2)

	sphereToRay := r.Origin.Subtract(s.Origin)
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant < 0 {
		return i
	}

	i1 := NewIntersection((-b-math.Sqrt(discriminant))/(2*a), s)
	i2 := NewIntersection((-b+math.Sqrt(discriminant))/(2*a), s)

	i = append(i, i1)
	i = append(i, i2)

	return i
}
