package examples

import (
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithMultiplePatterns(printProgress bool, jobs int) {
	Draw(printProgress, jobs, "tmp/world.jpg", func(world *World, camera *Camera) {
		camera.HSize = 1920
		camera.VSize = 1080
		camera.FieldOfView = math.Pi / 3

		camera.SetTransform(NewViewTransform(
			NewPoint(0, 1.5, -5),
			NewPoint(0, 1, 0),
			NewVector(0, 1, 0),
		))

		floor := NewPlane()
		// floor.Material.Color = NewColor(1, 0.9, 0.9)
		// floor.Material.Specular = NewColor(0, 0, 0)
		floor.Material.Pattern = NewCheckerPattern(Colors["White"], Colors["Red"])

		midSphere := NewSphere()
		midSphere.SetTransform(NewTranslation(-0.5, 1, 0.5))
		// midSphere.Material.Color = NewColor(0.1, 1, 0.5)
		// midSphere.Material.Diffuse = NewColor(0.7, 0.7, 0.7)
		// midSphere.Material.Specular = NewColor(0.3, 0,3 0.3)
		midSphere.Material.Pattern = NewStripePattern(Colors["Green"], Colors["Purple"])

		rightSphere := NewSphere()
		rightSphere.SetTransform(NewTranslation(1.5, 0.5, -0.5))
		rightSphere.SetTransform(rightSphere.Transform.Multiply(NewScale(0.5, 0.5, 0.5)))
		// rightSphere.Material.Color = NewColor(0.5, 1, 0.1)
		// rightSphere.Material.Diffuse = NewColor(0.7, 0.7, 0.7)
		// rightSphere.Material.Specular = NewColor(0.3, 0.3, 0.3)
		rightSphere.Material.Pattern = NewRingPattern(Colors["Red"], Colors["White"])
		rightSphere.Material.Pattern.SetTransform(NewScale(0.23, 0.23, 0.23))

		leftSphere := NewSphere()
		leftSphere.SetTransform(NewTranslation(-1.5, 0.33, -0.75))
		leftSphere.SetTransform(leftSphere.Transform.Multiply(NewScale(0.33, 0.33, 0.33)))
		leftSphere.Material.Pattern = NewCheckerPattern(Colors["White"], Colors["Black"])
		// leftSphere.Material.Pattern.SetTransform(NewScale(0.01, 0.01, 0.01))
		leftSphere.Material.Color = NewColor(1, 0.8, 0.1)
		leftSphere.Material.Diffuse = NewColor(0.7, 0.7, 0.7)
		leftSphere.Material.Specular = NewColor(0.3, 0.3, 0.3)

		world.Objects = []*Shape{
			floor,
			midSphere,
			leftSphere,
			rightSphere,
		}
		world.Lights = []*AreaLight{
			NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
		}
	})
}
