package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithCubeOfSpheres() {
	camera := NewCamera(200, 200, math.Pi/3)
	// camera := NewCamera(640, 400, math.Pi/3)

	camera.SetTransform(NewViewTransform(
		NewPoint(-5, 15, -5),
		NewPoint(0, 9, 0),
		NewVector(0, 1, 0),
	))

	sphereRowCount := 10.0

	group := NewGroup()
	for x := 0.0; x < sphereRowCount; x++ {
		for y := 0.0; y < sphereRowCount; y++ {
			for z := 0.0; z < sphereRowCount; z++ {
				sphere := NewSphere()
				sphere.Transform = NewTranslation(x, y, z)
				sphere.Transform = sphere.Transform.Multiply(NewUScale(0.5))
				sphere.Material.Color = NewColor(x/10.0, y/10.0, z/10.0)
				group.AddChildren(&sphere)
			}
		}
	}
	fmt.Printf("Dividing...")
	group.Divide(4)
	fmt.Printf(" ... done dividing!\n")

	world := NewWorld()
	world.Objects = []Shape{
		group,
	}
	world.Lights = []PointLight{
		NewPointLight(NewPoint(0, -10, -5), NewColor(1, 1, 1)),
		NewPointLight(NewPoint(0, 10, -5), NewColor(1, 1, 1)),
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
