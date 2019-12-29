package raytracer

import (
	"fmt"
	"math"
	"testing"
)

func TestCreatingAStripePattern(t *testing.T) {
	pattern := NewStripePattern(Colors["White"], Colors["Black"])

	stripePattern := pattern.LocalPattern.(StripePattern)
	assertEqualColor(t, Colors["White"], stripePattern.A)
	assertEqualColor(t, Colors["Black"], stripePattern.B)
}

func TestAStripePatternIsConstantInY(t *testing.T) {
	pattern := NewStripePattern(Colors["White"], Colors["Black"])

	assertEqualColor(t, Colors["White"], pattern.LocalPattern.LocalPatternAt(NewPoint(0, 0, 0)))
	assertEqualColor(t, Colors["White"], pattern.LocalPattern.LocalPatternAt(NewPoint(0, 1, 0)))
	assertEqualColor(t, Colors["White"], pattern.LocalPattern.LocalPatternAt(NewPoint(0, 2, 0)))
}

func TestAStripePatternIsConstantInZ(t *testing.T) {
	pattern := NewStripePattern(Colors["White"], Colors["Black"])

	assertEqualColor(t, Colors["White"], pattern.LocalPattern.LocalPatternAt(NewPoint(0, 0, 0)))
	assertEqualColor(t, Colors["White"], pattern.LocalPattern.LocalPatternAt(NewPoint(0, 0, 1)))
	assertEqualColor(t, Colors["White"], pattern.LocalPattern.LocalPatternAt(NewPoint(0, 0, 2)))
}

func TestAStripePatternAlternatesInX(t *testing.T) {
	pattern := NewStripePattern(Colors["White"], Colors["Black"])

	assertEqualColor(t, Colors["White"], pattern.LocalPattern.LocalPatternAt(NewPoint(0, 0, 0)))
	assertEqualColor(t, Colors["White"], pattern.LocalPattern.LocalPatternAt(NewPoint(0.9, 0, 0)))
	assertEqualColor(t, Colors["Black"], pattern.LocalPattern.LocalPatternAt(NewPoint(1, 0, 0)))
	assertEqualColor(t, Colors["Black"], pattern.LocalPattern.LocalPatternAt(NewPoint(-0.1, 0, 0)))
	assertEqualColor(t, Colors["Black"], pattern.LocalPattern.LocalPatternAt(NewPoint(-1, 0, 0)))
	assertEqualColor(t, Colors["White"], pattern.LocalPattern.LocalPatternAt(NewPoint(-1.1, 0, 0)))
}

func TestStripesWithAnObjectTransformation(t *testing.T) {
	sphere := NewSphere()
	sphere.SetTransform(NewScale(2, 2, 2))
	sphere.Material.Pattern = NewStripePattern(Colors["White"], Colors["Black"])

	assertEqualColor(t, Colors["White"], sphere.Material.Pattern.PatternAtShape(sphere, NewPoint(1.5, 0, 0)))
}

func TestStripesWithAPatternTransformation(t *testing.T) {
	sphere := NewSphere()
	sphere.Material.Pattern = NewStripePattern(Colors["White"], Colors["Black"])
	sphere.Material.Pattern.SetTransform(NewScale(2, 2, 2))

	assertEqualColor(t, Colors["White"], sphere.Material.Pattern.PatternAtShape(sphere, NewPoint(1.5, 0, 0)))
}

func TestStripesWithBothAnObjectAndAPatternTransformation(t *testing.T) {
	sphere := NewSphere()
	sphere.SetTransform(sphere.Transform.Multiply(NewScale(2, 2, 2)))
	sphere.Material.Pattern = NewStripePattern(Colors["White"], Colors["Black"])
	sphere.Material.Pattern.SetTransform(sphere.Material.Pattern.Transform.Multiply(NewTranslation(0.5, 0, 0)))

	assertEqualColor(t, Colors["White"], sphere.Material.Pattern.PatternAtShape(sphere, NewPoint(2.5, 0, 0)))
}

func TestTheDefaultPatternTransformation(t *testing.T) {
	pattern := NewTestPattern()

	assertEqualMatrix(t, IdentityMatrix(), pattern.Transform)
}

func TestAssigningATransformationToPattern(t *testing.T) {
	pattern := NewTestPattern()
	pattern.SetTransform(NewTranslation(1, 2, 3))

	assertEqualMatrix(t, NewTranslation(1, 2, 3), pattern.Transform)
}

