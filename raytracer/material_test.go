package raytracer

import (
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

	actual := mat.Lighting(obj, light, pos, eyeV, normalV, false)
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

	actual := mat.Lighting(obj, light, pos, eyeV, normalV, false)
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

	actual := mat.Lighting(obj, light, pos, eyeV, normalV, false)
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

	actual := mat.Lighting(obj, light, pos, eyeV, normalV, false)
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

	actual := mat.Lighting(obj, light, pos, eyeV, normalV, false)
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
	inShadow := true

	actual := mat.Lighting(obj, light, pos, eyeV, normalV, inShadow)
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

	c1 := mat.Lighting(obj, light, NewPoint(0.9, 0, 0), eyeV, normalV, false)
	c2 := mat.Lighting(obj, light, NewPoint(1.1, 0, 0), eyeV, normalV, false)

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
		mat.Lighting(obj, light, pos, eyeV, normalV, false)
	}
}
