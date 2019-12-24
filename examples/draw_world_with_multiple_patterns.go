package examples

import (
	"fmt"
	"io/ioutil"
	"math"
	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithMultiplePatterns() {
	// camera := NewCamera(320, 200, math.Pi/3)
	// camera := NewCamera(640, 480, math.Pi/3)
	// camera := NewCamera(400, 200, math.Pi/3)
	// camera := NewCamera(1000, 500, math.Pi/3)
	camera := NewCamera(1920, 1080, math.Pi/3)

	camera.Transform = NewViewTransform(
		NewPoint(0, 1.5, -5),
		NewPoint(0, 1, 0),
		NewVector(0, 1, 0),
	)

	floor := NewPlane()
	// floor.Material.Color = NewColor(1, 0.9, 0.9)
	// floor.Material.Specular = 0
	floor.Material.Pattern = NewCheckerPattern(Colors["White"], Colors["Red"])

	midSphere := NewSphere()
	midSphere.Transform = NewTranslation(-0.5, 1, 0.5)
	// midSphere.Material.Color = NewColor(0.1, 1, 0.5)
	// midSphere.Material.Diffuse = 0.7
	// midSphere.Material.Specular = 0.3
	midSphere.Material.Pattern = NewStripePattern(Colors["Green"], Colors["Purple"])

	rightSphere := NewSphere()
	rightSphere.Transform = NewTranslation(1.5, 0.5, -0.5)
	rightSphere.Transform = rightSphere.Transform.Multiply(NewScale(0.5, 0.5, 0.5))
	// rightSphere.Material.Color = NewColor(0.5, 1, 0.1)
	// rightSphere.Material.Diffuse = 0.7
	// rightSphere.Material.Specular = 0.3
	rightSphere.Material.Pattern = NewRingPattern(Colors["Red"], Colors["White"])
	rightSphere.Material.Pattern.Transform = NewScale(0.23, 0.23, 0.23)

	leftSphere := NewSphere()
	leftSphere.Transform = NewTranslation(-1.5, 0.33, -0.75)
	leftSphere.Transform = leftSphere.Transform.Multiply(NewScale(0.33, 0.33, 0.33))
	leftSphere.Material.Pattern = NewCheckerPattern(Colors["White"], Colors["Black"])
	// leftSphere.Material.Pattern.Transform = NewScale(0.01, 0.01, 0.01)
	leftSphere.Material.Color = NewColor(1, 0.8, 0.1)
	leftSphere.Material.Diffuse = 0.7
	leftSphere.Material.Specular = 0.3

	world := NewWorld()
	world.Objects = []Shape{
		floor,
		midSphere,
		leftSphere,
		rightSphere,
	}
	world.Lights = []PointLight{
		NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
	}

	canvas := camera.Render(world)

	fmt.Println("Generating PPM...")
	ppm := canvas.ToPpm()
	filename := "tmp/scene.ppm"
	ppmBytes := []byte(ppm)
	fmt.Printf("Saving scene to %s...\n", filename)
	if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
		panic(err)
	}
}