// The following tests replace the ones you wrote earlier in the chapter, testing the stripe pattern’s transformations.
func TestAPatternWithAnObjectTransformation(t *testing.T) {
	shape := NewSphere()
	shape.SetTransform(NewScale(2, 2, 2))
	pattern := NewTestPattern()

	assertEqualColor(t, NewColor(1, 1.5, 2), pattern.PatternAtShape(shape, NewPoint(2, 3, 4)))
}

func TestAPatternWithAPatternTransformation(t *testing.T) {
	shape := NewSphere()
	pattern := NewTestPattern()
	pattern.SetTransform(NewScale(2, 2, 2))

	assertEqualColor(t, NewColor(1, 1.5, 2), pattern.PatternAtShape(shape, NewPoint(2, 3, 4)))
}

func TestAPatternWithBothAnObjectAndAPatternTransformation(t *testing.T) {
	shape := NewSphere()
	shape.SetTransform(NewScale(2, 2, 2))
	pattern := NewTestPattern()
	pattern.SetTransform(NewTranslation(0.5, 1, 1.5))

	assertEqualColor(t, NewColor(0.75, 0.5, 0.25), pattern.PatternAtShape(shape, NewPoint(2.5, 3, 3.5)))
}

func TestAPatternInAGroup(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(NewRotateY(math.Pi / 2))

	g2 := NewGroup()
	g2.SetTransform(NewScale(2, 2, 2))

	shape := NewSphere()
	pattern := NewTestPattern()

	// NB: this is my own test from the Group chapter (p200) -- confirm these numbers are correct.
	assertEqualColor(t, NewColor(2.5, 3, 3.5), pattern.PatternAtShape(shape, NewPoint(2.5, 3, 3.5)))
}

func TestAGradientLinearlyInterpolatesBetweenColors(t *testing.T) {
	p := NewGradientPattern(Colors["White"], Colors["Black"])

	assertEqualColor(t, Colors["White"], p.LocalPattern.LocalPatternAt(NewPoint(0, 0, 0)))
	assertEqualColor(t, NewColor(0.75, 0.75, 0.75), p.LocalPattern.LocalPatternAt(NewPoint(0.25, 0, 0)))
	assertEqualColor(t, NewColor(0.5, 0.5, 0.5), p.LocalPattern.LocalPatternAt(NewPoint(0.5, 0, 0)))
	assertEqualColor(t, NewColor(0.25, 0.25, 0.25), p.LocalPattern.LocalPatternAt(NewPoint(0.75, 0, 0)))
}

func TestARingShouldExtendInBothXAndZ(t *testing.T) {
	p := NewRingPattern(Colors["White"], Colors["Black"])

	assertEqualColor(t, Colors["White"], p.LocalPattern.LocalPatternAt(NewPoint(0, 0, 0)))
	assertEqualColor(t, Colors["Black"], p.LocalPattern.LocalPatternAt(NewPoint(1, 0, 0)))
	assertEqualColor(t, Colors["Black"], p.LocalPattern.LocalPatternAt(NewPoint(0, 0, 1)))
	assertEqualColor(t, Colors["Black"], p.LocalPattern.LocalPatternAt(NewPoint(0.708, 0, 0.709)))
}

func TestCheckersShouldRepeatInX(t *testing.T) {
	p := NewCheckerPattern(Colors["White"], Colors["Black"])

	assertEqualColor(t, Colors["White"], p.LocalPattern.LocalPatternAt(NewPoint(0, 0, 0)))
	assertEqualColor(t, Colors["White"], p.LocalPattern.LocalPatternAt(NewPoint(0.99, 0, 0)))
	assertEqualColor(t, Colors["Black"], p.LocalPattern.LocalPatternAt(NewPoint(1.01, 0, 0)))
}

func TestCheckersShouldRepeatInY(t *testing.T) {
	p := NewCheckerPattern(Colors["White"], Colors["Black"])

	assertEqualColor(t, Colors["White"], p.LocalPattern.LocalPatternAt(NewPoint(0, 0, 0)))
	assertEqualColor(t, Colors["White"], p.LocalPattern.LocalPatternAt(NewPoint(0, 0.99, 0)))
	assertEqualColor(t, Colors["Black"], p.LocalPattern.LocalPatternAt(NewPoint(0, 1.01, 0)))
}

