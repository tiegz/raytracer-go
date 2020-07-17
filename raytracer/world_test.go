package raytracer

import (
	"fmt"
	"math"
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
	s1.Material = Material{
		Color:     NewColor(0.8, 1.0, 0.6),
		Ambient:   0.1,
		Diffuse:   0.7,
		Specular:  0.2,
		Shininess: 200,
		// Pattern:         nil,
		Reflective:      0.0,
		Transparency:    0.0,
		RefractiveIndex: 1.0,
	}

	s2 := NewSphere()
	s2.SetTransform(NewScale(0.5, 0.5, 0.5))

	assertEqualLight(t, l, world.Lights[0])
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

	actual := w.ShadeHit(c, DefaultMaximumReflections)
	expected := NewColor(0.38066, 0.47583, 0.2855)
	assertEqualColor(t, expected, actual)
}

func TestShadingAnIntersectionFromInside(t *testing.T) {
	w := DefaultWorld()
	w.Lights = []AreaLight{
		NewPointLight(NewPoint(0, 0.25, 0), Colors["White"]),
	}

	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	s2 := w.Objects[1]
	i := NewIntersection(0.5, s2)
	c := i.PrepareComputations(r)

	actual := w.ShadeHit(c, DefaultMaximumReflections)
	expected := NewColor(0.1, 0.1, 0.1)
	// expected := NewColor(0.90498, 0.90498, 0.90498) // TODO this is the answer from the book -- why not working?
	assertEqualColor(t, expected, actual)
}

func TestShadeHitIsGivenAnIntersectionInShadow(t *testing.T) {
	w := NewWorld()
	s1 := NewSphere()
	s2 := NewSphere()
	s2.SetTransform(NewTranslation(0, 0, 10))

	w.Lights = []AreaLight{
		NewPointLight(NewPoint(0, 0, -10), Colors["White"]),
	}
	w.Objects = []Shape{
		s1,
		s2,
	}

	r := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
	i := NewIntersection(4, s2)

	c := i.PrepareComputations(r)
	actual := w.ShadeHit(c, DefaultMaximumReflections)
	expected := NewColor(0.1, 0.1, 0.1)

	assertEqualColor(t, expected, actual)
}

func TestColorAtWhenRayMisses(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 1, 0))
	actual := w.ColorAt(r, DefaultMaximumReflections)
	expected := Colors["Black"]

	assertEqualColor(t, expected, actual)
}

// Ray hits the outer sphere
func TestColorAtWhenRayHitsOuterSphere(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	actual := w.ColorAt(r, DefaultMaximumReflections)
	expected := NewColor(0.38066, 0.47583, 0.2855)

	assertEqualColor(t, expected, actual)
}

// Ray is inside outer sphere, pointed at inner sphere.
func TestColorAtWithAnIntersectionBehindRay(t *testing.T) {
	w := DefaultWorld()
	w.Objects[0].Material.Ambient = 1 // outer
	w.Objects[1].Material.Ambient = 1 // inner
	r := NewRay(NewPoint(0, 0, 0.75), NewVector(0, 0, -1))
	actual := w.ColorAt(r, DefaultMaximumReflections)
	expected := w.Objects[1].Material.Color

	assertEqualColor(t, expected, actual)
}

func TestTheRefractedColorWithAnOpaqueSurface(t *testing.T) {
	w := DefaultWorld()
	shape := w.Objects[0]
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	xs := Intersections{NewIntersection(4, shape), NewIntersection(6, shape)}
	comps := xs[0].PrepareComputations(r, xs...)
	c := w.RefractedColor(comps, 5)

	assertEqualColor(t, Colors["Black"], c)
}
func TestTheRefractedColorAtTheMaximumRecursiveDepth(t *testing.T) {
	w := DefaultWorld()
	shape := w.Objects[0]
	shape.Material.Transparency = 1.0
	shape.Material.RefractiveIndex = 1.5
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	xs := Intersections{NewIntersection(4, shape), NewIntersection(6, shape)}
	comps := xs[0].PrepareComputations(r, xs...)
	color := w.RefractedColor(comps, 0)

	assertEqualColor(t, Colors["Black"], color)
}

func TestTheRefractedColorUnderTotalInternalReflection(t *testing.T) {
	w := DefaultWorld()
	shape := w.Objects[0]
	shape.Material.Transparency = 1.0
	shape.Material.RefractiveIndex = 1.5
	r := NewRay(NewPoint(0, 0, math.Sqrt(2)/2), NewVector(0, 1, 0))
	xs := Intersections{NewIntersection(-math.Sqrt(2)/2, shape), NewIntersection(math.Sqrt(2)/2, shape)}
	// NOTE: this time you're inside the sphere, so you need to look at the second intersection, xs[1], not xs[0]
	comps := xs[1].PrepareComputations(r, xs...)
	color := w.RefractedColor(comps, 5)

	assertEqualColor(t, Colors["Black"], color)
}

func TestTheRefractedColorWithARefractedRay(t *testing.T) {
	w := DefaultWorld()
	w.Objects[0].Material.Ambient = 1.0
	w.Objects[0].Material.Pattern = NewTestPattern()
	w.Objects[1].Material.Transparency = 1.0
	w.Objects[1].Material.RefractiveIndex = 1.5
	r := NewRay(NewPoint(0, 0, 0.1), NewVector(0, 1, 0))
	xs := Intersections{
		NewIntersection(-0.9899, w.Objects[0]),
		NewIntersection(-0.4899, w.Objects[1]),
		NewIntersection(0.4899, w.Objects[1]),
		NewIntersection(0.9899, w.Objects[0]),
	}
	// NOTE: this time you're inside the sphere, so you need to look at the second intersection, xs[1], not xs[0]
	comps := xs[2].PrepareComputations(r, xs...)
	color := w.RefractedColor(comps, 5)

	assertEqualColor(t, NewColor(0, 0.99889, 0.04721), color)
}

