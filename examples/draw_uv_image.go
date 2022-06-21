package examples

import (
	"io/ioutil"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawUVImage(printProgress bool, jobs int) {
	Draw(printProgress, jobs, "tmp/world.jpg", func(world *World, camera *Camera) {
		camera.SetTransform(NewViewTransform(
			NewPoint(1, 2, -10),
			NewPoint(0, 1.1, 0),
			NewVector(0, 1, 0),
		))

		floor := NewPlane()
		floor.Material.Color = Colors["White"]
		floor.Material.Diffuse = NewColor(0.1, 0.1, 0.1)
		floor.Material.Specular = NewColor(0.0, 0.0, 0.0)
		floor.Material.Ambient = NewColor(0.0, 0.0, 0.0)
		floor.Material.Reflective = 0.4

		platform := NewCylinder()
		platform.LocalShape.(*Cylinder).Minimum = 0
		platform.LocalShape.(*Cylinder).Maximum = 0.1
		platform.LocalShape.(*Cylinder).Closed = true
		platform.Material.Color = Colors["White"]
		platform.Material.Ambient = NewColor(0.0, 0.0, 0.0)
		platform.Material.Diffuse = NewColor(0.2, 0.2, 0.2)
		platform.Material.Specular = NewColor(0.0, 0.0, 0.0)
		platform.Material.Reflective = 0.1

		sphere := NewSphere()
		// Image setup:
		// 1) wget http://planetpixelemporium.com/download/download.php?earthmap1k.jpg
		// 2) convert tmp/earthmap1k.jpg -compress none tmp/earthmap1k.ppm
		image, err := ioutil.ReadFile("raytracer/files/earthmap1k.ppm")
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
		sphere.Material.Diffuse = NewColor(0.9, 0.9, 0.9)
		sphere.Material.Specular = NewColor(0.1, 0.1, 0.1)
		sphere.Material.Shininess = 10
		sphere.Material.Ambient = NewColor(0.1, 0.1, 0.1)

		world.Objects = []*Shape{
			floor,
			platform,
			sphere,
		}
	})
}
