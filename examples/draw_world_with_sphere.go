package examples

import (
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithSphere(printProgress bool, jobs int) {
	Draw(printProgress, jobs, "tmp/world.jpg", func(world *World, camera *Camera) {
		camera.HSize = 1280
		camera.VSize = 120
		camera.FieldOfView = math.Pi / 3

		camera.SetTransform(NewViewTransform(
			NewPoint(0, 5, -26),
			NewPoint(0, 1, 0),
			NewVector(0, 1, 0),
		))

		floor := NewPlane()
		floor.Material.Color = NewColor(1, 0.9, 0.9)
		floor.Material.Specular = 0
		floor.Material.Pattern = NewCheckerPattern(Colors["Gray"], Colors["White"])

		world.Objects = []*Shape{
			floor,
		}

		for i := 0; i < 10; i++ {
			sphere := NewSphere()
			sphere.SetTransform(sphere.Transform.Compose(
				NewTranslation((float64(i) * 3.0) - 13.5, 1, 0),
			))
			// sphere.Material.Ambient = 0.1 * (float64(i) + 1);
			// sphere.Material.Diffuse = 0.1 * (float64(i) + 1);
			sphere.Material.Specular = 0.1 * (float64(i) + 1);
			// sphere.Material.Shininess = (100 * float64(i)) + 1.0;
			// sphere.Material.Reflective = 0.1 * (float64(i) + 1);
			// sphere.Material.Transparency = 0.1 * (float64(i) + 1);
			// sphere.Material.RefractiveIndex = 0.5
			world.Objects = append(world.Objects, sphere)
		}
		// sphere.Material.Ambient = 0.1
		// sphere.Material.Diffuse = 0.6
		// sphere.Material.Specular = 0.0
		// sphere.Material.Reflective = 0.3
		// sphere.Material.Color = Colors["DarkGreen"]

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

		world.Lights = []*AreaLight{
			NewPointLight(NewPoint(0, 5, -1), NewColor(1, 1, 1)),
		}
	})
}
