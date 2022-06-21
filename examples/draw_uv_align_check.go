package examples

import (
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawUVAlignCheck(printProgress bool, jobs int) {
	Draw(printProgress, jobs, "tmp/world.jpg", func(world *World, camera *Camera) {
		camera.HSize = 400
		camera.VSize = 400
		camera.FieldOfView = math.Pi / 3

		camera.SetTransform(NewViewTransform(
			NewPoint(1, 2, -5),
			NewPoint(0, 0, 0),
			NewVector(0, 1, 0),
		))

		floor := NewPlane()
		floor.Material.Ambient = NewColor(0.1, 0.1, 0.1)
		floor.Material.Diffuse = NewColor(0.8, 0.8, 0.8)
		floor.Material.Pattern = NewTextureMapPattern(
			NewUVAlignCheckPattern(
				NewColor(1, 1, 1), // white
				NewColor(1, 0, 0), // red
				NewColor(1, 1, 0), // yellow
				NewColor(0, 1, 0), // green
				NewColor(0, 1, 1), // cyan
			),
			PlanarMap,
		)

		world.Objects = []*Shape{floor}
		world.Lights = []*AreaLight{
			NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
		}
	})
}