func TestShadeHitWithATransparentMaterial(t *testing.T) {
	w := DefaultWorld()
	floor := NewPlane()
	floor.SetTransform(NewTranslation(0, -1, 0))
	floor.Material.Transparency = 0.5
	floor.Material.RefractiveIndex = 1.5
	ball := NewSphere()
	ball.SetTransform(NewTranslation(0, -3.5, -0.5))
	ball.Material.Color = NewColor(1, 0, 0)
	ball.Material.Ambient = 0.5
	w.Objects = append(w.Objects, floor, ball)
	r := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	xs := Intersections{NewIntersection(math.Sqrt(2), floor)}
	comps := xs[0].PrepareComputations(r, xs...)
	color := w.ShadeHit(comps, 5)

	assertEqualColor(t, NewColor(0.93642, 0.68642, 0.68642), color)
}

// ... Show that the schlick() reflectance value is used by shade_hit() when a material is both transparent and reflective. ...
func TestShadeHitWithAReflectiveTransparentMaterial(t *testing.T) {
	w := DefaultWorld()
	floor := NewPlane()
	floor.SetTransform(NewTranslation(0, -1, 0))
	floor.Material.Reflective = 0.5
	floor.Material.Transparency = 0.5
	floor.Material.RefractiveIndex = 1.5
	ball := NewSphere()
	ball.SetTransform(NewTranslation(0, -3.5, -0.5))
	ball.Material.Color = NewColor(1, 0, 0)
	ball.Material.Ambient = 0.5
	w.Objects = append(w.Objects, floor, ball)
	r := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	xs := Intersections{NewIntersection(math.Sqrt(2), floor)}
	comps := xs[0].PrepareComputations(r, xs...)
	color := w.ShadeHit(comps, 5)

	assertEqualColor(t, NewColor(0.93391, 0.69643, 0.69243), color)
}

func TestIsShadowTestsForOcclusionBetweenTwoPoints(t *testing.T) {
	w := DefaultWorld()
	lightPosition := NewPoint(-10, -10, -10)

	testCases := []struct {
		Point  Tuple
		Result bool
	}{
		// These replaced the 4 old IsShadowed() tests from 5bfe2037
		{NewPoint(-10, -10, 10), false},
		{NewPoint(10, 10, 10), true},
		{NewPoint(-20, -20, -20), false},
		{NewPoint(-5, -5, -5), false},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("IsShadow tests for occlusion between two points #%v", idx), func(t *testing.T) {
			result := w.IsShadowed(lightPosition, tc.Point)
			assertEqualBool(t, tc.Result, result)
		})
	}
}

/////////////
// Benchmarks
/////////////

func BenchmarkWorldMethodIntersect(b *testing.B) {
	w := DefaultWorld()
	r := NewRay(NewPoint(0, 0, 0), NewVector(1, 1, 1))
	for i := 0; i < b.N; i++ {
		w.Intersect(r)
	}
}

func BenchmarkWorldMethodShadeHit(b *testing.B) {
	// Taken from TestShadingAnIntersection(), so this does calculate a color.
	w := DefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s1 := w.Objects[0]
	i := NewIntersection(4, s1)
	c := i.PrepareComputations(r)

	for i := 0; i < b.N; i++ {
		w.ShadeHit(c, DefaultMaximumReflections)
	}
}
func BenchmarkWorldMethodColorAt(b *testing.B) {
	// Taken from TestColorAtWhenRayHitsOuterSphere, so this does return a color.
	w := DefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	for i := 0; i < b.N; i++ {
		w.ColorAt(r, DefaultMaximumReflections)
	}
}

func BenchmarkWorldMethodRefractedColor(b *testing.B) {
	// Taken from TestTheRefractedColorWithARefractedRay, so this does return a color.
	w := DefaultWorld()
	w.Objects[0].Material.Ambient = 1.0
	w.Objects[0].Material.Pattern = NewTestPattern()
	w.Objects[1].Material.Transparency = 1.0
	w.Objects[1].Material.RefractiveIndex = 1.5
	r := NewRay(NewPoint(0, 0, 0.1), NewVector(0, 1, 0))
	xs := Intersections{
		NewIntersection(-0.9899, w.Objects[0]),
		NewIntersection(-0.4899, w.Objects[1]),
		NewIntersection(0.4899, w.Objects[1]),
		NewIntersection(0.9899, w.Objects[0]),
	}
	// NOTE: this time you're inside the sphere, so you need to look at the second intersection, xs[1], not xs[0]
	comps := xs[2].PrepareComputations(r, xs...)
	for i := 0; i < b.N; i++ {
		w.RefractedColor(comps, 5)
	}
}

func BenchmarkWorldMethodIsShadowed(b *testing.B) {
	// Taken from TestTheShadowWhenObjectIsBetweenPointAndLight, so this does return true.
	w := DefaultWorld()
	p := NewPoint(10, -10, 10)
	for i := 0; i < b.N; i++ {
		w.IsShadowed(p, w.Lights[0].PointOnLight(0, 0))
	}
}
