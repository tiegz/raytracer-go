package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawUVAlignCheck() {
	// camera := NewCamera(160, 70, math.Pi/3)
	camera := NewCamera(400, 400, math.Pi/3)
	// camera := NewCamera(1280, 800, math.Pi/3)

	camera.SetTransform(NewViewTransform(
		NewPoint(1, 2, -5),
		NewPoint(0, 0, 0),
		NewVector(0, 1, 0),
	))

	floor := NewPlane()
	floor.Material.Ambient = 0.1
	floor.Material.Diffuse = 0.8
	floor.Material.Pattern = NewTextureMapPattern(
		NewUVAlignCheckPattern(
			NewColor(1, 1, 1), // white
			NewColor(1, 0, 0), // red
			NewColor(1, 1, 0), // yellow
			NewColor(0, 1, 0), // green
			NewColor(0, 1, 1), // cyan
		),
		PlanarMap,
	)

	world := NewWorld()
	world.Objects = []Shape{
		floor,
	}
	world.Lights = []AreaLight{
		NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
		// NewAreaLight(NewPoint(3, 5, -5), NewVector(4, 0, 0), 4, NewVector(0, 4, 0), 4, NewColor(1, 1, 1)),
	}

	canvas := camera.RenderWithProgress(world)

	fmt.Println("Generating PPM...")
	ppm := canvas.ToPpm()
	filename := "tmp/world.ppm"
	ppmBytes := []byte(ppm)
	fmt.Printf("Saving scene to %s...\n", filename)
	if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
		panic(err)
	}
}
