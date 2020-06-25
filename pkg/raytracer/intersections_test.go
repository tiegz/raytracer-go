package raytracer

import (
	"math"
	"testing"
)

func TestNewIntersection(t *testing.T) {
	sphere := NewSphere()
	i := NewIntersection(1.23, sphere)

	assertEqualFloat64(t, 1.23, i.Time)
	assertEqualShape(t, sphere, i.Object)
}

func TestAggregatingIntersections(t *testing.T) {
	sphere := NewSphere()
	i1 := NewIntersection(1, sphere)
	i2 := NewIntersection(2, sphere)
	intersections := Intersections{i1, i2}

	assertEqualInt(t, 2, len(intersections))
	assertEqualFloat64(t, 1, intersections[0].Time)
	assertEqualFloat64(t, 2, intersections[1].Time)
}

func TestHitWithAllIntersectionsHavingPositiveT(t *testing.T) {
	sphere := NewSphere()
	i1 := NewIntersection(1, sphere)
	i2 := NewIntersection(2, sphere)
	intersections := Intersections{i1, i2}

	hit := intersections.Hit(false)
	assertEqualIntersection(t, i1, hit)
}

func TestHitWithSomeIntersectionsHavingNegativeT(t *testing.T) {
	sphere := NewSphere()
	i1 := NewIntersection(-1, sphere)
	i2 := NewIntersection(1, sphere)
	intersections := Intersections{i1, i2}

	hit := intersections.Hit(false)
	assertEqualIntersection(t, i2, hit)
}

func TestHitWhenAllIntersectionsHaveNegativeT(t *testing.T) {
	sphere := NewSphere()
	i1 := NewIntersection(-2, sphere)
	i2 := NewIntersection(-1, sphere)
	intersections := Intersections{i1, i2}

	hit := intersections.Hit(false)
	assert(t, hit.IsNull())
}

func TestHitIsAlwaysLowestNonNegativeIntersection(t *testing.T) {
	sphere := NewSphere()
	i1 := NewIntersection(5, sphere)
	i2 := NewIntersection(7, sphere)
	i3 := NewIntersection(-3, sphere)
	i4 := NewIntersection(2, sphere)
	intersections := Intersections{i1, i2, i3, i4}

	hit := intersections.Hit(false)
	assertEqualIntersection(t, i4, hit)
}

func TestPrecomputingStateOfIntersection(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewSphere()
	i := NewIntersection(4, s)
	c := i.PrepareComputations(r)

	assertEqualFloat64(t, i.Time, c.Time)
	assertEqualShape(t, s, c.Object)
	assertEqualTuple(t, NewPoint(0, 0, -1), c.Point)
	assertEqualTuple(t, NewVector(0, 0, -1), c.EyeV)
	assertEqualTuple(t, NewVector(0, 0, -1), c.NormalV)
}

func TestHitWhenIntersectionOccursOnOutside(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewSphere()
	i := NewIntersection(4, s)
	c := i.PrepareComputations(r)

	assert(t, !c.Inside)
}

func TestHitWhenIntersectionOccursOnInside(t *testing.T) {
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	s := NewSphere()
	i := NewIntersection(1, s)
	c := i.PrepareComputations(r)

	assertEqualShape(t, s, c.Object)
	assertEqualTuple(t, NewPoint(0, 0, 1), c.Point)
	assertEqualTuple(t, NewVector(0, 0, -1), c.EyeV)
	assertEqualTuple(t, NewVector(0, 0, -1), c.NormalV)
	assert(t, c.Inside)
}

// NB to avoid "raytracer acne" from shadows
func TestHitShouldOffsetPoint(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewSphere()
	s.SetTransform(NewTranslation(0, 0, 1))
	i := NewIntersection(5, s)
	c := i.PrepareComputations(r)

	assert(t, c.OverPoint.Z < -EPSILON/2)
	assert(t, c.Point.Z > c.OverPoint.Z)
}

func TestPrecomputingTheReflectionVector(t *testing.T) {
	shape := NewPlane()
	r := NewRay(NewPoint(0, 1, -1), NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2)) // ray arrives at 45° angle
	i := NewIntersection(math.Sqrt(2), shape)                                      // √2 units away, given the above line + pythagoream theorem
	c := i.PrepareComputations(r)

	assertEqualTuple(t, NewVector(0, math.Sqrt(2)/2, math.Sqrt(2)/2), c.ReflectV) // ray reflects at 45° angle
}

