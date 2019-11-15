package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithTriangles() {
	camera := NewCamera(320, 200, math.Pi/3)

	camera.Transform = NewViewTransform(
		NewPoint(-3, 2, 3),
		NewPoint(0, 0, 0),
		NewVector(0, 1, 0),
	)

	pyramid := func() *Shape {
		bottom := NewTriangle(NewPoint(-1, 0, 0), NewPoint(0, 0, 1), NewPoint(1, 0, 0))
		bottom.Material.Color = Colors["White"]
		bottom.Material.Reflective = 0.8
		bottom.Material.Ambient = 0.5
		bottom.Material.Transparency = 0.9
		side1 := NewTriangle(NewPoint(-1, 0, 0), NewPoint(0, 0, 1), NewPoint(0, 1, 0))
		side1.Material.Color = Colors["Green"]
		side1.Material.Reflective = 0.8
		side1.Material.Ambient = 0.5
		side1.Material.Transparency = 0.9
		side2 := NewTriangle(NewPoint(0, 1, 0), NewPoint(0, 0, 1), NewPoint(1, 0, 0))
		side2.Material.Color = Colors["Blue"]
		side2.Material.Reflective = 0.8
		side2.Material.Ambient = 0.5
		side2.Material.Transparency = 0.9
		side3 := NewTriangle(NewPoint(-1, 0, 0), NewPoint(0, 1, 0), NewPoint(1, 0, 0))
		side3.Material.Color = Colors["Red"]
		side3.Material.Reflective = 0.8
		side3.Material.Ambient = 0.5
		side3.Material.Transparency = 0.9
		g := NewGroup()
		g.AddChildren(&side1, &side2, &side3, &bottom)
		return &g
	}

	pyramid1 := pyramid()
	pyramid2 := pyramid()
	pyramid2.Transform = pyramid2.Transform.Compose(
		NewTranslation(-1.5, 0, -1.5),
		NewRotateY(math.Pi),
	)
	pyramid3 := pyramid()
	pyramid3.Transform = pyramid3.Transform.Compose(
		NewTranslation(2, 0, 0),
		NewRotateY(math.Pi/2),
	)

	world := NewWorld()
	world.Objects = []Shape{
		*pyramid1,
		*pyramid2,
		*pyramid3,
	}
	world.Lights = []PointLight{
		NewPointLight(NewPoint(0, 5, 0), NewColor(1, 1, 1)),
	}

	canvas := camera.Render(world)

	fmt.Println("Generating PPM...")
	ppm := canvas.ToPpm()
	filename := "tmp/world.ppm"
	ppmBytes := []byte(ppm)
	fmt.Printf("Saving scene to %s...\n", filename)
	if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
		panic(err)
	}
}
