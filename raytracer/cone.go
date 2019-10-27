package raytracer

import (
	"fmt"
	"math"
)

type Cone struct {
	Origin  Tuple
	Minimum float64
	Maximum float64
	Closed  bool
}

func NewCone() Shape {
	return NewShape(&Cone{NewPoint(0, 0, 0), math.Inf(-1), math.Inf(1), false})
}

func (cone Cone) String() string {
	return fmt.Sprintf("Cone( %v )", cone.Origin)
}

func (cone Cone) intersectCaps(xs Intersections, r Ray, shape *Shape) Intersections {
	if !cone.Closed || math.Abs(r.Direction.Y) < EPSILON {
		return xs
	}

	t := (cone.Minimum - r.Origin.Y) / r.Direction.Y
	if checkCap(r, t, math.Abs(cone.Minimum)) {
		xs = append(xs, NewIntersection(t, *shape))
	}

	// Does the ray intersect top cap?
	t = (cone.Maximum - r.Origin.Y) / r.Direction.Y
	if checkCap(r, t, math.Abs(cone.Maximum)) {
		xs = append(xs, NewIntersection(t, *shape))
	}

	return xs
}

/////////////////////////
// ShapeInterface methods
/////////////////////////

func (cone Cone) localString() string {
	return cone.String()
}

// TODO can we remove Shape arg somehow? It's only there because ShapeInterface
// has no knowledge of its parent, but we need to put its aprent in the Intersection :(
// We treat a cube like 6 planes, with 2 parallel planes per axis.
func (cone Cone) LocalIntersect(r Ray, shape *Shape) Intersections {
	o := r.Origin
	d := r.Direction
	a := math.Pow(d.X, 2) - math.Pow(d.Y, 2) + math.Pow(d.Z, 2)
	b := (2 * o.X * d.X) - (2 * o.Y * d.Y) + (2 * o.Z * d.Z)
	c := math.Pow(o.X, 2) - math.Pow(o.Y, 2) + math.Pow(o.Z, 2)

	// Ray is not parallel to y axis
	if math.Abs(a) <= EPSILON {
		if math.Abs(b) <= EPSILON {
			return Intersections{}
		}

		t := -c / (2 * b)
		xs := Intersections{NewIntersection(t, *shape)}

		return cone.intersectCaps(xs, r, shape)
	}

	// Ray is parallel to y axis
	discriminant := math.Pow(b, 2) - 4*a*c

	// Ray does not intersect cone
	xs := make(Intersections, 0, 2)
	t0 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t1 := (-b + math.Sqrt(discriminant)) / (2 * a)

	// swap min max?
	if t0 > t1 {
		t0, t1 = t1, t0
	}

	y0 := r.Origin.Y + (t0 * r.Direction.Y)
	if cone.Minimum < y0 && y0 < cone.Maximum {
		xs = append(xs, NewIntersection(t0, *shape))
	}

	y1 := r.Origin.Y + (t1 * r.Direction.Y)
	if cone.Minimum < y1 && y1 < cone.Maximum {
		xs = append(xs, NewIntersection(t1, *shape))
	}

	// Caps only matter if the cone is closed, and might possibly be intersected by the ray.
	xs = cone.intersectCaps(xs, r, shape)

	return xs
	// }
}

func (cone Cone) LocalNormalAt(localPoint Tuple) Tuple {
	// ... compute the square of the distance from the y axis ...
	distance := math.Pow(localPoint.X, 2) + math.Pow(localPoint.Z, 2)

	if distance < 1 && localPoint.Y >= cone.Maximum-EPSILON {
		return NewVector(0, 1, 0)
	} else if distance < 1 && localPoint.Y <= cone.Minimum+EPSILON {
		return NewVector(0, -1, 0)
	} else {
		y := math.Sqrt(math.Pow(localPoint.X, 2) + math.Pow(localPoint.Z, 2))
		if localPoint.Y > 0 {
			y = -y
		}
		return NewVector(localPoint.X, y, localPoint.Z)
	}
}

func (cone Cone) localIsEqualTo(c2 ShapeInterface) bool {
	// NB I still don't know why we need to do a type assertion to *Cone
	// instead of Cone, but it fixes this panic when comparing two
	// cones: 'panic: interface conversion: raytracer.ShapeInterface is *raytracer.Cone, not raytracer.Cone [recovered]
	c2Cone := c2.(*Cone)
	if !cone.Origin.IsEqualTo(c2Cone.Origin) {
		return false
	}
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (cone Cone) localType() string {
	return "Cone"
}
