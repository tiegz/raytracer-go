package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorld() {
	camera := NewCamera(320, 200, math.Pi/3) // VGA
	// camera := NewCamera(640, 480, math.Pi/3) // VGA
	// camera := NewCamera(400, 200, math.Pi/3)
	// camera := NewCamera(1000, 500, math.Pi/3)
	// camera := NewCamera(1920, 1080, math.Pi/3) // VGA

	camera.Transform = NewViewTransform(
		NewPoint(0, 1.5, -5),
		NewPoint(0, 1, 0),
		NewVector(0, 1, 0),
	)

	floor := NewSphere()
	floor.Transform = NewScale(10, 0.01, 10)
	floor.Material.Color = NewColor(1, 0.9, 0.9)
	floor.Material.Specular = 0

	leftWall := NewSphere()
	leftWall.Transform = NewTranslation(0, 0, 5)
	leftWall.Transform = leftWall.Transform.Multiply(NewRotateY(-math.Pi / 4))
	leftWall.Transform = leftWall.Transform.Multiply(NewRotateX(math.Pi / 2))
	leftWall.Transform = leftWall.Transform.Multiply(NewScale(10, 0.01, 10))
	leftWall.Material = floor.Material

	rightWall := NewSphere()
	rightWall.Transform = NewTranslation(0, 0, 5)
	rightWall.Transform = rightWall.Transform.Multiply(NewRotateY(math.Pi / 4))
	rightWall.Transform = rightWall.Transform.Multiply(NewRotateX(math.Pi / 2))
	rightWall.Transform = rightWall.Transform.Multiply(NewScale(10, 0.01, 10))
	rightWall.Material = floor.Material

	midSphere := NewSphere()
	midSphere.Transform = NewTranslation(-0.5, 1, 0.5)
	midSphere.Material.Color = NewColor(0.1, 1, 0.5)
	midSphere.Material.Diffuse = 0.7
	midSphere.Material.Specular = 0.3

	rightSphere := NewSphere()
	rightSphere.Transform = NewTranslation(1.5, 0.5, -0.5)
	rightSphere.Transform = rightSphere.Transform.Multiply(NewScale(0.5, 0.5, 0.5))
	rightSphere.Material.Color = NewColor(0.5, 1, 0.1)
	rightSphere.Material.Diffuse = 0.7
	rightSphere.Material.Specular = 0.3

	leftSphere := NewSphere()
	leftSphere.Transform = NewTranslation(-1.5, 0.33, -0.75)
	leftSphere.Transform = leftSphere.Transform.Multiply(NewScale(0.33, 0.33, 0.33))
	leftSphere.Material.Color = NewColor(1, 0.8, 0.1)
	leftSphere.Material.Diffuse = 0.7
	leftSphere.Material.Specular = 0.3

	world := NewWorld()
	world.Objects = []Shape{
		floor,
		leftWall,
		rightWall,
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
	filename := "tmp/sphere_silhouette.ppm"
	ppmBytes := []byte(ppm)
	fmt.Printf("Saving scene to %s...\n", filename)
	if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
		panic(err)
	}
}
