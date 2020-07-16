package examples

import (
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorld() {
	Draw("tmp/world.jpg", func(world *World, camera *Camera) {
		camera.SetSize(320, 240, math.Pi/3)
		camera.SetTransform(NewViewTransform(
			NewPoint(0, 1.5, -5),
			NewPoint(0, 1, 0),
			NewVector(0, 1, 0),
		))

		floor := NewSphere()
		floor.SetTransform(NewScale(10, 0.01, 10))
		floor.Material.Color = NewColor(1, 0.9, 0.9)
		floor.Material.Specular = 0

		leftWall := NewSphere()
		leftWall.SetTransform(NewTranslation(0, 0, 5))
		leftWall.SetTransform(leftWall.Transform.Multiply(NewRotateY(-math.Pi / 4)))
		leftWall.SetTransform(leftWall.Transform.Multiply(NewRotateX(math.Pi / 2)))
		leftWall.SetTransform(leftWall.Transform.Multiply(NewScale(10, 0.01, 10)))
		leftWall.Material = floor.Material

		rightWall := NewSphere()
		rightWall.SetTransform(NewTranslation(0, 0, 5))
		rightWall.SetTransform(rightWall.Transform.Multiply(NewRotateY(math.Pi / 4)))
		rightWall.SetTransform(rightWall.Transform.Multiply(NewRotateX(math.Pi / 2)))
		rightWall.SetTransform(rightWall.Transform.Multiply(NewScale(10, 0.01, 10)))
		rightWall.Material = floor.Material

		midSphere := NewSphere()
		midSphere.SetTransform(NewTranslation(-0.5, 1, 0.5))
		midSphere.Material.Color = NewColor(0.1, 1, 0.5)
		midSphere.Material.Diffuse = 0.7
		midSphere.Material.Specular = 0.3

		rightSphere := NewSphere()
		rightSphere.SetTransform(NewTranslation(1.5, 0.5, -0.5))
		rightSphere.SetTransform(rightSphere.Transform.Multiply(NewScale(0.5, 0.5, 0.5)))
		rightSphere.Material.Color = NewColor(0.5, 1, 0.1)
		rightSphere.Material.Diffuse = 0.7
		rightSphere.Material.Specular = 0.3

		leftSphere := NewSphere()
		leftSphere.SetTransform(NewTranslation(-1.5, 0.33, -0.75))
		leftSphere.SetTransform(leftSphere.Transform.Multiply(NewScale(0.33, 0.33, 0.33)))
		leftSphere.Material.Color = NewColor(1, 0.8, 0.1)
		leftSphere.Material.Diffuse = 0.7
		leftSphere.Material.Specular = 0.3

		world.Objects = []Shape{
			floor,
			leftWall,
			rightWall,
			midSphere,
			leftSphere,
			rightSphere,
		}
		world.Lights = []AreaLight{
			NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
		}
	})
}
