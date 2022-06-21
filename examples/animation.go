package examples

import (
	"fmt"
	"math"
	"os/exec"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunAnimation(printProgress bool, jobs int) {
	midSphere := NewSphere()
	midSphere.SetTransform(NewTranslation(-0.5, 1, 0.5))
	midSphere.Material.Reflective = 1.0
	midSphere.Material.Diffuse = NewColor(0.1, 0.1, 0.1)
	midSphere.Material.Color = NewColor(0.75, 0.75, 0.75)

	rightSphere := NewSphere()
	rightSphere.SetTransform(NewTranslation(1.5, 0.5, -1.5))
	rightSphere.SetTransform(rightSphere.Transform.Multiply(NewScale(0.5, 0.5, 0.5)))
	rightSphere.Material.Pattern = NewRingPattern(Colors["Red"], Colors["White"])
	rightSphere.Material.Pattern.SetTransform(NewScale(0.23, 0.1, 0.23))

	leftSphere := NewSphere()
	leftSphere.SetTransform(NewTranslation(-1.5, 0.33, -0.75))
	leftSphere.SetTransform(leftSphere.Transform.Multiply(NewScale(0.33, 0.33, 0.33)))
	leftSphere.Material.Pattern = NewRingPattern(Colors["Blue"], Colors["White"])
	leftSphere.Material.Pattern.SetTransform(NewScale(0.23, 0.23, 0.23))
	leftSphere.Material.Color = NewColor(1, 0.8, 0.1)

	leftSphereTranslation := NewTranslation(0, 0.05, 0)
	leftSphereTranslation = leftSphereTranslation.Multiply(NewRotateZ(math.Pi / 40))
	leftSphereTranslation = leftSphereTranslation.Multiply(NewTranslation(0, 0.1, 0))
	leftSphereTranslation = leftSphereTranslation.Multiply(NewRotateX(math.Pi / 20))
	midSphereTranslation := NewTranslation(0.075, 0, 0)
	rightSphereTranslation := NewTranslation(0, 0.05, 0.1)

	frameCount := 100

	for i := 0; i < frameCount; i++ {
		midSphere.SetTransform(midSphereTranslation.Multiply(midSphere.Transform))
		leftSphere.SetTransform(leftSphereTranslation.Multiply(leftSphere.Transform))
		rightSphere.SetTransform(rightSphereTranslation.Multiply(rightSphere.Transform))

		Draw(printProgress, jobs, fmt.Sprintf("tmp/world_%03d.jpg", i), func(world *World, camera *Camera) {
			camera.HSize = 320
			camera.VSize = 245
			camera.FieldOfView = math.Pi / 3

			camera.SetTransform(NewViewTransform(
				NewPoint(0, 1.5, -5),
				NewPoint(0, 1, 0),
				NewVector(0, 1, 0),
			))

			floor := NewPlane()
			floor.Material.Pattern = NewCheckerPattern(Colors["Black"], Colors["White"])

			rightWall := NewPlane()
			rightWall.Material.Ambient = NewColor(0.5, 0.5, 0.5)
			rightWall.Material.Reflective = 0.5
			rightWall.SetTransform(rightWall.Transform.Multiply(NewTranslation(0, 0, 5)))
			rightWall.SetTransform(rightWall.Transform.Multiply(NewRotateY(math.Pi / 4)))
			rightWall.SetTransform(rightWall.Transform.Multiply(NewRotateX(math.Pi / 2)))
			rightWall.Material.Pattern = NewCheckerPattern(Colors["Blue"], Colors["Green"])

			leftWall := NewPlane()
			leftWall.Material.Ambient = NewColor(0.5, 0.5, 0.5)
			leftWall.Material.Reflective = 0.2
			leftWall.SetTransform(leftWall.Transform.Multiply(NewTranslation(0, 0, 5)))
			leftWall.SetTransform(leftWall.Transform.Multiply(NewRotateY(-math.Pi / 4)))
			leftWall.SetTransform(leftWall.Transform.Multiply(NewRotateX(math.Pi / 2)))
			leftWall.Material.Pattern = NewCheckerPattern(Colors["Purple"], Colors["Yellow"])

			world.Objects = []*Shape{
				floor,
				rightWall,
				leftWall,
				midSphere,
				leftSphere,
				rightSphere,
			}
			world.Lights = []*AreaLight{
				NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
			}
		})
	}

	if _, err := exec.Command("convert -delay 5 tmp/world*jpg tmp/movie.gif").Output(); err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
}
