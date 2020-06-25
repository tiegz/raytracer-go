package raytracer

import (
	"fmt"
	"sort"
)

type Group struct {
	Children    []*Shape
	LeftBounds  BoundingBox
	RightBounds BoundingBox
}

func NewGroup() Shape {
	c := []*Shape{}
	return NewShape(Group{Children: c})
}

func (g Group) String() string {
	return fmt.Sprintf("Group(\nChildren: %d\n  LeftBounds: %v\nRightBounds: %v\n)", len(g.Children), g.LeftBounds, g.RightBounds)
}

/////////////////////////
// ShapeInterface methods
/////////////////////////

func (g Group) LocalIntersect(r Ray, shape *Shape) Intersections {
	xs := Intersections{}

	// This is the optimization that Groups offers: only calculate its Children
	// intersections if the ray interects the BoundingBox itself.
	if g.LocalBounds().Intersects(r) {
		for _, s := range g.Children {
			xs = append(xs, s.Intersect(r)...)
		}
		sort.Slice(xs, func(i, j int) bool { return xs[i].Time < xs[j].Time })
	}

	return xs
}

func (g Group) LocalNormalAt(localPoint Tuple, hit Intersection) Tuple {
	// TODO: return error instead
	//  ... if your code ever tries to call local_normal_at() on a group, that means there’s a bug somewhere (p200) ...
	return NewVector(localPoint.X, localPoint.Y, localPoint.Z)
}

func (g Group) localIsEqualTo(s2 ShapeInterface) bool {
	s2Group := s2.(Group)
	if len(g.Children) != len(s2Group.Children) {
		return false
	} else {
		// TODO: test to confirm this works?
		for idx, childShape := range g.Children {
			if childShape != s2Group.Children[idx] {
				return false
			}
		}
	}
	return true
}

func (g Group) LocalBounds() BoundingBox {
	b := NullBoundingBox()
	for _, c := range g.Children {
		b.AddBoundingBoxes(c.ParentSpaceBounds())
	}
	return b
}

func (g Group) localString() string {
	return g.String()
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (g Group) localType() string {
	return "Group"
}
