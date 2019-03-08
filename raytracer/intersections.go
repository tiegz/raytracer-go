package raytracer

import (
	"fmt"
	"math"
)

type Intersection struct {
	Time   float64
	Object Sphere
}

type Computation struct {
	Time    float64
	Object  Sphere
	Point   Tuple
	EyeV    Tuple
	NormalV Tuple
	Inside  bool
}

func NullIntersection() Intersection {
	return NewIntersection(math.MaxFloat64, NewSphere())
}

type Intersections []Intersection

func NewIntersection(t float64, obj Sphere) Intersection {
	return Intersection{t, obj}
}

func (i Intersection) IsEqualTo(i2 Intersection) bool {
	if i.Time != i2.Time {
		return false
	} else if !i.Object.IsEqualTo(i2.Object) {
		return false
	}
	return true
}

func (i *Intersection) String() string {
	return fmt.Sprintf("Intersection( %.3f, %v )", i.Time, i.Object)
}

func (is *Intersections) Hit() Intersection {
	minIntersection := NullIntersection()
	for _, intersection := range *is {
		if intersection.Time > 0 {
			if minIntersection.IsEqualTo(intersection) || intersection.Time < minIntersection.Time {
				minIntersection = intersection
			}
		}
	}

	return minIntersection
}

func (i *Intersection) PrepareComputations(r Ray) Computation {
	c := Computation{}
	c.Time = i.Time
	c.Object = i.Object
	c.Point = r.Position(c.Time)
	c.EyeV = r.Direction.Negate()
	c.NormalV = c.Object.NormalAt(c.Point)
	if c.NormalV.Dot(c.EyeV) < 0 {
		c.Inside = true
		c.NormalV = c.NormalV.Negate()
	} else {
		c.Inside = false
	}

	return c
}
