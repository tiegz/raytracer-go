package raytracer

import (
	"fmt"
	"math"
)

// Plane is implemented on xz axes.
type Plane struct {
}

func NewPlane() Shape {
	return NewShape(&Plane{})
}

func (p Plane) String() string {
	return fmt.Sprintf("Plane( )")
}

func (p Plane) localString() string {
	return p.String()
}

/////////////////////////
// ShapeInterface methods
/////////////////////////

// TODO can we remove Shape arg somehow? It's only there because ShapeInterface
// has no knowledge of its parent, but we need to put its aprent in the Intersection :(
func (p Plane) LocalIntersect(r Ray, shape *Shape) Intersections {
	i := make(Intersections, 0, 1)

	// TODO worth short-circuiting this to return a slice with 0,0 instead of 0,1?
	// No intersections for planes that are parallel or coplanar to plane.
	if math.Abs(r.Direction.Y) < EPSILON {
		return i
	}

	i1 := NewIntersection(-r.Origin.Y/r.Direction.Y, *shape)

	i = append(i, i1)

	return i
}

// ... every single point on the plane has the same normal: vector(0, 1, 0). ...
func (p Plane) LocalNormalAt(localPoint Tuple) Tuple {
	return NewVector(0, 1, 0)
}

func (p Plane) localIsEqualTo(s2 ShapeInterface) bool {
	// Other than the transformation, all planes are equal.
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (p Plane) localType() string {
	return "Plane"
}