func TestCheckersShouldRepeatInZ(t *testing.T) {
	p := NewCheckerPattern(Colors["White"], Colors["Black"])

	assertEqualColor(t, Colors["White"], p.LocalPattern.LocalPatternAt(NewPoint(0, 0, 0)))
	assertEqualColor(t, Colors["White"], p.LocalPattern.LocalPatternAt(NewPoint(0, 0, 0.99)))
	assertEqualColor(t, Colors["Black"], p.LocalPattern.LocalPatternAt(NewPoint(0, 0, 1.01)))
}

func TestCheckerPatternIn2D(t *testing.T) {
	checkers := NewUVCheckerPattern(2, 2, Colors["Black"], Colors["White"])
	testCases := []struct {
		u        float64
		v        float64
		expected Color
	}{
		{0.0, 0.0, Colors["Black"]},
		{0.5, 0.0, Colors["White"]},
		{0.0, 0.5, Colors["White"]},
		{0.5, 0.5, Colors["Black"]},
		{1.0, 1.0, Colors["Black"]},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("testCases[%d]", idx), func(t *testing.T) {
			assertEqualColor(t, tc.expected, checkers.UVPatternAt(tc.u, tc.v))
		})
	}
}

func TestUsingASphericalMappingOnA3DPoint(t *testing.T) {
	testCases := []struct {
		point Tuple
		u     float64
		v     float64
	}{
		{NewPoint(0, 0, -1), 0.0, 0.5},
		{NewPoint(1, 0, 0), 0.25, 0.5},
		{NewPoint(0, 0, 1), 0.5, 0.5},
		{NewPoint(-1, 0, 0), 0.75, 0.5},
		{NewPoint(0, 1, 0), 0.5, 1.0},
		{NewPoint(0, -1, 0), 0.5, 0.0},
		{NewPoint(math.Sqrt(2)/2, math.Sqrt(2)/2, 0), 0.25, 0.75},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("testCases[%d]", idx), func(t *testing.T) {
			u, v := SphericalMap(tc.point)
			assertEqualFloat64(t, u, tc.u)
			assertEqualFloat64(t, v, tc.v)
		})
	}
}

func TestUsingATextureMapPatternWithASphericalMap(t *testing.T) {
	testCases := []struct {
		point Tuple
		color Color
	}{
		{NewPoint(0.4315, 0.4670, 0.7719), Colors["White"]},
		{NewPoint(-0.9654, 0.2552, -0.0534), Colors["Black"]},
		{NewPoint(0.1039, 0.7090, 0.6975), Colors["White"]},
		{NewPoint(-0.4986, -0.7856, -0.3663), Colors["Black"]},
		{NewPoint(-0.0317, -0.9395, 0.3411), Colors["Black"]},
		{NewPoint(0.4809, -0.7721, 0.4154), Colors["Black"]},
		{NewPoint(0.0285, -0.9612, -0.2745), Colors["Black"]},
		{NewPoint(-0.5734, -0.2162, -0.7903), Colors["White"]},
		{NewPoint(0.7688, -0.1470, 0.6223), Colors["Black"]},
		{NewPoint(-0.7652, 0.2175, 0.6060), Colors["Black"]},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("testCases[%d]", idx), func(t *testing.T) {
			checkers := NewUVCheckerPattern(16, 8, Colors["Black"], Colors["White"])
			pattern := NewTextureMapPattern(checkers, SphericalMap)
			p := pattern.LocalPattern.(TextureMapPattern)
			assertEqualColor(t, tc.color, p.LocalPatternAt(tc.point))
		})
	}
}

/////////////
// Benchmarks
/////////////

func BenchmarkPatternMethodIsEqualTo(b *testing.B) {
	pattern := NewTestPattern()
	for i := 0; i < b.N; i++ {
		pattern.IsEqualTo(pattern)
	}
}

func BenchmarkPatternMethodPatternAtShape(b *testing.B) {
	// Taken from TestStripesWithBothAnObjectAndAPatternTransformation
	sphere := NewSphere()
	sphere.SetTransform(sphere.Transform.Multiply(NewScale(2, 2, 2)))
	sphere.Material.Pattern = NewStripePattern(Colors["White"], Colors["Black"])
	sphere.Material.Pattern.SetTransform(sphere.Material.Pattern.Transform.Multiply(NewTranslation(0.5, 0, 0)))
	for i := 0; i < b.N; i++ {
		sphere.Material.Pattern.PatternAtShape(sphere, NewPoint(2.5, 0, 0))
	}
}
