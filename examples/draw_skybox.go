package examples

import (
	"io/ioutil"
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawSkybox() {
	Draw("tmp/world.jpg", func(world *World, camera *Camera) {
		camera.FieldOfView = math.Pi / 2.5

		camera.SetTransform(NewViewTransform(
			NewPoint(0, 0, 0),
			NewPoint(0, 0, 5),
			NewVector(0, 1, 0),
		))

		getCubeSide := func(filepath string) *Pattern {
			image, err := ioutil.ReadFile(filepath)
			if err != nil {
				panic(err)
			}
			c, err := NewCanvasFromPpm(string(image))
			if err != nil {
				panic(err)
			}
			return NewTextureMapPattern(
				NewUVImagePattern(c),
				PlanarMap,
			)
		}

		cube := NewCube()
		cube.SetTransform(cube.Transform.Compose(
			NewUScale(1000),
		))
		// This cubemap can be found here: http://www.humus.name/index.php?page=Textures&ID=110
		// Fetch the JPGs, save to tmp/, and convert to PPM with `convert XXX.jpg -compress none XXX.ppm`
		cube.Material.Pattern = NewCubeMapPattern(
			getCubeSide("tmp/LancellotiChapel/negx.ppm"), // l
			getCubeSide("tmp/LancellotiChapel/posz.ppm"), // f
			getCubeSide("tmp/LancellotiChapel/posx.ppm"), // r
			getCubeSide("tmp/LancellotiChapel/negz.ppm"), // b
			getCubeSide("tmp/LancellotiChapel/posy.ppm"), // u
			getCubeSide("tmp/LancellotiChapel/negy.ppm"), // d
		)
		cube.Material.Ambient = 1.0
		cube.Material.Specular = 0
		cube.Material.Diffuse = 0

		sphere := NewSphere()
		sphere.SetTransform(sphere.Transform.Compose(
			NewUScale(0.75),
			NewTranslation(0, 0, 5),
		))
		sphere.Material.Diffuse = 0.4
		sphere.Material.Specular = 0.6
		sphere.Material.Shininess = 20
		sphere.Material.Reflective = .6
		sphere.Material.Ambient = 0

		world.Objects = []Shape{
			cube,
			sphere,
		}
		world.Lights = []AreaLight{
			NewPointLight(NewPoint(0, 100, 0), NewColor(1, 1, 1)),
		}
	})
}
