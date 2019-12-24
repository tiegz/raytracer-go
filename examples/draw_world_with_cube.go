package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithCube() {
	camera := NewCamera(320, 200, math.Pi/3)
	// camera := NewCamera(640, 480, math.Pi/3)
	// camera := NewCamera(400, 200, math.Pi/3)
	// camera := NewCamera(1000, 500, math.Pi/3)
	// camera := NewCamera(1920, 1080, math.Pi/3)

	camera.SetTransform(NewViewTransform(
		NewPoint(0, 3, -10),
		NewPoint(0, 2, 0),
		NewVector(0, 1, 0),
	))

	floor := NewPlane()
	floor.Material.Color = NewColor(1, 0.9, 0.9)
	floor.Material.Specular = 0

	cubeLeft := NewCube()
	cubeLeft.Transform = NewTranslation(-2, 1, 1)
	cubeLeft.Material.Color = Colors["Red"]
	cubeLeft.Material.Diffuse = 0.7
	cubeLeft.Material.Specular = 0.3

	cubeMiddle := NewCube()
	cubeMiddle.Transform = NewTranslation(-0.5, 3, 1)
	cubeMiddle.Material.Color = Colors["Green"]
	cubeMiddle.Material.Diffuse = 0.7
	cubeMiddle.Material.Specular = 0.3

	cubeRight := NewCube()
	cubeRight.Material.Color = Colors["Blue"]
	cubeRight.Transform = NewTranslation(1, 1, 1)
	cubeRight.Material.Diffuse = 0.7
	cubeRight.Material.Specular = 0.3

	world := NewWorld()
	world.Objects = []Shape{
		floor,
		cubeLeft,
		cubeMiddle,
		cubeRight,
	}
	world.Lights = []PointLight{
		NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
	}

	canvas := camera.Render(world)

	fmt.Println("Generating PPM...")
	ppm := canvas.ToPpm()
	filename := "tmp/cube.ppm"
	ppmBytes := []byte(ppm)
	fmt.Printf("Saving scene to %s...\n", filename)
	if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
		panic(err)
	}
}
