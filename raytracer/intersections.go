package raytracer

import (
	"fmt"
)

type Intersection struct {
	Time   float64
	Object Sphere
}
type Intersections []Intersection

// func (is Intersections) Len() int           { return len(is) }
// func (is Intersections) Less(i, j int) bool { is[i], is[j] = is[j], is[i] }
// func (is Intersections) Swap(i, j int) int  { return is[i].Time < is[j].Time }

func NewIntersection(t float64, obj Sphere) Intersection {
	return Intersection{t, obj}
}

func (i Intersection) String() string {
	return fmt.Sprintf("Intersection( %.3f, %v )", i.Time, i.Object)
}

func (is *Intersections) Hit() (Intersection, error) {
	var minIntersection Intersection // initial value is zero value of Intersection
	hasFoundAPositiveIntersection := false
	for _, intersection := range *is {
		if intersection.Time > 0 {
			if !hasFoundAPositiveIntersection || intersection.Time < minIntersection.Time {
				hasFoundAPositiveIntersection = true
				minIntersection = intersection
			}
		}
	}

	// TODO should this function just return a pointer instead, so we can return nil?
	if minIntersection == (Intersection{}) {
		return minIntersection, fmt.Errorf("Couldn't find intersection.")
	} else {
		return minIntersection, nil
	}
}
