package raytracer

import (
	"math"
	"testing"
)

func TestDefaultTransformation(t *testing.T) {
	s := NewTestShape()

	assertEqualMatrix(t, s.Transform, IdentityMatrix())
}

func TestAssigningATransformation(t *testing.T) {
	s := NewTestShape()
	s.SetTransform(NewTranslation(2, 3, 4))

	assertEqualMatrix(t, NewTranslation(2, 3, 4), s.Transform)
}

// These replace the tests named “A sphere has a default material” and “A sphere may be assigned a material” (from the sphere scenarios on page 85).
func TestDefaultMatrial(t *testing.T) {
	s := NewTestShape()

	assertEqualMaterial(t, s.Material, DefaultMaterial())
}

func TestAssigningAMaterial(t *testing.T) {
	s := NewTestShape()
	m := DefaultMaterial()
	m.Ambient = 1
	s.Material = m

	assertEqualMaterial(t, s.Material, m)
}

//  These tests are both based on (and replace) the tests called “Intersecting a scaled sphere with a ray” and “Intersecting a translated sphere with a ray”
func TestIntersectingAScaledShapeWithARay(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewTestShape()
	s.SetTransform(NewScale(2, 2, 2))
	s.Intersect(r)

	assertEqualTuple(t, NewPoint(0, 0, -2.5), s.SavedRay.Origin)
	assertEqualTuple(t, NewVector(0, 0, 0.5), s.SavedRay.Direction)
}

func TestIntersectingATranslatedShapeWithARay(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewTestShape()
	s.SetTransform(NewTranslation(5, 0, 0))
	s.Intersect(r)

	assertEqualTuple(t, NewPoint(-5, 0, -5), s.SavedRay.Origin)
	assertEqualTuple(t, NewVector(0, 0, 1), s.SavedRay.Direction)
}

// The following two tests replace the ones called “Computing the normal on a translated sphere” and “Computing the normal on a transformed sphere”
func TestComputingTheNormalOnATranslatedShape(t *testing.T) {
	s := NewTestShape()
	s.SetTransform(NewTranslation(0, 1, 0))
	n := s.NormalAt(NewPoint(0, 1.70711, -0.70711), NewIntersection(0, s))

	assertEqualTuple(t, NewVector(0, 0.70711, -0.70711), n)
}

func TestComputingTheNormalOnATransformedShape(t *testing.T) {
	s := NewTestShape()
	transform := NewScale(1, 0.5, 1)
	transform = transform.Multiply(NewRotateZ(math.Pi / 5))
	s.SetTransform(transform)
	n := s.NormalAt(NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2), NewIntersection(0, s))

	assertEqualTuple(t, NewVector(0, 0.97014, -0.24254), n)
}

func TestAShapeHasAParentAttribute(t *testing.T) {
	s := NewTestShape()

	assertNil(t, s.Parent)
}

