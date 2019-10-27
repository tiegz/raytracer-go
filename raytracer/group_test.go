package raytracer

import (
	"testing"
)

func TestCreatingANewGroup(t *testing.T) {
	s := NewGroup()
	g := s.LocalShape.(Group)

	assertEqualMatrix(t, IdentityMatrix(), s.Transform)
	assertEqualInt(t, 0, len(g.Children))
}

func TestAddingAChildToAGroup(t *testing.T) {
	g := NewGroup()
	ts := NewTestShape()
	g.AddChildren(&ts)
	gs := g.LocalShape.(Group)

	assertEqualInt(t, 1, len(gs.Children))
	assert(t, gs.Contains(&ts))
}

func TestIntersectingARayWithAnEmptyGroup(t *testing.T) {
	g := NewGroup()
	gs := g.LocalShape.(Group)
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	xs := gs.LocalIntersect(r, &g)

	assertEqualInt(t, 0, len(xs))
}

func TestIntersectingARayWithANonemptyGroup(t *testing.T) {
	g := NewGroup()
	s1 := NewSphere()
	s2 := NewSphere()
	s2.Transform = NewTranslation(0, 0, -3)
	s3 := NewSphere()
	s3.Transform = NewTranslation(5, 0, 0)

	g.AddChildren(&s1, &s2, &s3)
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	gs := g.LocalShape.(Group)
	xs := gs.LocalIntersect(r, &g)

	assertEqualInt(t, 4, len(xs))
	assertEqualShape(t, xs[0].Object, s2)
	assertEqualShape(t, xs[1].Object, s2)
	assertEqualShape(t, xs[2].Object, s1)
	assertEqualShape(t, xs[3].Object, s1)
}

func TestIntersectingATransformedGroup(t *testing.T) {
	g := NewGroup()
	g.Transform = NewScale(2, 2, 2)
	s := NewSphere()
	s.Transform = NewTranslation(5, 0, 0)
	g.AddChildren(&s)
	r := NewRay(NewPoint(10, 0, -10), NewVector(0, 0, 1))
	xs := g.Intersect(r)

	assertEqualInt(t, 2, len(xs))
}
