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
	mat := DefaultMaterial()
	pos := NewPoint(0, 0, 0)

	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), Colors["White"])

	actual := mat.Lighting(light, pos, eyeV, normalV)
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
	mat := DefaultMaterial()
	pos := NewPoint(0, 0, 0)

	eyeV := NewVector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), Colors["White"])

	actual := mat.Lighting(light, pos, eyeV, normalV)
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
	mat := DefaultMaterial()
	pos := NewPoint(0, 0, 0)

	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 10, -10), Colors["White"])

	actual := mat.Lighting(light, pos, eyeV, normalV)
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
	mat := DefaultMaterial()
	pos := NewPoint(0, 0, 0)

	eyeV := NewVector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 10, -10), Colors["White"])

	actual := mat.Lighting(light, pos, eyeV, normalV)
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
	mat := DefaultMaterial()
	pos := NewPoint(0, 0, 0)

	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, 10), Colors["White"])

	actual := mat.Lighting(light, pos, eyeV, normalV)
	expected := NewColor(0.1, 0.1, 0.1)

	assertEqualColor(t, expected, actual)
}
