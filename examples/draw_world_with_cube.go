package examples

import (
	"fmt"
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
	cubeLeft.SetTransform(NewTranslation(-2, 1, 1))
	cubeLeft.Material.Color = Colors["Red"]
	cubeLeft.Material.Diffuse = 0.7
	cubeLeft.Material.Specular = 0.3

	cubeMiddle := NewCube()
	cubeMiddle.SetTransform(NewTranslation(-0.5, 3, 1))
	cubeMiddle.Material.Color = Colors["Green"]
	cubeMiddle.Material.Diffuse = 0.7
	cubeMiddle.Material.Specular = 0.3

	cubeRight := NewCube()
	cubeRight.Material.Color = Colors["Blue"]
	cubeRight.SetTransform(NewTranslation(1, 1, 1))
	cubeRight.Material.Diffuse = 0.7
	cubeRight.Material.Specular = 0.3

	world := NewWorld()
	world.Objects = []Shape{
		floor,
		cubeLeft,
		cubeMiddle,
		cubeRight,
	}
	world.Lights = []AreaLight{
		NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
	}

	canvas := camera.Render(world)

	if err := canvas.SavePpm("tmp/world.ppm"); err != nil {
		fmt.Printf("Something went wrong! %s\n", err)
	} else {
		fmt.Println("Saved to tmp/world.ppm")
	}
}
