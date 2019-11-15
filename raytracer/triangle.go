package raytracer

import "fmt"

type Triangle struct {
	P1        Tuple
	P2        Tuple
	P3        Tuple
	E1        Tuple
	E2        Tuple
	Normal    Tuple
	boundsMin Tuple
	boundsMax Tuple
}

func NewTriangle(p1, p2, p3 Tuple) Shape {
	tri := Triangle{P1: p1, P2: p2, P3: p3}
	// Pre-calculate edge vectors and normal
	tri.E1 = p2.Subtract(p1)
	tri.E2 = p3.Subtract(p1)
	cross := tri.E2.Cross(tri.E1)
	tri.Normal = cross.Normalized()
	tri.boundsMin = NewPoint(
		minFloat64(p1.X, p2.X, p3.X),
		minFloat64(p1.Y, p2.Y, p3.Y),
		minFloat64(p1.Z, p2.Z, p3.Z),
	)
	tri.boundsMax = NewPoint(
		maxFloat64(p1.X, p2.X, p3.X),
		maxFloat64(p1.Y, p2.Y, p3.Y),
		maxFloat64(p1.Z, p2.Z, p3.Z),
	)
	return NewShape(&tri)
}

/////////////////////////
// ShapeInterface methods
/////////////////////////

func (t Triangle) LocalBounds() BoundingBox {
	return NewBoundingBox(t.boundsMin, t.boundsMax)
}

// Uses the Möller–Trumbore algorithm to find the intersection of the
// ray and the triangle.
func (t Triangle) LocalIntersect(r Ray, shape *Shape) Intersections {
	directionCrossE2 := r.Direction.Cross(t.E2)
	determinant := directionCrossE2.Dot(t.E1)

	if equalFloat64s(determinant, 0) { // ray is parallel to triangle
		return Intersections{}
	} else {
		f := 1.0 / determinant
		p1ToOrigin := r.Origin.Subtract(t.P1)
		u := f * p1ToOrigin.Dot(directionCrossE2)
		if u < EPSILON || u > 1 {
			return Intersections{}
		}

		originCrossE1 := p1ToOrigin.Cross(t.E1)
		v := f * r.Direction.Dot(originCrossE1)
		if v < EPSILON || (u+v) > 1 {
			return Intersections{}
		}

		time := f * t.E2.Dot(originCrossE1)
		return Intersections{NewIntersection(time, *shape)}
	}
}

func (t Triangle) LocalNormalAt(worldPoint Tuple) Tuple {
	return t.Normal
}

func (t Triangle) localIsEqualTo(t2 ShapeInterface) bool {
	t2Triangle := t2.(*Triangle)
	if !t.P1.IsEqualTo(t2Triangle.P1) {
		return false
	} else if !t.P2.IsEqualTo(t2Triangle.P2) {
		return false
	} else if !t.P3.IsEqualTo(t2Triangle.P3) {
		return false
	}
	return true
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle( %v %v %v )", t.P1, t.P2, t.P3)
}

func (t Triangle) localString() string {
	return t.String()
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (t Triangle) localType() string {
	return "Triangle"
}
