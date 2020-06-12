package raytracer

import (
	"fmt"
	"math"
)

type Material struct {
	Label           string
	Color           Color
	Ambient         float64
	Diffuse         float64
	Specular        float64
	Shininess       float64
	Pattern         Pattern
	Reflective      float64
	Transparency    float64
	RefractiveIndex float64 // Vacuum=1, Water=1.333, Glass=1.52, Diamond=2.42
}

// Beware: use this instead of Material{}, for Material{} without all the args will throw errors when rendering.
func DefaultMaterial() Material {
	return Material{
		Label:           "default-material",
		Color:           Colors["White"],
		Ambient:         0.1,
		Diffuse:         0.9,
		Specular:        0.9,
		Shininess:       200,
		Pattern:         NewNullPattern(),
		Reflective:      0.0,
		Transparency:    0.0,
		RefractiveIndex: 1.0,
	}
}

func (m Material) IsEqualTo(m2 Material) bool {
	// TODO add check for Pattern equality too
	if !m.Color.IsEqualTo(m2.Color) {
		return false
	} else if m.Ambient != m2.Ambient {
		return false
	} else if m.Diffuse != m2.Diffuse {
		return false
	} else if m.Specular != m2.Specular {
		return false
	} else if m.Shininess != m2.Shininess {
		return false
	} else if m.Reflective != m2.Reflective {
		return false
	} else if m.Transparency != m2.Transparency {
		return false
	} else if m.RefractiveIndex != m2.RefractiveIndex {
		return false
	}
	return true
}

func (m Material) String() string {
	return fmt.Sprintf(
		"Material(\n  Label: %v\n  Color: %v\n  Ambient: %v\n  Diffuse: %v\n  Specular: %v\n  Shininess: %v\n  Pattern: %v\n  Reflective: %v\n  Transparency: %v\n  ReflectiveIndex: %v\n)",
		m.Label,
		m.Color,
		m.Ambient,
		m.Diffuse,
		m.Specular,
		m.Shininess,
		m.Pattern,
		m.Reflective,
		m.Transparency,
		m.RefractiveIndex,
	)
}

// Calculates the lighting for a given point and material, based on Phong reflection.
// Phone reflection model:
//   * Ambient reflection:  background lighting; a constant value.
//   * Diffuse reflection:  reflection from matte surface; depends on angle btwn light and surface.
//   * Specular reflection: reflection of the light source; depends on angle btwn the reflection
//      										and eye vectors. Intensity is controlled by "shininess".
// Inensity: 0.0 = in shadow, 1.0 = not in shadow.
func (m Material) Lighting(obj Shape, light AreaLight, point Tuple, eyeVector, normalVector Tuple, intensity float64) Color {
	var baseColor, ambient, specular, diffuse Color

	if !m.Pattern.IsEqualTo(NewNullPattern()) {
		baseColor = m.Pattern.PatternAtShape(obj, point)
	} else {
		baseColor = m.Color
	}

	effectiveColor := baseColor.MultiplyColor(light.GetIntensity()) // Combine the surface color with the light's color/intensity
	ambient = effectiveColor.Multiply(m.Ambient)                    // Compute the ambient contribution

	samples := []Tuple{}
	// TODO: can we memoize and abstract this out, with the shared code in IntensityAt()?
	for v := 0.0; v < light.VSteps; v++ {
		for u := 0.0; u < light.USteps; u++ {
			samples = append(samples, light.PointOnLight(u, v))
		}
	}

	sum := Colors["Black"]

	// Loop through each of the points in the area light
	for _, sample := range samples {
		// The direction to the light source
		lightVector := sample.Subtract(point).Normalized()
		// The cosine of the angle between the light vector and the normal vector. Negative means light is on other side of surface.
		lightDotNormal := lightVector.Dot(normalVector)
		if lightDotNormal < EPSILON || intensity < EPSILON {
			// When inside a shadow, you only need ambient, not diffuse & specular.
			continue
		}

		// Compute the diffuse contribution
		diffuse = effectiveColor.Multiply(m.Diffuse).Multiply(lightDotNormal)
		sum = sum.Add(diffuse)

		// Compute the specular contribution
		reflectVector := lightVector.Negate().Reflect(normalVector)
		reflectDotEye := reflectVector.Dot(eyeVector) // The cosine of angle between reflection vector + eye vector (nenative means light reflects away from eye)
		if reflectDotEye > 0 {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = light.GetIntensity().Multiply(m.Specular).Multiply(factor)
			sum = sum.Add(specular)
		}
	}

	return ambient.Add(sum.Divide(light.Samples).Multiply(intensity))
}
