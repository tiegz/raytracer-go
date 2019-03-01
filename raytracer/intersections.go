package raytracer

import (
	"fmt"
	"math"
)

type Intersection struct {
	Time   float64
	Object Sphere
}

func NullIntersection() Intersection {
	return NewIntersection(math.MaxFloat64, NewSphere())
}

type Intersections []Intersection

// func (is Intersections) Len() int           { return len(is) }
// func (is Intersections) Less(i, j int) bool { is[i], is[j] = is[j], is[i] }
// func (is Intersections) Swap(i, j int) int  { return is[i].Time < is[j].Time }

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
