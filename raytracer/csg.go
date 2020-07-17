package raytracer

import (
	"fmt"
	"sort"
)

// Constructive Solid Geometry: combines shapes using one of three operations
type Csg struct {
	Operation string
	Left      *Shape
	Right     *Shape
}

func NewCsg(t string, l, r *Shape) *Shape {
	csg := Csg{t, l, r}
	s := NewShape(csg)
	l.Parent = s
	r.Parent = s
	return s
}

func IntersectionAllowed(op string, lhit, inl, inr bool) bool {
	if op == "union" { // ... preserves all intersections on the exterior of both shapes ...
		// ... only want the intersections that are not inside another object ...
		return (lhit && !inr) || (!lhit && !inl)
	} else if op == "intersection" { // ... preserves all intersections where both shapes overlap ...
		// ... only those intersections that strike one object while inside the other ...
		return (lhit && inr) || (!lhit && inl)
	} else if op == "difference" { // ... preserves all intersections not exclusively inside the object on the right ...
		// ... keep every intersection on left that is not inside right, and every intersection on right that is inside left ..
		return (lhit && !inr) || (!lhit && inl)
	} else {
		panic("Invalid Csg operation!") // TODO: replace eith returned error?
	}
}

// For this to work, your filter_intersections() function needs to loop over each intersection in xs, keeping
// track of which child the intersection hits and which children it is currently inside, and then passing
// that information to intersec- tion_allowed(). If the intersection is allowed, it’s added to the list
// of passing intersections.
func (c *Csg) FilterIntersections(xs Intersections) Intersections {
	inl, inr := false, false      // outside of both children
	filteredXs := Intersections{} // list of filtered intersections

	for _, x := range xs {
		lhit := c.Left.Includes(x.Object)

		if IntersectionAllowed(c.Operation, lhit, inl, inr) {
			filteredXs = append(filteredXs, x)
		}

		if lhit {
			inl = !inl
		} else {
			inr = !inr
		}
	}
	return filteredXs
}

////////////////
// ShapeInterface methods
/////////////////////////

func (c Csg) LocalBounds() BoundingBox {
	b := NullBoundingBox()
	b.AddBoundingBoxes(
		c.Left.ParentSpaceBounds(),
		c.Right.ParentSpaceBounds(),
	)
	return b
}

func (c Csg) LocalIntersect(r Ray, shape *Shape) Intersections {
	// This is the optimization that Csg offers: only calculate its Left/Right
	// intersections if the ray interects the BoundingBox itself.
	if c.LocalBounds().Intersects(r) {
		xs := Intersections{}

		leftXs, rightXs := c.Left.Intersect(r), c.Right.Intersect(r)
		xs = append(xs, leftXs...)
		xs = append(xs, rightXs...)

		// TODO: extract into an Intersections.SortByTime() method?
		sort.Slice(xs, func(i, j int) bool { return xs[i].Time < xs[j].Time })

		return c.FilterIntersections(xs)
	} else {
		return Intersections{}
	}
}

func (c Csg) LocalNormalAt(localPoint Tuple, hit *Intersection) Tuple {
	// TODO: return error instead
	//  ... if your code ever tries to call local_normal_at() on a Csg, that means there’s a bug somewhere (p200) ...
	return NewVector(localPoint.X, localPoint.Y, localPoint.Z)
}

func (c Csg) localIsEqualTo(c2 ShapeInterface) bool {
	c2Csg := c2.(Csg)
	if !c.Left.IsEqualTo(c2Csg.Left) {
		return false
	} else if !c.Right.IsEqualTo(c2Csg.Right) {
		return false
	}
	return true
}

func (c Csg) String() string {
	return fmt.Sprintf("Csg(\n  Operation: %s\n  Left: %v\n  Right: %v\n)", c.Operation, c.Left, c.Right)
}

func (c Csg) localString() string {
	return c.String()
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (c Csg) localType() string {
	return "Csg"
}
