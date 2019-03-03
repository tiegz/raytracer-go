package raytracer

import (
	"testing"
)

// Scenario: A point light has a position and intensity Given intensity ← color(1, 1, 1)
// And position ← point(0, 0, 0)
// When light ← point_light(position, intensity) Then light.position = position
// And light.intensity = intensity

func TestPointLightHasPositionAndIntensity(t *testing.T) {
	i := Colors["White"]
	p := NewPoint(0, 0, 0)
	pl := NewPointLight(p, i)
	assertEqualTuple(t, p, pl.Position)
	assertEqualColor(t, i, pl.Intensity)
}
