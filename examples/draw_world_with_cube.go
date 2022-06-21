package examples

import (
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithCube(printProgress bool, jobs int) {
	Draw(printProgress, jobs, "tmp/world.jpg", func(world *World, camera *Camera) {
		camera.HSize = 320
		camera.VSize = 200
		camera.FieldOfView = math.Pi / 3

		camera.SetTransform(NewViewTransform(
			NewPoint(0, 3, -10),
			NewPoint(0, 2, 0),
			NewVector(0, 1, 0),
		))

		floor := NewPlane()
		floor.Material.Color = NewColor(1, 0.9, 0.9)
		floor.Material.Specular = NewColor(0, 0, 0)

		cubeLeft := NewCube()
		cubeLeft.SetTransform(NewTranslation(-2, 1, 1))
		cubeLeft.Material.Color = Colors["Red"]
		cubeLeft.Material.Diffuse = NewColor(0.7, 0.7, 0.7)
		cubeLeft.Material.Specular = NewColor(0.3, 0.3, 0.3)

		cubeMiddle := NewCube()
		cubeMiddle.SetTransform(NewTranslation(-0.5, 3, 1))
		cubeMiddle.Material.Color = Colors["Green"]
		cubeMiddle.Material.Diffuse = NewColor(0.7, 0.7, 0.7)
		cubeMiddle.Material.Specular = NewColor(0.3, 0.3, 0.3)

		cubeRight := NewCube()
		cubeRight.Material.Color = Colors["Blue"]
		cubeRight.SetTransform(NewTranslation(1, 1, 1))
		cubeRight.Material.Diffuse = NewColor(0.7, 0.7, 0.7)
		cubeRight.Material.Specular = NewColor(0.3, 0.3, 0.3)

		world.Objects = []*Shape{
			floor,
			cubeLeft,
			cubeMiddle,
			cubeRight,
		}
		world.Lights = []*AreaLight{
			NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
		}
	})
}