func TestTheReflectedColorForANonreflectiveMaterial(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	shape := w.Objects[1]
	i := NewIntersection(1, shape)
	c := i.PrepareComputations(r)

	assertEqualColor(t, NewColor(0, 0, 0), w.ReflectedColor(c, DefaultMaximumReflections))
}

func TestTheReflectedColorForAReflectiveMaterial(t *testing.T) {
	w := DefaultWorld()
	shape := NewPlane()
	shape.Material.Reflective = 0.5
	shape.SetTransform(NewTranslation(0, -1, 0))
	w.Objects = append(w.Objects, shape)

	r := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	i := NewIntersection(math.Sqrt(2), shape)
	c := i.PrepareComputations(r)

	// ERRATA book values were: 0.19032, 0.2379, 0.14274
	assertEqualColor(t, NewColor(0.19033, 0.23791, 0.14274), w.ReflectedColor(c, DefaultMaximumReflections))
}

func TestShadeHitWithAReflectiveMaterial(t *testing.T) {
	w := DefaultWorld()
	shape := NewPlane()
	shape.Material.Reflective = 0.5
	shape.SetTransform(NewTranslation(0, -1, -0))
	w.Objects = append(w.Objects, shape)
	r := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	i := NewIntersection(math.Sqrt(2), shape)
	c := i.PrepareComputations(r)

	// ERRATA book values were: 0.87677, 0.92436, 0.82918
	assertEqualColor(t, NewColor(0.87675, 0.92434, 0.82917), w.ShadeHit(c, DefaultMaximumReflections))
}

// Avoid infinite recursion (ShadeHit -> ReflectedColor -> ColorAt -> ShadeHit)
// ERRATA page 145 says to use OverPoint in ReflectedColor, which doesn't result
// in infinite recursion. If I use Point though, this test will recurse infinitely.
func TestColorAtWithMutuallyReflectiveSurfaces(t *testing.T) {
	w := DefaultWorld()
	light := NewPointLight(NewPoint(0, 0, 0), NewColor(1, 1, 1))
	w.Lights = []AreaLight{light}

	lower := NewPlane()
	lower.Material.Reflective = 1
	lower.SetTransform(NewTranslation(0, -1, 0))

	upper := NewPlane()
	upper.Material.Reflective = 1
	upper.SetTransform(NewTranslation(0, 1, 0))

	w.Objects = []Shape{lower, upper}

	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 1, 0))

	// The book only wanted us to write this to prove that infinite
	// recursion was possible, but I added this assertion anyway.
	assertEqualColor(t, NewColor(0.2, 0.2, 0.2), w.ColorAt(r, 999999))
}

func TestTheReflectedColorAtTheMaximumRecursiveDepth(t *testing.T) {
	w := DefaultWorld()
	shape := NewPlane()
	shape.Material.Reflective = 0.5
	shape.SetTransform(NewTranslation(0, -1, 0))
	w.Objects = append(w.Objects, shape)
	r := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	i := NewIntersection(math.Sqrt(2), shape)
	c := i.PrepareComputations(r)

	assertEqualColor(t, NewColor(0, 0, 0), w.ReflectedColor(c, 0))
}

func TestFindingN1AndN2AtVariousIntersections(t *testing.T) {
	a := NewGlassSphere()
	a.SetTransform(NewScale(2, 2, 2))
	a.Material.RefractiveIndex = 1.5

	b := NewGlassSphere()
	b.SetTransform(NewTranslation(0, 0, -0.25))
	b.Material.RefractiveIndex = 2.0

	c := NewGlassSphere()
	c.SetTransform(NewTranslation(0, 0, 0.25))
	c.Material.RefractiveIndex = 2.5

	r := NewRay(NewPoint(0, 0, -4), NewVector(0, 0, 1))
	// NB the book uses Intersection[], but I'm making the intersections arg to
	// PrepareComputations() a variadic array of Intersection instead, because
	// it's optional and Go lacks default values.
	xs := Intersections{
		NewIntersection(2, a),
		NewIntersection(2.75, b),
		NewIntersection(3.25, c),
		NewIntersection(4.75, b),
		NewIntersection(5.25, c),
		NewIntersection(6, a),
	}

	var expectedRefractedIndexValues = [][]float64{
		{1.0, 1.5},
		{1.5, 2.0},
		{2.0, 2.5},
		{2.5, 2.5},
		{2.5, 1.5},
		{1.5, 1.0},
	}
	for idx, expected := range expectedRefractedIndexValues {
		comps := xs[idx].PrepareComputations(r, xs...)

		assertEqualFloat64(t, expected[0], comps.N1)
		assertEqualFloat64(t, expected[1], comps.N2)
	}
}

