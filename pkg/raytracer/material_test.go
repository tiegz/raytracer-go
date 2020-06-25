package raytracer

import (
	"fmt"
	"math"
	"testing"
)

func TestDefaultMaterial(t *testing.T) {
	m := DefaultMaterial()

	assertEqualColor(t, Colors["White"], m.Color)
	assertEqualFloat64(t, 0.1, m.Ambient)
	assertEqualFloat64(t, 0.9, m.Diffuse)
	assertEqualFloat64(t, 0.9, m.Specular)
	assertEqualFloat64(t, 200, m.Shininess)
}

//  			 			|
//  				 		|
// 	 				 	  |
// 💡-----👁 <---|
//  			 			|
//  				 		|
// 	 				 	  |
func TestLightingWithEyeBetweenLightAndSurface(t *testing.T) {
	obj := NewSphere()
	mat := DefaultMaterial()
	pos := NewPoint(0, 0, 0)

	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), Colors["White"])

	actual := mat.Lighting(obj, light, pos, eyeV, normalV, 1.0)
	expected := NewColor(1.9, 1.9, 1.9)

	assertEqualColor(t, expected, actual)
}

//  			👁 	|
//  				\ |
// 	 				 \|
// 💡------------|
//  			 		|
//  				 	|
// 	 				 	|

func TestLightingWithEyeBetweenLightAndSurfaceAndEyeOffset45Degrees(t *testing.T) {
	obj := NewSphere()
	mat := DefaultMaterial()
	pos := NewPoint(0, 0, 0)

	eyeV := NewVector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), Colors["White"])

	actual := mat.Lighting(obj, light, pos, eyeV, normalV, 1.0)
	expected := NewColor(1, 1, 1)

	assertEqualColor(t, expected, actual)
}

//  			💡 	|
//  				\ |
// 	 				 \|
// 👁---------|
//  			 		|
//  				 	|
// 	 				 	|

func TestLightingWithEyeOppositeSurfaceAndLightOffset45Degrees(t *testing.T) {
	obj := NewSphere()
	mat := DefaultMaterial()
	pos := NewPoint(0, 0, 0)

	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 10, -10), Colors["White"])

	actual := mat.Lighting(obj, light, pos, eyeV, normalV, 1.0)
	expected := NewColor(0.7364, 0.7364, 0.7364)

	assertEqualColor(t, expected, actual)
}

//  			💡 	|
//  				\ |
// 	 				 \|
//            |
//  			 	 /|
//  				/	|
// 	 			👁	|
func TestLightingWithEyeInPathOfReflectionVector(t *testing.T) {
	obj := NewSphere()
	mat := DefaultMaterial()
	pos := NewPoint(0, 0, 0)

	eyeV := NewVector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 10, -10), Colors["White"])

	actual := mat.Lighting(obj, light, pos, eyeV, normalV, 1.0)
	expected := NewColor(1.6364, 1.6364, 1.6364)

	assertEqualColor(t, expected, actual)
}

//  			 			|
//  				 		|
// 	 				 	  |
// 👁 <---------| ------->💡
//  			 			|
//  				 		|
// 	 				 	  |
func TestLightingWithLightBehindSurface(t *testing.T) {
	obj := NewSphere()
	mat := DefaultMaterial()
	pos := NewPoint(0, 0, 0)

	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, 10), Colors["White"])

	actual := mat.Lighting(obj, light, pos, eyeV, normalV, 1.0)
	expected := NewColor(0.1, 0.1, 0.1)

	assertEqualColor(t, expected, actual)
}

//  			 			 |
//  				 		 |
// 	 				 	   |
// 💡-----👁 <---|
//  			 			 |
//  				 		 |
// 	 				 	   |
func TestLightingWithTheSurfaceInShadow(t *testing.T) {
	obj := NewSphere()
	mat := DefaultMaterial()
	pos := NewPoint(0, 0, 0)

	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), Colors["White"])
	actual := mat.Lighting(obj, light, pos, eyeV, normalV, 0.0) // in shadow
	expected := NewColor(0.1, 0.1, 0.1)

	assertEqualColor(t, expected, actual)
}

