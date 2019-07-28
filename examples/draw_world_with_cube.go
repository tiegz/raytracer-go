package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	"github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithCube() {
	camera := raytracer.NewCamera(320, 200, math.Pi/3)
	// camera := raytracer.NewCamera(640, 480, math.Pi/3)
	// camera := raytracer.NewCamera(400, 200, math.Pi/3)
	// camera := raytracer.NewCamera(1000, 500, math.Pi/3)
	// camera := raytracer.NewCamera(1920, 1080, math.Pi/3)

	camera.Transform = raytracer.NewViewTransform(
		raytracer.NewPoint(0, 3, -10),
		raytracer.NewPoint(0, 2, 0),
		raytracer.NewVector(0, 1, 0),
	)

	floor := raytracer.NewPlane()
	floor.Material.Color = raytracer.NewColor(1, 0.9, 0.9)
	floor.Material.Specular = 0

	cubeLeft := raytracer.NewCube()
	cubeLeft.Transform = raytracer.NewTranslation(-2, 1, 1)
	cubeLeft.Material.Color = raytracer.Colors["Red"]
	cubeLeft.Material.Diffuse = 0.7
	cubeLeft.Material.Specular = 0.3

	cubeMiddle := raytracer.NewCube()
	cubeMiddle.Transform = raytracer.NewTranslation(-0.5, 3, 1)
	cubeMiddle.Material.Color = raytracer.Colors["Green"]
	cubeMiddle.Material.Diffuse = 0.7
	cubeMiddle.Material.Specular = 0.3

	cubeRight := raytracer.NewCube()
	cubeRight.Material.Color = raytracer.Colors["Blue"]
	cubeRight.Transform = raytracer.NewTranslation(1, 1, 1)
	cubeRight.Material.Diffuse = 0.7
	cubeRight.Material.Specular = 0.3

	world := raytracer.NewWorld()
	world.Objects = []raytracer.Shape{
		floor,
		cubeLeft,
		cubeMiddle,
		cubeRight,
	}
	world.Lights = []raytracer.PointLight{
		raytracer.NewPointLight(raytracer.NewPoint(-10, 10, -10), raytracer.NewColor(1, 1, 1)),
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
