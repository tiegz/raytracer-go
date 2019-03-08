package raytracer

import (
	"testing"
)

func TestNewWorld(t *testing.T) {
	world := NewWorld()

	assertEqualInt(t, 0, len(world.Objects))
	assertEqualInt(t, 0, len(world.Lights))
}

func TestDefaultWorld(t *testing.T) {
	world := DefaultWorld()

	l := NewPointLight(NewPoint(-10, 10, -10), Colors["White"])

	s1 := NewSphere()
	s1.Material = Material{NewColor(0.8, 1.0, 0.6), 0, 0.7, 0.2, 0}

	s2 := NewSphere()
	s2.Transform = NewScale(0.5, 0.5, 0.5)

	assertEqualPointLight(t, l, world.Lights[0])
	assert(t, world.Contains(s1))
	assert(t, world.Contains(s2))
}

func TestIntersectWorldWithRay(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	xs := w.Intersect(r)

	assertEqualInt(t, 4, len(xs))
	assertEqualFloat64(t, 4, xs[0].Time)
	assertEqualFloat64(t, 4.5, xs[1].Time)
	assertEqualFloat64(t, 5.5, xs[2].Time)
	assertEqualFloat64(t, 6, xs[3].Time)
}

func TestShadingAnIntersection(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s1 := w.Objects[0]
	i := NewIntersection(4, s1)
	c := i.PrepareComputations(r)

	actual := w.ShadeHit(c)
	expected := NewColor(0.38066, 0.47583, 0.2855)
	assertEqualColor(t, expected, actual)
}

func TestShadingAnIntersectionFromInside(t *testing.T) {
	w := DefaultWorld()
	w.Lights = []PointLight{NewPointLight(NewPoint(0, 0.25, 0), Colors["White"])}

	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	s2 := w.Objects[1]
	i := NewIntersection(0.5, s2)
	c := i.PrepareComputations(r)

	actual := w.ShadeHit(c)
	expected := NewColor(0.90498, 0.90498, 0.90498)
	assertEqualColor(t, expected, actual)
}
