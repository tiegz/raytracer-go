package raytracer

import (
	"fmt"
	"math"
)

type Sphere struct {
	Origin Tuple
	Radius float64
}

func NewSphere() *Shape {
	return NewShape(&Sphere{NewPoint(0, 0, 0), 1})
}

func NewGlassSphere() *Shape {
	shape := NewShape(&Sphere{NewPoint(0, 0, 0), 1})
	shape.Material.Transparency = 1.0
	shape.Material.RefractiveIndex = 1.5
	return shape
}

func (s Sphere) String() string {
	return fmt.Sprintf("Sphere(\n  Origin: %v\n  Radius: %.1f\n)", s.Origin, s.Radius)
}

/////////////////////////
// ShapeInterface methods
/////////////////////////

func (s Sphere) LocalBounds() BoundingBox {
	return NewBoundingBox(NewPoint(-1, -1, -1), NewPoint(1, 1, 1))
}

func (s Sphere) localString() string {
	return s.String()
}

// TODO can we remove Shape arg somehow? It's only there because ShapeInterface
// has no knowledge of its parent, but we need to put its aprent in the Intersection :(
func (s Sphere) LocalIntersect(r Ray, shape *Shape) Intersections {
	i := make(Intersections, 0, 2)

	sphereToRay := r.Origin.Subtract(s.Origin)
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1.0
	discriminant := (b * b) - 4*a*c

	if discriminant < 0 {
		return i
	}

	i1 := NewIntersection((-b-math.Sqrt(discriminant))/(2*a), shape)
	i2 := NewIntersection((-b+math.Sqrt(discriminant))/(2*a), shape)

	i = append(i, i1, i2)

	return i
}

func (s Sphere) LocalNormalAt(localPoint Tuple, hit Intersection) Tuple {
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
