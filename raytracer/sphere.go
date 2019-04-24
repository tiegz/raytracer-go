package raytracer

import (
	"fmt"
	"math"
)

type Sphere struct {
	Origin Tuple
	Radius float64
}

func NewSphere() Shape {
	return NewShape(&Sphere{NewPoint(0, 0, 0), 1})
}

func (s Sphere) String() string {
	return fmt.Sprintf("Sphere( %v, %.1f )", s.Origin, s.Radius)
}

func (s Sphere) localString() string {
	return s.String()
}

/////////////////////////
// ShapeInterface methods
/////////////////////////

// TODO can we remove Shape arg somehow? It's only there because ShapeInterface
// has no knowledge of its parent, but we need to put its aprent in the Intersection :(
func (s Sphere) LocalIntersect(localRay Ray, shape *Shape) Intersections {
	i := make(Intersections, 0, 2)

	sphereToRay := localRay.Origin.Subtract(s.Origin)
	a := localRay.Direction.Dot(localRay.Direction)
	b := 2 * localRay.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1.0
	discriminant := (b * b) - 4*a*c

	if discriminant < 0 {
		return i
	}

	i1 := NewIntersection((-b-math.Sqrt(discriminant))/(2*a), *shape)
	i2 := NewIntersection((-b+math.Sqrt(discriminant))/(2*a), *shape)

	i = append(i, i1, i2)

	return i
}

func (s Sphere) LocalNormalAt(localPoint Tuple) Tuple {
	return localPoint.Subtract(s.Origin)
}

func (s Sphere) localIsEqualTo(s2 ShapeInterface) bool {
	// NB I still don't know why we need to do a type assertiong to *Sphere
	// instead of Sphere, but it fixes this panic when comparing two
	// spheres: 'panic: interface conversion: raytracer.ShapeInterface is *raytracer.Sphere, not raytracer.Sphere [recovered]
	s2Sphere := s2.(*Sphere)
	if !s.Origin.IsEqualTo(s2Sphere.Origin) {
		return false
	} else if s.Radius != s2Sphere.Radius {
		return false
	}
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (s Sphere) localType() string {
	return "Sphere"
}
