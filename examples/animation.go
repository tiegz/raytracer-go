package examples

import (
	"fmt"
	"math"
	"os/exec"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunAnimation() {
	// camera := NewCamera(160, 100, math.Pi/3)
	// camera := NewCamera(320, 200, math.Pi/3)
	camera := NewCamera(640, 480, math.Pi/3)
	// camera := NewCamera(400, 200, math.Pi/3)
	// camera := NewCamera(1000, 500, math.Pi/3)
	// camera := NewCamera(1920, 1080, math.Pi/3)

	camera.SetTransform(NewViewTransform(
		NewPoint(0, 1.5, -5),
		NewPoint(0, 1, 0),
		NewVector(0, 1, 0),
	))

	floor := NewPlane()
	floor.Material.Pattern = NewCheckerPattern(Colors["Black"], Colors["White"])

	rightWall := NewPlane()
	rightWall.Material.Ambient = 0.5
	rightWall.Material.Reflective = 0.5
	rightWall.SetTransform(rightWall.Transform.Multiply(NewTranslation(0, 0, 5)))
	rightWall.SetTransform(rightWall.Transform.Multiply(NewRotateY(math.Pi / 4)))
	rightWall.SetTransform(rightWall.Transform.Multiply(NewRotateX(math.Pi / 2)))
	rightWall.Material.Pattern = NewCheckerPattern(Colors["Blue"], Colors["Green"])

	leftWall := NewPlane()
	leftWall.Material.Ambient = 0.5
	leftWall.Material.Reflective = 0.2
	leftWall.SetTransform(leftWall.Transform.Multiply(NewTranslation(0, 0, 5)))
	leftWall.SetTransform(leftWall.Transform.Multiply(NewRotateY(-math.Pi / 4)))
	leftWall.SetTransform(leftWall.Transform.Multiply(NewRotateX(math.Pi / 2)))
	leftWall.Material.Pattern = NewCheckerPattern(Colors["Purple"], Colors["Yellow"])

	midSphere := NewSphere()
	midSphere.SetTransform(NewTranslation(-0.5, 1, 0.5))
	midSphere.Material.Reflective = 1.0
	midSphere.Material.Diffuse = 0.1
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

	world := NewWorld()
	world.Objects = []Shape{
		floor,
		rightWall,
		leftWall,
		midSphere,
		leftSphere,
		rightSphere,
	}
	world.Lights = []AreaLight{
		NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
	}

	frameCount := 100
	leftSphereTranslation := NewTranslation(0, 0.05, 0)
	leftSphereTranslation = leftSphereTranslation.Multiply(NewRotateZ(math.Pi / 40))
	leftSphereTranslation = leftSphereTranslation.Multiply(NewTranslation(0, 0.1, 0))
	leftSphereTranslation = leftSphereTranslation.Multiply(NewRotateX(math.Pi / 20))
	midSphereTranslation := NewTranslation(0.075, 0, 0)
	rightSphereTranslation := NewTranslation(0, 0.05, 0.1)

	for i := 0; i < frameCount; i++ {
		canvas := camera.Render(world)

		world.Objects[3].SetTransform(midSphereTranslation.Multiply(world.Objects[3].Transform))
		world.Objects[4].SetTransform(leftSphereTranslation.Multiply(world.Objects[4].Transform))
		world.Objects[5].SetTransform(rightSphereTranslation.Multiply(world.Objects[5].Transform))

		fmt.Println("Generating PPM...")

		filepath := fmt.Sprintf("tmp/world_%03d.ppm", i)
		if err := canvas.SavePpm(filepath); err != nil {
			fmt.Printf("Something went wrong! %s\n", err)
		} else {
			fmt.Printf("Saved to %s\n", filepath)
		}
	}

	if _, err := exec.Command("convert -delay 5 tmp/world*ppm movie.gif").Output(); err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
}
