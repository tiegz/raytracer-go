package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	"github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithPatterns() {
	camera := raytracer.NewCamera(320, 200, math.Pi/3)
	// camera := raytracer.NewCamera(640, 480, math.Pi/3)
	// camera := raytracer.NewCamera(400, 200, math.Pi/3)
	// camera := raytracer.NewCamera(1000, 500, math.Pi/3)
	// camera := raytracer.NewCamera(1920, 1080, math.Pi/3)

	camera.Transform = raytracer.NewViewTransform(
		raytracer.NewPoint(0, 1.5, -5),
		raytracer.NewPoint(0, 1, 0),
		raytracer.NewVector(0, 1, 0),
	)

	floor := raytracer.NewPlane()
	floor.Material.Color = raytracer.NewColor(1, 0.9, 0.9)
	floor.Material.Specular = 0
	floor.Material.Pattern = raytracer.NewStripePattern(raytracer.Colors["White"], raytracer.Colors["Red"])

	midSphere := raytracer.NewSphere()
	midSphere.Transform = raytracer.NewTranslation(-0.5, 1, 0.5)
	midSphere.Material.Color = raytracer.NewColor(0.1, 1, 0.5)
	midSphere.Material.Diffuse = 0.7
	midSphere.Material.Specular = 0.3
	midSphere.Material.Pattern = raytracer.NewStripePattern(raytracer.Colors["Green"], raytracer.Colors["Purple"])

	rightSphere := raytracer.NewSphere()
	rightSphere.Transform = raytracer.NewTranslation(1.5, 0.5, -0.5)
	rightSphere.Transform = rightSphere.Transform.Multiply(raytracer.NewScale(0.5, 0.5, 0.5))
	rightSphere.Material.Color = raytracer.NewColor(0.5, 1, 0.1)
	rightSphere.Material.Diffuse = 0.7
	rightSphere.Material.Specular = 0.3
	rightSphere.Material.Pattern = raytracer.NewStripePattern(raytracer.Colors["Red"], raytracer.Colors["Orange"])
	rightSphere.Material.Pattern.Transform = raytracer.NewScale(0.33, 0.33, 0.33)

	leftSphere := raytracer.NewSphere()
	leftSphere.Transform = raytracer.NewTranslation(-1.5, 0.33, -0.75)
	leftSphere.Transform = leftSphere.Transform.Multiply(raytracer.NewScale(0.33, 0.33, 0.33))
	leftSphere.Material.Pattern = raytracer.NewStripePattern(raytracer.Colors["White"], raytracer.Colors["Black"])
	leftSphere.Material.Pattern.Transform = raytracer.NewScale(0.01, 0.01, 0.01)
	leftSphere.Material.Color = raytracer.NewColor(1, 0.8, 0.1)
	leftSphere.Material.Diffuse = 0.7
	leftSphere.Material.Specular = 0.3

	world := raytracer.NewWorld()
	world.Objects = []raytracer.Shape{
		floor,
		midSphere,
		leftSphere,
		rightSphere,
	}
	world.Lights = []raytracer.PointLight{
		raytracer.NewPointLight(raytracer.NewPoint(-10, 10, -10), raytracer.NewColor(1, 1, 1)),
	}

	canvas := camera.Render(world)

	fmt.Println("Generating PPM...")
	ppm := canvas.ToPpm()
	filename := "tmp/sphere_silhouette.ppm"
	ppmBytes := []byte(ppm)
	fmt.Printf("Saving scene to %s...\n", filename)
	if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
		panic(err)
	}
}
