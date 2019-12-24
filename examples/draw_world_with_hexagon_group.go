package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithHexagonGroup() {
	camera := NewCamera(320, 200, math.Pi/3)

	camera.SetTransform(NewViewTransform(
		NewPoint(0, 2, -4),
		NewPoint(0, 0, 0),
		NewVector(0, 1, 0),
	))

	hexagonCorner := func() *Shape {
		corner := NewSphere()
		corner.Transform = corner.Transform.Compose(
			NewUScale(0.25),
			NewTranslation(0, 0, -1),
		)
		return &corner
	}

	hexagonEdge := func() *Shape {
		edge := NewCylinder()
		cyl := edge.LocalShape.(*Cylinder)
		cyl.Minimum = 0
		cyl.Maximum = 1
		edge.Transform = edge.Transform.Compose(
			NewScale(0.25, 1, 0.25),
			NewRotateZ(-math.Pi/2),
			NewRotateY(-math.Pi/6),
			NewTranslation(0, 0, -1),
		)
		return &edge
	}

	hexagonSide := func() *Shape {
		side := NewGroup()
		corner := hexagonCorner()
		edge := hexagonEdge()
		side.AddChildren(corner, edge)
		return &side
	}

	hexagon := func() *Shape {
		hex := hexagonSide()
		for i := 0; i <= 5; i++ {
			side := hexagonSide()
			side.Transform = NewRotateY(float64(i) * math.Pi / 3.0)
			hex.AddChildren(side)
		}
		return hex
	}

	hex := hexagon()

	world := NewWorld()
	world.Objects = []Shape{
		*hex,
	}
	world.Lights = []PointLight{
		NewPointLight(NewPoint(5, 8, -9), NewColor(1, 1, 1)),
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
