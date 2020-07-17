package raytracer

import "fmt"

// SmoothTriangle is like Triangle
type SmoothTriangle struct {
	P1        Tuple
	P2        Tuple
	P3        Tuple
	N1        Tuple
	N2        Tuple
	N3        Tuple
	E1        Tuple
	E2        Tuple
	Normal    Tuple
	boundsMin Tuple
	boundsMax Tuple
}

func NewSmoothTriangle(p1, p2, p3, n1, n2, n3 Tuple) *Shape {
	tri := SmoothTriangle{P1: p1, P2: p2, P3: p3, N1: n1, N2: n2, N3: n3}
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

func (t SmoothTriangle) LocalBounds() BoundingBox {
	return NewBoundingBox(t.boundsMin, t.boundsMax)
}

// Uses the Möller–Trumbore algorithm to find the intersection of the
// ray and the triangle.
func (t SmoothTriangle) LocalIntersect(r Ray, shape *Shape) Intersections {
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

		return Intersections{NewIntersectionWithUV(time, shape, u, v)}
	}
}

func (t SmoothTriangle) LocalNormalAt(localPoint Tuple, hit Intersection) Tuple {
	return t.N2.Multiply(hit.U).
		Add(t.N3.Multiply(hit.V)).
		Add(t.N1.Multiply(1 - hit.U - hit.V))
}

func (t SmoothTriangle) localIsEqualTo(t2 ShapeInterface) bool {
	t2SmoothTriangle := t2.(*SmoothTriangle)
	if !t.P1.IsEqualTo(t2SmoothTriangle.P1) {
		return false
	} else if !t.P2.IsEqualTo(t2SmoothTriangle.P2) {
		return false
	} else if !t.P3.IsEqualTo(t2SmoothTriangle.P3) {
		return false
	} else if !t.P3.IsEqualTo(t2SmoothTriangle.P3) {
		return false
	} else if !t.N1.IsEqualTo(t2SmoothTriangle.N1) {
		return false
	} else if !t.N2.IsEqualTo(t2SmoothTriangle.N2) {
		return false
	} else if !t.N3.IsEqualTo(t2SmoothTriangle.N3) {
		return false
	}
	return true
}

func (t SmoothTriangle) String() string {
	return fmt.Sprintf(
		"SmoothTriangle( P1: %v P2: %v P3: %v N1: %v N2: %v N3: %v E1: %v E2: %v Norma: %v boundsMin: %v boundsMax: %v )",
		t.P1,
		t.P2,
		t.P3,
		t.N1,
		t.N2,
		t.N3,
		t.E1,
		t.E2,
		t.Normal,
		t.boundsMin,
		t.boundsMax,
	)
}

func (t SmoothTriangle) localString() string {
	return t.String()
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (t SmoothTriangle) localType() string {
	return "SmoothTriangle"
}
