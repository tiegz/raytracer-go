package raytracer

import (
	"fmt"
	"math"
)

// BoundingBox represents a bounding box, which can be used to optimize intersection-checking on groups of objects.
type BoundingBox struct {
	MinPoint Tuple
	MaxPoint Tuple
}

func NewBoundingBox(min Tuple, max Tuple) BoundingBox {
	return BoundingBox{min, max}
}

// NB: this should just be a BoundingBox with no min/max provided, but Go doesn't support optional arguments.
func NullBoundingBox() BoundingBox {
	return NewBoundingBox(
		NewPoint(math.Inf(1), math.Inf(1), math.Inf(1)),
		NewPoint(math.Inf(-1), math.Inf(-1), math.Inf(-1)),
	)
}

// Adds points to this BoundingBox and re-calculates the MinPoint and MaxPoint.
func (b *BoundingBox) AddPoints(points ...Tuple) {
	for _, point := range points {
		b.MinPoint.X = math.Min(b.MinPoint.X, point.X)
		b.MinPoint.Y = math.Min(b.MinPoint.Y, point.Y)
		b.MinPoint.Z = math.Min(b.MinPoint.Z, point.Z)

		b.MaxPoint.X = math.Max(b.MaxPoint.X, point.X)
		b.MaxPoint.Y = math.Max(b.MaxPoint.Y, point.Y)
		b.MaxPoint.Z = math.Max(b.MaxPoint.Z, point.Z)
	}
}

func (b *BoundingBox) AddBoundingBoxes(boundingBoxes ...BoundingBox) {
	for _, bb := range boundingBoxes {
		b.AddPoints(bb.MinPoint)
		b.AddPoints(bb.MaxPoint)
	}
}

// Divide a bounding box into two sub-boxes.
func (b BoundingBox) SplitBounds() (BoundingBox, BoundingBox) {
	dx := b.MaxPoint.X - b.MinPoint.X
	dy := b.MaxPoint.Y - b.MinPoint.Y
	dz := b.MaxPoint.Z - b.MinPoint.Z

	greatestDimmension := maxFloat64(dx, dy, dz)

	x0, y0, z0 := b.MinPoint.X, b.MinPoint.Y, b.MinPoint.Z
	x1, y1, z1 := b.MaxPoint.X, b.MaxPoint.Y, b.MaxPoint.Z

	// ... adjust the points so that they lie on the dividing plane ...
	if greatestDimmension == dx {
		x0 = x0 + dx/2.0
		x1 = x0
	} else if greatestDimmension == dy {
		y0 = y0 + dy/2.0
		y1 = y0
	} else {
		z0 = z0 + dz/2.0
		z1 = z0
	}

	midMin := NewPoint(x0, y0, z0)
	midMax := NewPoint(x1, y1, z1)

	return NewBoundingBox(b.MinPoint, midMax), NewBoundingBox(midMin, b.MaxPoint)
}

func (b BoundingBox) ContainsPoint(point Tuple) bool {
	return (point.X >= b.MinPoint.X && point.X <= b.MaxPoint.X) &&
		(point.Y >= b.MinPoint.Y && point.Y <= b.MaxPoint.Y) &&
		(point.Z >= b.MinPoint.Z && point.Z <= b.MaxPoint.Z)
}

func (b BoundingBox) ContainsBox(b2 BoundingBox) bool {
	return b.ContainsPoint(b2.MinPoint) && b.ContainsPoint(b2.MaxPoint)
}

// Transforms a BoundingBox according to a Matrix. To work for scaling,
// translation and rotation, we need to get all the points from the
// box, transform each by the matrix, and then find the new min/max.
func (b BoundingBox) Transform(m Matrix) BoundingBox {
	p1 := b.MinPoint
	p2 := NewPoint(b.MinPoint.X, b.MinPoint.Y, b.MaxPoint.Z)
	p3 := NewPoint(b.MinPoint.X, b.MaxPoint.Y, b.MinPoint.Z)
	p4 := NewPoint(b.MinPoint.X, b.MaxPoint.Y, b.MaxPoint.Z)
	p5 := NewPoint(b.MaxPoint.X, b.MinPoint.Y, b.MinPoint.Z)
	p6 := NewPoint(b.MaxPoint.X, b.MinPoint.Y, b.MaxPoint.Z)
	p7 := NewPoint(b.MaxPoint.X, b.MaxPoint.Y, b.MinPoint.Z)
	p8 := b.MaxPoint

	b2 := NullBoundingBox()

	for _, p := range []Tuple{p1, p2, p3, p4, p5, p6, p7, p8} {
		b2.AddPoints(m.MultiplyByTuple(p))
	}

	return b2
}

// TODO: we can reuse the Cube LocalIntersect code here if we make
// this return an Interesctions instead.
func (b BoundingBox) Intersects(r Ray) bool {
	xTMin, xTMax := b.checkAxis(r.Origin.X, r.Direction.X, b.MinPoint.X, b.MaxPoint.X)
	yTMin, yTMax := b.checkAxis(r.Origin.Y, r.Direction.Y, b.MinPoint.Y, b.MaxPoint.Y)
	zTMin, zTMax := b.checkAxis(r.Origin.Z, r.Direction.Z, b.MinPoint.Z, b.MaxPoint.Z)

	// ... The intersection of the ray with that square will always be those two points: the largest minimum t value and the smallest maximum t value. ...
	tMin := maxFloat64(xTMin, yTMin, zTMin)
	tMax := minFloat64(xTMax, yTMax, zTMax)

	if tMin > tMax {
		return false
	} else {
		return true
	}
}

func (b BoundingBox) checkAxis(origin float64, direction float64, min float64, max float64) (float64, float64) {
	tMinNumerator := (min - origin)
	tMaxNumerator := (max - origin)
	var tMin, tMax float64

	if math.Abs(direction) >= EPSILON {
		tMin = tMinNumerator / direction
		tMax = tMaxNumerator / direction
	} else {
		tMin = tMinNumerator * math.Inf(1)
		tMax = tMaxNumerator * math.Inf(1)
	}
	if tMin > tMax {
		return tMax, tMin
	} else {
		return tMin, tMax
	}
}

func (b BoundingBox) String() string {
	return fmt.Sprintf("BoundingBox( MinPoint:%v MaxPoint:%v )", b.MinPoint, b.MaxPoint)
}

func (b BoundingBox) IsEqualTo(b2 BoundingBox) bool {
	return b.MinPoint.IsEqualTo(b2.MinPoint) && b.MaxPoint.IsEqualTo(b2.MaxPoint)
}
