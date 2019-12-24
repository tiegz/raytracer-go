package raytracer

import (
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
	sphere.Transform = NewScale(2, 2, 2)
	sphere.Material.Pattern = NewStripePattern(Colors["White"], Colors["Black"])

	assertEqualColor(t, Colors["White"], sphere.Material.Pattern.PatternAtShape(sphere, NewPoint(1.5, 0, 0)))
}

func TestStripesWithAPatternTransformation(t *testing.T) {
	sphere := NewSphere()
	sphere.Material.Pattern = NewStripePattern(Colors["White"], Colors["Black"])
	sphere.Material.Pattern.Transform = NewScale(2, 2, 2)

	assertEqualColor(t, Colors["White"], sphere.Material.Pattern.PatternAtShape(sphere, NewPoint(1.5, 0, 0)))
}

func TestStripesWithBothAnObjectAndAPatternTransformation(t *testing.T) {
	sphere := NewSphere()
	sphere.Transform = sphere.Transform.Multiply(NewScale(2, 2, 2))
	sphere.Material.Pattern = NewStripePattern(Colors["White"], Colors["Black"])
	sphere.Material.Pattern.Transform = sphere.Material.Pattern.Transform.Multiply(NewTranslation(0.5, 0, 0))

	assertEqualColor(t, Colors["White"], sphere.Material.Pattern.PatternAtShape(sphere, NewPoint(2.5, 0, 0)))
}

func TestTheDefaultPatternTransformation(t *testing.T) {
	pattern := NewTestPattern()

	assertEqualMatrix(t, IdentityMatrix(), pattern.Transform)
}

func TestAssigningATransformationToPattern(t *testing.T) {
	pattern := NewTestPattern()
	pattern.Transform = NewTranslation(1, 2, 3)

	assertEqualMatrix(t, NewTranslation(1, 2, 3), pattern.Transform)
}

// The following tests replace the ones you wrote earlier in the chapter, testing the stripe pattern’s transformations.
func TestAPatternWithAnObjectTransformation(t *testing.T) {
	shape := NewSphere()
	shape.Transform = NewScale(2, 2, 2)
	pattern := NewTestPattern()

	assertEqualColor(t, NewColor(1, 1.5, 2), pattern.PatternAtShape(shape, NewPoint(2, 3, 4)))
}

func TestAPatternWithAPatternTransformation(t *testing.T) {
	shape := NewSphere()
	pattern := NewTestPattern()
	pattern.Transform = NewScale(2, 2, 2)

	assertEqualColor(t, NewColor(1, 1.5, 2), pattern.PatternAtShape(shape, NewPoint(2, 3, 4)))
}

func TestAPatternWithBothAnObjectAndAPatternTransformation(t *testing.T) {
	shape := NewSphere()
	shape.Transform = NewScale(2, 2, 2)
	pattern := NewTestPattern()
	pattern.Transform = NewTranslation(0.5, 1, 1.5)

	assertEqualColor(t, NewColor(0.75, 0.5, 0.25), pattern.PatternAtShape(shape, NewPoint(2.5, 3, 3.5)))
}

func TestAPatternInAGroup(t *testing.T) {
	g1 := NewGroup()
	g1.Transform = NewRotateY(math.Pi / 2)

	g2 := NewGroup()
	g2.Transform = NewScale(2, 2, 2)

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
	sphere.Transform = sphere.Transform.Multiply(NewScale(2, 2, 2))
	sphere.Material.Pattern = NewStripePattern(Colors["White"], Colors["Black"])
	sphere.Material.Pattern.Transform = sphere.Material.Pattern.Transform.Multiply(NewTranslation(0.5, 0, 0))
	for i := 0; i < b.N; i++ {
		sphere.Material.Pattern.PatternAtShape(sphere, NewPoint(2.5, 0, 0))
	}
}