func TestTheUnderPointIsOffsetBelowTheSurface(t *testing.T) {
	shape := NewGlassSphere()
	shape.SetTransform(NewTranslation(0, 0, 1))
	i := NewIntersection(5, shape)
	xs := Intersections{i}
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	comps := i.PrepareComputations(ray, xs...)

	assert(t, comps.UnderPoint.Z > EPSILON/2)
	assert(t, comps.Point.Z < comps.UnderPoint.Z)
}

// ...  "total internal reflection" means all the light is reflected and none is refracted. ...
func TestTheSchlickApproximationUnderTotalInternalReflection(t *testing.T) {
	shape := NewGlassSphere()
	r := NewRay(NewPoint(0, 0, math.Sqrt(2)/2), NewVector(0, 1, 0))
	xs := Intersections{
		NewIntersection(-math.Sqrt(2)/2, shape),
		NewIntersection(math.Sqrt(2)/2, shape),
	}
	comps := xs[1].PrepareComputations(r, xs...)
	reflectance := comps.Schlick()

	assertEqualFloat64(t, 1.0, reflectance)
}

// ... Show that reflectance (via schlick()) is small when a ray strikes the surface at a perpendicular angle. ...
func TestTheSchlickApproximationWithAPerpendicularViewingAngle(t *testing.T) {
	shape := NewGlassSphere()
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 1, 0))
	xs := Intersections{
		NewIntersection(-1, shape),
		NewIntersection(1, shape),
	}
	comps := xs[1].PrepareComputations(r, xs...)
	reflectance := comps.Schlick()

	assertEqualFloat64(t, 0.04, reflectance)
}

// ... Show that reflectance (via schlick()) is significant when n2 > n1 and the ray strikes the surface at a small angle. ...
func TestTheSchlickApproximationWithSmallAngleAndN2GreaterThanN1(t *testing.T) {
	shape := NewGlassSphere()
	r := NewRay(NewPoint(0, 0.99, -2), NewVector(0, 0, 1))
	xs := Intersections{NewIntersection(1.8589, shape)}
	comps := xs[0].PrepareComputations(r, xs...)
	reflectance := comps.Schlick()

	assertEqualFloat64(t, 0.48873, reflectance)
}

func TestAnIntersectionCanEncapsulateUAndV(t *testing.T) {
	s := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))
	i := NewIntersectionWithUV(3.5, s, 0.2, 0.4)

	assertEqualFloat64(t, 0.2, i.U)
	assertEqualFloat64(t, 0.4, i.V)
}

/////////////
// Benchmarks
/////////////

func BenchmarkIntersectionMethodIsEqualTo(b *testing.B) {
	sphere := NewSphere()
	i1 := NewIntersection(1, sphere)
	for i := 0; i < b.N; i++ {
		i1.IsEqualTo(i1)
	}
}

func BenchmarkIntersectionMethodHit(b *testing.B) {
	sphere := NewSphere()
	i1 := NewIntersection(1, sphere)
	i2 := NewIntersection(2, sphere)
	intersections := Intersections{i1, i2}
	for i := 0; i < b.N; i++ {
		intersections.Hit(false)
	}
}

func BenchmarkIntersectionMethodPrepareComputations(b *testing.B) {
	// Taken from TestPrecomputingStateOfIntersection().
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewSphere()
	x := NewIntersection(4, s)
	for i := 0; i < b.N; i++ {
		x.PrepareComputations(r)
	}
}

func BenchmarkIntersectionMethodShlick(b *testing.B) {
	// Taken from TestTheSchlickApproximationUnderTotalInternalReflection().
	shape := NewGlassSphere()
	r := NewRay(NewPoint(0, 0, math.Sqrt(2)/2), NewVector(0, 1, 0))
	xs := Intersections{
		NewIntersection(-math.Sqrt(2)/2, shape),
		NewIntersection(math.Sqrt(2)/2, shape),
	}
	comps := xs[1].PrepareComputations(r, xs...)
	for i := 0; i < b.N; i++ {
		comps.Schlick()
	}
}