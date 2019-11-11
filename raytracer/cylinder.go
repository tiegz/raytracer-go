package raytracer

import (
	"fmt"
	"math"
)

type Cylinder struct {
	Origin  Tuple
	Minimum float64
	Maximum float64
	Closed  bool
}

func NewCylinder() Shape {
	return NewShape(&Cylinder{NewPoint(0, 0, 0), math.Inf(-1), math.Inf(1), false})
}

func (cyl Cylinder) String() string {
	return fmt.Sprintf("Cylinder( %v )", cyl.Origin)
}

// Is the ray at time t within a radius of 1 in y axis
// TODO: better name, e.g. isInsideCap
func checkCap(r Ray, t float64, radius float64) bool {
	x := r.Origin.X + t*r.Direction.X
	z := r.Origin.Z + t*r.Direction.Z
	return ((x * x) + (z * z)) <= (radius * radius)
}

/////////////////////////
// ShapeInterface methods
/////////////////////////

func (c Cylinder) LocalBounds() BoundingBox {
	return NewBoundingBox(NewPoint(-1, c.Minimum, -1), NewPoint(1, c.Maximum, 1))
}

func (c Cylinder) localString() string {
	return c.String()
}

// TODO can we remove Shape arg somehow? It's only there because ShapeInterface
// has no knowledge of its parent, but we need to put its aprent in the Intersection :(
// We treat a cube like 6 planes, with 2 parallel planes per axis.
func (cyl Cylinder) LocalIntersect(r Ray, shape *Shape) Intersections {
	a := (r.Direction.X * r.Direction.X) + (r.Direction.Z * r.Direction.Z)
	xs := make(Intersections, 0, 2)

	// Ray is parallel to y axis
	if !equalFloat64s(a, 0.0) {
		b := (2 * r.Origin.X * r.Direction.X) + (2 * r.Origin.Z * r.Direction.Z)
		c := (r.Origin.X * r.Origin.X) + (r.Origin.Z * r.Origin.Z) - 1
		disc := (b * b) - 4*a*c

		// Ray does not intersect cylinder
		if disc < 0 {
			return Intersections{}
		}

		t0 := (-b - math.Sqrt(disc)) / (2 * a)
		t1 := (-b + math.Sqrt(disc)) / (2 * a)

		y0 := r.Origin.Y + (t0 * r.Direction.Y)
		if cyl.Minimum < y0 && y0 < cyl.Maximum {
			xs = append(xs, NewIntersection(t0, *shape))
		}

		y1 := r.Origin.Y + (t1 * r.Direction.Y)
		if cyl.Minimum < y1 && y1 < cyl.Maximum {
			xs = append(xs, NewIntersection(t1, *shape))
		}
	}

	// (This is the logic for intersectCaps() in the book)
	// Caps only matter if the cylinder is closed, and might possibly be # intersected by the ray.
	// if cyl is not closed or ray.direction.y is close to zero
	if cyl.Closed && !equalFloat64s(r.Direction.Y, 0) {
		// Does the ray intersect bottom cap?
		t := (cyl.Minimum - r.Origin.Y) / r.Direction.Y
		if checkCap(r, t, 1) {
			xs = append(xs, NewIntersection(t, *shape))
		}

		// Does the ray intersect top cap?
		t = (cyl.Maximum - r.Origin.Y) / r.Direction.Y
		if checkCap(r, t, 1) {
			xs = append(xs, NewIntersection(t, *shape))
		}
	}

	return xs
}

func (cyl Cylinder) LocalNormalAt(localPoint Tuple, hit Intersection) Tuple {
	// ... compute the square of the distance from the y axis ...
	distance := math.Pow(localPoint.X, 2) + math.Pow(localPoint.Z, 2)
	if distance < 1 && localPoint.Y >= cyl.Maximum-EPSILON {
		return NewVector(0, 1, 0)
	} else if distance < 1 && localPoint.Y <= cyl.Minimum+EPSILON {
		return NewVector(0, -1, 0)
	} else {
		return NewVector(localPoint.X, 0, localPoint.Z)
	}
}

func (c Cylinder) localIsEqualTo(c2 ShapeInterface) bool {
	// NB I still don't know why we need to do a type assertion to *Cylinder
	// instead of Cylinder, but it fixes this panic when comparing two
	// cylinders: 'panic: interface conversion: raytracer.ShapeInterface is *raytracer.Cylinder, not raytracer.Cylinder [recovered]
	c2Cylinder := c2.(*Cylinder)
	if !c.Origin.IsEqualTo(c2Cylinder.Origin) {
		return false
	}
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (c Cylinder) localType() string {
	return "Cylinder"
}