func TestConvertingAPointFromWorldToObjectSpace(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(NewRotateY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(NewScale(2, 2, 2))
	g1.AddChildren(g2)

	s := NewSphere()
	s.SetTransform(NewTranslation(5, 0, 0))
	g2.AddChildren(s)

	p := s.WorldToObject(NewPoint(-2, 0, -10))

	assertEqualTuple(t, NewPoint(0, 0, -1), p)
}

func TestConvertingANormalFromObjectToWorldSpace(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(NewRotateY(math.Pi / 2))

	g2 := NewGroup()
	g2.SetTransform(NewScale(1, 2, 3))
	g1.AddChildren(g2)

	s := NewSphere()
	s.SetTransform(NewTranslation(5, 0, 0))
	g2.AddChildren(s)

	n := s.NormalToWorld(NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	// NB: values in book were slightly different: (0.2857, 0.4286, -0.8571)
	assertEqualTuple(t, NewVector(0.28571, 0.42857, -0.85714), n)
}

func TestFindingTheNormalOnAChildObject(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(NewRotateY(math.Pi / 2))

	g2 := NewGroup()
	g2.SetTransform(NewScale(1, 2, 3))
	g1.AddChildren(g2)

	s := NewSphere()
	s.SetTransform(NewTranslation(5, 0, 0))
	g2.AddChildren(s)

	n := s.NormalAt(NewPoint(1.7321, 1.1547, -5.5774), NewIntersection(0, s))

	// NB: values in book were slightly different: (0.2857, 0.4286, -0.8571)
	assertEqualTuple(t, NewVector(0.28570, 0.42854, -0.85716), n)
}

func TestIncludesInspectsShapesChildren(t *testing.T) {
	sphere := NewSphere()
	cylinder := NewCylinder()
	cube := NewCube()
	group1 := NewGroup()
	group1.AddChildren(sphere)
	group2 := NewGroup()
	group2.AddChildren(group1)
	csg := NewCsg("intersection", sphere, cube)

	// group
	assert(t, group1.Includes(sphere))
	assert(t, group2.Includes(group1))
	assert(t, group2.Includes(sphere))

	// csg
	assert(t, csg.Includes(sphere))
	assert(t, csg.Includes(cube))
	assert(t, !sphere.Includes(cylinder))

	// etc
	assert(t, sphere.Includes(sphere))
	assert(t, !sphere.Includes(cylinder))
}

func TestFindingTheBoundsOnAnObject(t *testing.T) {
	s := NewSphere()

	assertEqualBoundingBox(t, NewBoundingBox(NewPoint(-1, -1, -1), NewPoint(1, 1, 1)), s.Bounds())
}

func TestTestShapeHasArbitraryBounds(t *testing.T) {
	s := NewTestShape()
	b := s.Bounds()

	assertEqualTuple(t, NewPoint(-1, -1, -1), b.MinPoint)
	assertEqualTuple(t, NewPoint(1, 1, 1), b.MaxPoint)
}

func TestQueryingAShapesBoundingBoxInItsParentsSpace(t *testing.T) {
	s := NewSphere()
	s.SetTransform(NewTranslation(1, -3, 5))
	s.SetTransform(s.Transform.Multiply(NewScale(0.5, 2, 4)))

	b := s.ParentSpaceBounds()

	assertEqualTuple(t, NewPoint(0.5, -5, 1), b.MinPoint)
	assertEqualTuple(t, NewPoint(1.5, -1, 9), b.MaxPoint)
}

/////////////
// Benchmarks
/////////////

func BenchmarkShapeMethodIsEqualTo(b *testing.B) {
	s := NewTestShape()
	for i := 0; i < b.N; i++ {
		s.IsEqualTo(s)
	}
}

func BenchmarkShapeMethodIntersect(b *testing.B) {
	// Taken from TestIntersectingAScaledShapeWithARay().
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewTestShape()
	s.SetTransform(NewScale(2, 2, 2))
	for i := 0; i < b.N; i++ {
		s.Intersect(r)
	}
}

func BenchmarkShapeMethodNormalAt(b *testing.B) {
	// Taken from TestComputingTheNormalOnATranslatedShape().
	s := NewTestShape()
	s.SetTransform(NewTranslation(0, 1, 0))
	for i := 0; i < b.N; i++ {
		s.NormalAt(NewPoint(0, 1.70711, -0.70711), NewIntersection(0, s))
	}
}

func BenchmarkShapeMethodWorldToObject(b *testing.B) {
	// Taken from TestConvertingAPointFromWorldToObjectSpace().
	g1 := NewGroup()
	g1.SetTransform(NewRotateY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(NewScale(2, 2, 2))
	g1.AddChildren(g2)
	s := NewSphere()
	s.SetTransform(NewTranslation(5, 0, 0))
	g2.AddChildren(s)
	for i := 0; i < b.N; i++ {
		s.WorldToObject(NewPoint(-2, 0, -10))
	}
}

func BenchmarkShapeMethodNormalToWorld(b *testing.B) {
	// Taken from TestConvertingANormalFromObjectToWorldSpace().
	g1 := NewGroup()
	g1.SetTransform(NewRotateY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(NewScale(1, 2, 3))
	g1.AddChildren(g2)
	s := NewSphere()
	s.SetTransform(NewTranslation(5, 0, 0))
	g2.AddChildren(s)

	for i := 0; i < b.N; i++ {
		s.NormalToWorld(NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	}
}
