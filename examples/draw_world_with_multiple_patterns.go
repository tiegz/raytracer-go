package examples

import (
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithMultiplePatterns() {
	Draw("tmp/world.jpg", func(world *World, camera *Camera) {
		camera.SetSize(1920, 1080, math.Pi/3)
		camera.SetTransform(NewViewTransform(
			NewPoint(0, 1.5, -5),
			NewPoint(0, 1, 0),
			NewVector(0, 1, 0),
		))

		floor := NewPlane()
		// floor.Material.Color = NewColor(1, 0.9, 0.9)
		// floor.Material.Specular = 0
		floor.Material.Pattern = NewCheckerPattern(Colors["White"], Colors["Red"])

		midSphere := NewSphere()
		midSphere.SetTransform(NewTranslation(-0.5, 1, 0.5))
		// midSphere.Material.Color = NewColor(0.1, 1, 0.5)
		// midSphere.Material.Diffuse = 0.7
		// midSphere.Material.Specular = 0.3
		midSphere.Material.Pattern = NewStripePattern(Colors["Green"], Colors["Purple"])

		rightSphere := NewSphere()
		rightSphere.SetTransform(NewTranslation(1.5, 0.5, -0.5))
		rightSphere.SetTransform(rightSphere.Transform.Multiply(NewScale(0.5, 0.5, 0.5)))
		// rightSphere.Material.Color = NewColor(0.5, 1, 0.1)
		// rightSphere.Material.Diffuse = 0.7
		// rightSphere.Material.Specular = 0.3
		rightSphere.Material.Pattern = NewRingPattern(Colors["Red"], Colors["White"])
		rightSphere.Material.Pattern.SetTransform(NewScale(0.23, 0.23, 0.23))

		leftSphere := NewSphere()
		leftSphere.SetTransform(NewTranslation(-1.5, 0.33, -0.75))
		leftSphere.SetTransform(leftSphere.Transform.Multiply(NewScale(0.33, 0.33, 0.33)))
		leftSphere.Material.Pattern = NewCheckerPattern(Colors["White"], Colors["Black"])
		// leftSphere.Material.Pattern.SetTransform(NewScale(0.01, 0.01, 0.01))
		leftSphere.Material.Color = NewColor(1, 0.8, 0.1)
		leftSphere.Material.Diffuse = 0.7
		leftSphere.Material.Specular = 0.3

		world.Objects = []Shape{
			floor,
			midSphere,
			leftSphere,
			rightSphere,
		}
		world.Lights = []AreaLight{
			NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
		}
	})
}
