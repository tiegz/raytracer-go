package raytracer

import (
	"fmt"
	"math"
)

type Intersection struct {
	Time   float64
	Object Shape
}

type Computation struct {
	Time      float64
	Object    Shape
	Point     Tuple
	OverPoint Tuple
	EyeV      Tuple
	NormalV   Tuple
	Inside    bool
}

func NullIntersection() Intersection {
	return Intersection{math.MaxFloat64, NewNullShape()}
}

func (i Intersection) IsNull() bool {
	return i.Time == math.MaxFloat64
}

type Intersections []Intersection

func NewIntersection(t float64, obj Shape) Intersection {
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

func (i Intersection) String() string {
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
	c.OverPoint = c.Point.Add(c.NormalV.Multiply(EPSILON)) // to avoid "raytracer acne" with shadows
	if c.NormalV.Dot(c.EyeV) < 0 {
		c.Inside = true
		c.NormalV = c.NormalV.Negate()
	} else {
		c.Inside = false
	}

	return c
}
