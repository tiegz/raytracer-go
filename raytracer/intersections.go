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
	Time      float64 // the moment (in time units) at which the intersection happened
	Object    Shape   // the object that was intersected
	Point     Tuple   // the point where intersection happened
	OverPoint Tuple   // the Point value adjusted slightly to avoid "raytracer acne"
	EyeV      Tuple   // the vector from eye to the Point
	NormalV   Tuple   // the normal on the object at the given Point
	ReflectV  Tuple   // the vector of reflection
	Inside    bool    // was the intersection inside the object?
	N1        float64 // refractive index of object being exited at ray-object intersection
	N2        float64 // refractive index of object being entered at ray-object intersection
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

// r:  the ray that hit the intersection
// xs: "... the collection of all intersections, which can tell you where the hit is relative to the rest of the intersections ...""
func (i *Intersection) PrepareComputations(r Ray, xs ...Intersection) Computation {
	if len(xs) == 0 {
		xs = []Intersection{*i}
	}

	c := Computation{}
	c.Time = i.Time
	c.Object = i.Object
	c.Point = r.Position(c.Time)
	c.EyeV = r.Direction.Negate()
	c.NormalV = c.Object.NormalAt(c.Point)
	c.ReflectV = r.Direction.Reflect(c.NormalV)            // TODO after negating the normal, if necessary
	c.OverPoint = c.Point.Add(c.NormalV.Multiply(EPSILON)) // to avoid "raytracer acne" with shadows
	if c.NormalV.Dot(c.EyeV) < 0 {
		c.Inside = true
		c.NormalV = c.NormalV.Negate()
	} else {
		c.Inside = false
	}

	visitedShapes := []Shape{}
	for _, intersection := range xs {
		isHit := intersection.IsEqualTo(*i) // is this intersection the hit?

		if isHit {
			if len(visitedShapes) == 0 {
				c.N1 = 1.0
			} else {
				c.N1 = visitedShapes[len(visitedShapes)-1].Material.RefractiveIndex
			}
		}

		// TODO move somewhere else?
		indexOf := func(shape Shape, shapes []Shape) int {
			for idx, s := range shapes {
				if s.IsEqualTo(shape) {
					return idx
				}
			}
			return -1
		}

		if containerIdx := indexOf(intersection.Object, visitedShapes); containerIdx == -1 { // enter
			visitedShapes = append(visitedShapes, intersection.Object)
		} else { // exit
			visitedShapes = append(visitedShapes[:containerIdx], visitedShapes[containerIdx+1:]...)
		}

		if isHit {
			if len(visitedShapes) == 0 {
				c.N2 = 1.0
			} else {
				c.N2 = visitedShapes[len(visitedShapes)-1].Material.RefractiveIndex
			}
			break
		}
	}

	return c
}
