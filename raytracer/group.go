package raytracer

import (
	"fmt"
	"sort"
)

type Group struct {
	Children []*Shape
}

func NewGroup() Shape {
	c := []*Shape{}
	return NewShape(Group{c})
}

/////////////////////////
// ShapeInterface methods
/////////////////////////

func (g Group) LocalIntersect(r Ray, shape *Shape) Intersections {
	xs := Intersections{}

	for _, s := range g.Children {
		xs = append(xs, s.Intersect(r)...)
	}

	// TODO: sort all xs by T
	sort.Slice(xs, func(i, j int) bool { return xs[i].Time < xs[j].Time })

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

func (g Group) String() string {
	return fmt.Sprintf("Group( Children:%d )", len(g.Children))
}

func (g Group) localString() string {
	return g.String()
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (g Group) localType() string {
	return "Group"
}
