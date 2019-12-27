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
	assert(t, g.Includes(&ts))
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
	s2.SetTransform(NewTranslation(0, 0, -3))
	s3 := NewSphere()
	s3.SetTransform(NewTranslation(5, 0, 0))

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
	g.SetTransform(NewScale(2, 2, 2))
	s := NewSphere()
	s.SetTransform(NewTranslation(5, 0, 0))
	g.AddChildren(&s)
	r := NewRay(NewPoint(10, 0, -10), NewVector(0, 0, 1))
	xs := g.Intersect(r)

	assertEqualInt(t, 2, len(xs))
}

func TestAGroupHasABoundingBoxThatContainsItsChildren(t *testing.T) {
	s := NewSphere()
	s.SetTransform(NewTranslation(2, 5, -3))
	s.SetTransform(s.Transform.Multiply(NewScale(2, 2, 2)))
	c := NewCylinder()
	cc := c.LocalShape.(*Cylinder)
	cc.Minimum = -2
	cc.Maximum = 2
	c.SetTransform(NewTranslation(-4, -1, 4))
	c.SetTransform(c.Transform.Multiply(NewScale(0.5, 1, 0.5)))
	g := NewGroup()
	g.AddChildren(&s, &c)

	b := g.Bounds()

	assertEqualTuple(t, NewPoint(-4.5, -3, -5), b.MinPoint)
	assertEqualTuple(t, NewPoint(4, 7, 4.5), b.MaxPoint)
}

func TestIntersectingRayAndGroupDoesntTestChildrenIfBoxIsMissed(t *testing.T) {
	ts := NewTestShape()
	g := NewGroup()
	g.AddChildren(&ts)
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 1, 0))

	g.Intersect(r)

	assertEqualRay(t, NullRay(), ts.SavedRay)
}
func TestIntersectingRayAndGroupDoesntTestsChildrenIfBoxIsHit(t *testing.T) {
	ts := NewTestShape()
	g := NewGroup()
	g.AddChildren(&ts)
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))

	g.Intersect(r)

	assertEqualRay(t, r, ts.SavedRay)
}

func TestCreatingASubGroupFromAListOfChildren(t *testing.T) {
	s1 := NewSphere()
	s2 := NewSphere()
	g := NewGroup()
	g.MakeSubGroup(&s1, &s2)
	group := g.LocalShape.(Group)

	expectedSubGroup := NewGroup()
	expectedSubGroup.AddChildren(&s1, &s2)

	assertEqualInt(t, 1, len(group.Children))
	assertEqualShape(t, expectedSubGroup, *group.Children[0])
}

func TestSubdividingAPrimitiveDoesNothing(t *testing.T) {
	shape := NewSphere()
	shape.Divide(1)

	assertEqualString(t, "Sphere", shape.LocalShape.localType())
}

func TestSubdividingAGroupPartitionsItsChildren(t *testing.T) {
	s1 := NewSphere()
	s1.SetTransform(NewTranslation(-2, -2, 0))
	s2 := NewSphere()
	s2.SetTransform(NewTranslation(-2, 2, 0))
	s3 := NewSphere()
	s3.SetTransform(NewScale(4, 4, 4))
	g := NewGroup()
	g.AddChildren(&s1, &s2, &s3)
	g.Divide(1)
	group := g.LocalShape.(Group)

	assertEqualShape(t, s3, *group.Children[0])

	subGroup := group.Children[1].LocalShape.(Group)
	assertEqualString(t, "Group", subGroup.localType())
	assertEqualInt(t, 2, len(subGroup.Children))

	subSubGroup1 := subGroup.Children[0].LocalShape.(Group)
	subSubGroup2 := subGroup.Children[1].LocalShape.(Group)

	assertEqualInt(t, 1, len(subSubGroup1.Children))
	assertEqualShape(t, s1, *subSubGroup1.Children[0])
	assertEqualInt(t, 1, len(subSubGroup2.Children))
	assertEqualShape(t, s2, *subSubGroup2.Children[0])
}

func TestSubdividingAGroupWithTooFewChildren(t *testing.T) {
	s1 := NewSphere()
	s1.SetTransform(NewTranslation(-2, 0, 0))
	s2 := NewSphere()
	s2.SetTransform(NewTranslation(2, 1, 0))
	s3 := NewSphere()
	s3.SetTransform(NewTranslation(2, -1, 0))
	sg := NewGroup()
	sg.AddChildren(&s1, &s2, &s3)
	s4 := NewSphere()
	g := NewGroup()
	g.AddChildren(&sg, &s4)
	g.Divide(3)

	group := g.LocalShape.(Group)
	subGroup := sg.LocalShape.(Group)

	assertEqualShape(t, sg, *group.Children[0])
	assertEqualShape(t, s4, *group.Children[1])

	assertEqualInt(t, 2, len(subGroup.Children))
	assertEqualGroup(t, []*Shape{&s1}, *subGroup.Children[0])
	assertEqualGroup(t, []*Shape{&s2, &s3}, *subGroup.Children[1])
}

func TestSubdividingACsgShapeSubdividesItsChildren(t *testing.T) {
	s1 := NewSphere()
	s1.SetTransform(NewTranslation(-1.5, 0, 0))
	s1.Label = "s1"
	s2 := NewSphere()
	s2.SetTransform(NewTranslation(1.5, 0, 0))
	s2.Label = "s2"
	l := NewGroup()
	l.AddChildren(&s1, &s2)

	s3 := NewSphere()
	s3.SetTransform(NewTranslation(0, 0, -1.5))
	s3.Label = "s3"
	s4 := NewSphere()
	s4.SetTransform(NewTranslation(0, 0, 1.5))
	s4.Label = "s4"
	r := NewGroup()
	r.AddChildren(&s3, &s4)

	shape := NewCsg("difference", &l, &r)

	shape.Divide(1)

	assertEqualGroup(t, []*Shape{&s1}, *l.LocalShape.(Group).Children[0])
	assertEqualGroup(t, []*Shape{&s2}, *l.LocalShape.(Group).Children[1])
	assertEqualGroup(t, []*Shape{&s3}, *r.LocalShape.(Group).Children[0])
	assertEqualGroup(t, []*Shape{&s4}, *r.LocalShape.(Group).Children[1])
}