func TestLightingWithAPatternApplied(t *testing.T) {
	obj := NewSphere()
	mat := DefaultMaterial()
	mat.Pattern = NewStripePattern(Colors["White"], Colors["Black"])
	mat.Ambient = 1
	mat.Diffuse = 0
	mat.Specular = 0
	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), Colors["White"])

	c1 := mat.Lighting(obj, light, NewPoint(0.9, 0, 0), eyeV, normalV, 1.0)
	c2 := mat.Lighting(obj, light, NewPoint(1.1, 0, 0), eyeV, normalV, 1.0)

	assertEqualColor(t, Colors["White"], c1)
	assertEqualColor(t, Colors["Black"], c2)
}

func TestReflectivityForTheDefaultMaterial(t *testing.T) {
	m := DefaultMaterial()

	assertEqualFloat64(t, 0.0, m.Reflective)
}

func TestTransparencyAndRefractiveIndexForTheDefaultMaterial(t *testing.T) {
	m := DefaultMaterial()

	assertEqualFloat64(t, 0.0, m.Transparency)
	assertEqualFloat64(t, 1.0, m.RefractiveIndex)
}

func TestLightingUsesLightIntensityToAttentuateColor(t *testing.T) {
	w := DefaultWorld()
	w.Lights[0] = NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	s := w.Objects[0]
	s.Material.Ambient = 0.1
	s.Material.Diffuse = 0.9
	s.Material.Specular = 0
	s.Material.Color = NewColor(1, 1, 1)
	pt := NewPoint(0, 0, -1)
	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)

	testCases := []struct {
		intensity float64
		result    Color
	}{
		{1.0, NewColor(1, 1, 1)},
		{0.5, NewColor(0.55, 0.55, 0.55)},
		{0.0, NewColor(0.1, 0.1, 0.1)},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("testCases[#%d]", idx), func(t *testing.T) {
			result := s.Material.Lighting(s, w.Lights[0], pt, eyeV, normalV, tc.intensity)
			assertEqualColor(t, tc.result, result)
		})
	}
}

func TestLightingSamplesTheAreaLight(t *testing.T) {
	light := NewAreaLight(NewPoint(-0.5, -0.5, -5), NewVector(1, 0, 0), 2, NewVector(0, 1, 0), 2, NewColor(1, 1, 1))
	shape := NewSphere()
	shape.Material.Ambient = 0.1
	shape.Material.Diffuse = 0.9
	shape.Material.Specular = 0
	shape.Material.Color = NewColor(1, 1, 1)
	eye := NewPoint(0, 0, -5)
	testCases := []struct {
		point  Tuple
		result Color
	}{
		{NewPoint(0, 0, -1), NewColor(0.9965, 0.9965, 0.9965)},
		{NewPoint(0, 0.7071, -0.7071), NewColor(0.62318, 0.62318, 0.62318)}, // HACK: this is a slightly different value than the bonus chapter's: NewColor(0.6232, 0.6232, 0.6232)

	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("testCases[#%d]", idx), func(t *testing.T) {
			eyeV := eye.Subtract(tc.point)
			normalV := NewVector(tc.point.X, tc.point.Y, tc.point.Z)
			result := shape.Material.Lighting(shape, light, tc.point, eyeV, normalV, 1.0)
			assertEqualColor(t, tc.result, result)
		})
	}

}

/////////////
// Benchmarks
/////////////

func BenchmarkMaterialMethodIsEqualTo(b *testing.B) {
	mat := DefaultMaterial()
	for i := 0; i < b.N; i++ {
		mat.IsEqualTo(mat)
	}
}

func BenchmarkMaterialMethodLighting(b *testing.B) {
	// Taken from TestLightingWithEyeBetweenLightAndSurface()
	obj := NewSphere()
	mat := DefaultMaterial()
	pos := NewPoint(0, 0, 0)
	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), Colors["White"])
	for i := 0; i < b.N; i++ {
		mat.Lighting(obj, light, pos, eyeV, normalV, 1.0)
	}
}
