package examples

import (
	"fmt"
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithCubeOfSpheres(jobs int) {
	Draw(jobs, "tmp/world.jpg", func(world *World, camera *Camera) {
		camera.HSize = 200
		camera.VSize = 200
		camera.FieldOfView = math.Pi / 3

		camera.SetTransform(NewViewTransform(
			NewPoint(-5, 15, -5),
			NewPoint(0, 9, 0),
			NewVector(0, 1, 0),
		))

		sphereRowCount := 10.0

		group := NewGroup()
		for x := 0.0; x < sphereRowCount; x++ {
			for y := 0.0; y < sphereRowCount; y++ {
				for z := 0.0; z < sphereRowCount; z++ {
					sphere := NewSphere()
					sphere.SetTransform(NewTranslation(x, y, z))
					sphere.SetTransform(sphere.Transform.Multiply(NewUScale(0.5)))
					sphere.Material.Color = NewColor(x/10.0, y/10.0, z/10.0)
					group.AddChildren(sphere)
				}
			}
		}
		fmt.Printf("Dividing...")
		group.Divide(4)
		fmt.Printf("done dividing!\n")

		world.Objects = []*Shape{
			group,
		}
		world.Lights = []*AreaLight{
			NewPointLight(NewPoint(0, -10, -5), NewColor(1, 1, 1)),
			NewPointLight(NewPoint(0, 10, -5), NewColor(1, 1, 1)),
		}
	})
}
