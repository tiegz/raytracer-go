package examples

import (
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithTriangles(printProgress bool, jobs int) {
	Draw(printProgress, jobs, "tmp/world.jpg", func(world *World, camera *Camera) {
		camera.HSize = 320
		camera.VSize = 200
		camera.FieldOfView = math.Pi / 3

		camera.SetTransform(NewViewTransform(
			NewPoint(-3, 2, 3),
			NewPoint(0, 0, 0),
			NewVector(0, 1, 0),
		))

		sphere := NewSphere()

		// pyramid1 := pyramid()
		// pyramid2 := pyramid()
		// pyramid2.SetTransform(pyramid2.Transform.Compose(
		// 	NewTranslation(-1.5, 0, -1.5),
		// 	NewRotateY(math.Pi),
		// ))
		// pyramid3 := pyramid()
		// pyramid3.SetTransform(pyramid3.Transform.Compose(
		// 	NewTranslation(2, 0, 0),
		// 	NewRotateY(math.Pi/2),
		// ))

		world.Objects = []*Shape{
			sphere,
		}
		world.Lights = []*AreaLight{
			NewPointLight(NewPoint(0, 5, 0), NewColor(1, 1, 1)),
		}
	})
}
