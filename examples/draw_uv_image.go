package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	. "github.com/tiegz/raytracer-go/pkg/raytracer"
)

func RunDrawUVImage() {
	camera := NewCamera(800, 400, math.Pi/3)

	camera.SetTransform(NewViewTransform(
		NewPoint(1, 2, -10),
		NewPoint(0, 1.1, 0),
		NewVector(0, 1, 0),
	))

	floor := NewPlane()
	floor.Material.Color = Colors["White"]
	floor.Material.Diffuse = 0.1
	floor.Material.Specular = 0.0
	floor.Material.Ambient = 0.0
	floor.Material.Reflective = 0.4

	platform := NewCylinder()
	platform.LocalShape.(*Cylinder).Minimum = 0
	platform.LocalShape.(*Cylinder).Maximum = 0.1
	platform.LocalShape.(*Cylinder).Closed = true
	platform.Material.Color = Colors["White"]
	platform.Material.Ambient = 0.0
	platform.Material.Diffuse = 0.2
	platform.Material.Specular = 0.0
	platform.Material.Reflective = 0.1

	sphere := NewSphere()
	// This cubemap can be found at http://planetpixelemporium.com/earth.html > cube map
	// Fetch the files, save to tmp, and convert to PPM: convert earthmap1k.jpg -compress none earthmap1k.ppm
	image, err := ioutil.ReadFile("tmp/earthmap1k.ppm")
	if err != nil {
		panic(err)
	}

	c, err := NewCanvasFromPpm(string(image))
	if err != nil {
		panic(err)
	}
	sphere.Material.Pattern = NewTextureMapPattern(
		NewUVImagePattern(c),
		SphericalMap,
	)
	sphere.SetTransform(sphere.Transform.Compose(
		NewRotateY(1.9),
		NewTranslation(0, 1.1, 0),
	))
	sphere.Material.Diffuse = 0.9
	sphere.Material.Specular = 0.1
	sphere.Material.Shininess = 10
	sphere.Material.Ambient = 0.1

	world := NewWorld()
	world.Objects = []Shape{
		floor,
		platform,
		sphere,
	}
	world.Lights = []AreaLight{
		NewPointLight(NewPoint(-100, 100, -100), NewColor(1, 1, 1)),
	}

	canvas := camera.RenderWithProgress(world)

	fmt.Println("Generating PPM...")
	ppm := canvas.ToPpm()
	filename := "tmp/world.ppm"
	ppmBytes := []byte(ppm)
	fmt.Printf("Saving scene to %s...\n", filename)
	if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
		panic(err)
	}
}
