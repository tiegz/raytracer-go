package raytracer

import (
	"testing"
)

func TestCreatingAStripePattern(t *testing.T) {
	pattern := NewStripePattern(Colors["White"], Colors["Black"])

	assertEqualColor(t, Colors["White"], pattern.A)
	assertEqualColor(t, Colors["Black"], pattern.B)
}

func TestAStripePatternIsConstantInY(t *testing.T) {
	pattern := NewStripePattern(Colors["White"], Colors["Black"])

	assertEqualColor(t, Colors["White"], pattern.StripeAt(NewPoint(0, 0, 0)))
	assertEqualColor(t, Colors["White"], pattern.StripeAt(NewPoint(0, 1, 0)))
	assertEqualColor(t, Colors["White"], pattern.StripeAt(NewPoint(0, 2, 0)))
}

func TestAStripePatternIsConstantInZ(t *testing.T) {
	pattern := NewStripePattern(Colors["White"], Colors["Black"])

	assertEqualColor(t, Colors["White"], pattern.StripeAt(NewPoint(0, 0, 0)))
	assertEqualColor(t, Colors["White"], pattern.StripeAt(NewPoint(0, 0, 1)))
	assertEqualColor(t, Colors["White"], pattern.StripeAt(NewPoint(0, 0, 2)))
}

func TestAStripePatternAlternatesInX(t *testing.T) {
	pattern := NewStripePattern(Colors["White"], Colors["Black"])

	assertEqualColor(t, Colors["White"], pattern.StripeAt(NewPoint(0, 0, 0)))
	assertEqualColor(t, Colors["White"], pattern.StripeAt(NewPoint(0.9, 0, 0)))
	assertEqualColor(t, Colors["Black"], pattern.StripeAt(NewPoint(1, 0, 0)))
	assertEqualColor(t, Colors["Black"], pattern.StripeAt(NewPoint(-0.1, 0, 0)))
	assertEqualColor(t, Colors["Black"], pattern.StripeAt(NewPoint(-1, 0, 0)))
	assertEqualColor(t, Colors["White"], pattern.StripeAt(NewPoint(-1.1, 0, 0)))
}

func TestStripesWithAnObjectTransformation(t *testing.T) {
	sphere := NewSphere()
	sphere.Transform = NewScale(2, 2, 2)
	sphere.Material.Pattern = NewStripePattern(Colors["White"], Colors["Black"])

	assertEqualColor(t, Colors["White"], sphere.Material.Pattern.StripeAtObject(sphere, NewPoint(1.5, 0, 0)))
}

func TestStripesWithAPatternTransformation(t *testing.T) {
	sphere := NewSphere()
	sphere.Material.Pattern = NewStripePattern(Colors["White"], Colors["Black"])
	sphere.Material.Pattern.Transform = NewScale(2, 2, 2)

	assertEqualColor(t, Colors["White"], sphere.Material.Pattern.StripeAtObject(sphere, NewPoint(1.5, 0, 0)))
}

func TestStripesWithBothAnObjectAndAPatternTransformation(t *testing.T) {
	sphere := NewSphere()
	sphere.Transform = sphere.Transform.Multiply(NewScale(2, 2, 2))
	sphere.Material.Pattern = NewStripePattern(Colors["White"], Colors["Black"])
	sphere.Material.Pattern.Transform = sphere.Material.Pattern.Transform.Multiply(NewTranslation(0.5, 0, 0))

	assertEqualColor(t, Colors["White"], sphere.Material.Pattern.StripeAtObject(sphere, NewPoint(2.5, 0, 0)))
}
