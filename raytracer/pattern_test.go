package raytracer

import (
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
