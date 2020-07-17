package examples

import (
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithSphereAndAreaLight() {
	Draw("tmp/world.jpg", func(world *World, camera *Camera) {
		camera.HSize = 320
		camera.VSize = 200
		camera.FieldOfView = math.Pi / 3

		camera.SetTransform(NewViewTransform(
			NewPoint(0, 3, -7),
			NewPoint(0, 1, 0),
			NewVector(0, 1, 0),
		))

		floor := NewPlane()
		floor.Material.Color = NewColor(1, 0.9, 0.9)
		floor.Material.Specular = 0

		sphere := NewSphere()
		sphere.Material.Ambient = 0.1
		sphere.Material.Diffuse = 0.6
		sphere.Material.Specular = 0.0
		sphere.Material.Reflective = 0.3
		sphere.Material.Color = Colors["DarkGreen"]
		sphere.SetTransform(sphere.Transform.Compose(
			NewTranslation(0, 1, 0),
		))

		cube := NewCube()
		cube.Material.Ambient = 0.1
		cube.Material.Diffuse = 0.6
		cube.Material.Specular = 0.0
		cube.Material.Reflective = 0.3
		cube.Material.Color = Colors["DarkRed"]
		cube.SetTransform(cube.Transform.Compose(
			NewTranslation(-4, 1, 5),
		))

		cylinder := NewCylinder()
		cylinder.LocalShape.(*Cylinder).Closed = true
		cylinder.LocalShape.(*Cylinder).Minimum = 0
		cylinder.LocalShape.(*Cylinder).Maximum = 2
		cylinder.Material.Color = Colors["DarkBlue"]
		cylinder.Material.Ambient = 0.6
		cylinder.Material.Diffuse = 0.6
		cylinder.Material.Specular = 0.0
		cylinder.Material.Reflective = 0.3
		cylinder.SetTransform(cylinder.Transform.Compose(
			NewTranslation(4, 0, 5),
		))

		world.Objects = []*Shape{
			floor,
			cube,
			sphere,
			cylinder,
		}
		world.Lights = []AreaLight{
			NewAreaLight(NewPoint(0, 3, -3), NewVector(2, 0, 0), 4, NewVector(0, 2, 0), 4, NewColor(1.5, 1.5, 1.5)),
		}
	})
}
