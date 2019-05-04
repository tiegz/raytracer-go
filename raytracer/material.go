package raytracer

import (
	"fmt"
	"math"
)

type Material struct {
	Color      Color
	Ambient    float64
	Diffuse    float64
	Specular   float64
	Shininess  float64
	Pattern    Pattern
	Reflective float64
}

func DefaultMaterial() Material {
	return Material{
		Color:      Colors["White"],
		Ambient:    0.1,
		Diffuse:    0.9,
		Specular:   0.9,
		Shininess:  200,
		Pattern:    NewNullPattern(),
		Reflective: 0.0,
	}
}

func (m *Material) IsEqualTo(m2 Material) bool {
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
	}
	return true
}

func (m Material) String() string {
	return fmt.Sprintf("Material( %v %v %v %v %v %v %v\n)", m.Color, m.Ambient, m.Diffuse, m.Specular, m.Shininess, m.Pattern, m.Reflective)
}

// Calculates the lighting for a given point and material, based on Phong reflection.
// Phone reflection model:
//   * Ambient reflection:  background lighting; a constant value.
//   * Diffuse reflection:  reflection from matte surface; depends on angle btwn light and surface.
//   * Specular reflection: reflection of the light source; depends on angle btwn the reflection
//      										and eye vectors. Intensity is controlled by "shininess".
func (m *Material) Lighting(obj Shape, light PointLight, point Tuple, eyeVector, normalVector Tuple, inShadow bool) Color {
	var baseColor, ambient, specular, diffuse Color

	if !m.Pattern.IsEqualTo(NewNullPattern()) {
		baseColor = m.Pattern.PatternAtShape(obj, point)
	} else {
		baseColor = m.Color
	}

	// combine the surface color with the light's color/intensity
	effectiveColor := baseColor.MultiplyColor(light.Intensity)

	// find the direction to the light source
	lightVector := light.Position.Subtract(point)
	lightVector = lightVector.Normalized()

	// compute the ambient contribution
	ambient = effectiveColor.Multiply(m.Ambient)

	if inShadow {
		// when in a shadow, you only need ambient, not duffse & specular.
		return ambient
	}

	// light_dot_normal represents the cosine of the angle between the
	// light vector and the normal vector. A negative number means the
	// light is on the other side of the surface.
	lightDotNormal := lightVector.Dot(normalVector)
	if lightDotNormal < 0 {
		diffuse = Colors["Black"]
		specular = Colors["Black"]
	} else {
		// compute the diffuse contribution
		diffuse = effectiveColor.Multiply(m.Diffuse)
		diffuse = diffuse.Multiply(lightDotNormal)
		// reflect_dot_eye represents the cosine of the angle between the
		// reflection vector and the eye vector. A negative number means the
		// light reflects away from the eye.
		reflectVector := lightVector.Negate()
		reflectVector = reflectVector.Reflect(normalVector)
		reflectDotEye := reflectVector.Dot(eyeVector)

		if reflectDotEye <= 0 {
			specular = Colors["Black"]
		} else {
			// compute the specular contribution
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = light.Intensity.Multiply(m.Specular)
			specular = specular.Multiply(factor)
		}
	}

	// add the three contributions together to get the final shading
	shading := ambient.Add(diffuse)
	shading = shading.Add(specular)

	return shading
}
