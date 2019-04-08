package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	"github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorld() {
	camera := raytracer.NewCamera(320, 200, math.Pi/3) // VGA
	// camera := raytracer.NewCamera(640, 480, math.Pi/3) // VGA
	// camera := raytracer.NewCamera(400, 200, math.Pi/3)
	// camera := raytracer.NewCamera(1000, 500, math.Pi/3)
	// camera := raytracer.NewCamera(1920, 1080, math.Pi/3) // VGA

	camera.Transform = raytracer.NewViewTransform(
		raytracer.NewPoint(0, 1.5, -5),
		raytracer.NewPoint(0, 1, 0),
		raytracer.NewVector(0, 1, 0),
	)

	floor := raytracer.NewSphere()
	floor.Transform = raytracer.NewScale(10, 0.01, 10)
	floor.Material.Color = raytracer.NewColor(1, 0.9, 0.9)
	floor.Material.Specular = 0

	leftWall := raytracer.NewSphere()
	leftWall.Transform = raytracer.NewTranslation(0, 0, 5)
	leftWall.Transform = leftWall.Transform.Multiply(raytracer.NewRotateY(-math.Pi / 4))
	leftWall.Transform = leftWall.Transform.Multiply(raytracer.NewRotateX(math.Pi / 2))
	leftWall.Transform = leftWall.Transform.Multiply(raytracer.NewScale(10, 0.01, 10))
	leftWall.Material = floor.Material

	rightWall := raytracer.NewSphere()
	rightWall.Transform = raytracer.NewTranslation(0, 0, 5)
	rightWall.Transform = rightWall.Transform.Multiply(raytracer.NewRotateY(math.Pi / 4))
	rightWall.Transform = rightWall.Transform.Multiply(raytracer.NewRotateX(math.Pi / 2))
	rightWall.Transform = rightWall.Transform.Multiply(raytracer.NewScale(10, 0.01, 10))
	rightWall.Material = floor.Material

	midSphere := raytracer.NewSphere()
	midSphere.Transform = raytracer.NewTranslation(-0.5, 1, 0.5)
	midSphere.Material.Color = raytracer.NewColor(0.1, 1, 0.5)
	midSphere.Material.Diffuse = 0.7
	midSphere.Material.Specular = 0.3

	rightSphere := raytracer.NewSphere()
	rightSphere.Transform = raytracer.NewTranslation(1.5, 0.5, -0.5)
	rightSphere.Transform = rightSphere.Transform.Multiply(raytracer.NewScale(0.5, 0.5, 0.5))
	rightSphere.Material.Color = raytracer.NewColor(0.5, 1, 0.1)
	rightSphere.Material.Diffuse = 0.7
	rightSphere.Material.Specular = 0.3

	leftSphere := raytracer.NewSphere()
	leftSphere.Transform = raytracer.NewTranslation(-1.5, 0.33, -0.75)
	leftSphere.Transform = leftSphere.Transform.Multiply(raytracer.NewScale(0.33, 0.33, 0.33))
	leftSphere.Material.Color = raytracer.NewColor(1, 0.8, 0.1)
	leftSphere.Material.Diffuse = 0.7
	leftSphere.Material.Specular = 0.3

	world := raytracer.NewWorld()
	world.Objects = []raytracer.Shape{
		floor,
		leftWall,
		rightWall,
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
	fmt.Printf("Saving analog clock to %s...\n", filename)
	if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
		panic(err)
	}
}
