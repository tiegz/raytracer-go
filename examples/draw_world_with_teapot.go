package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	"github.com/tiegz/raytracer-go/raytracer"
	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithTeapot() {
	// camera := NewCamera(320, 200, math.Pi/3)
	camera := NewCamera(640, 400, math.Pi/3)

	camera.Transform = NewViewTransform(
		NewPoint(0, 20, -20),
		NewPoint(2, 5, 0),
		NewVector(0, 1, 0),
	)

	floor := raytracer.NewPlane()
	floor.Material.Pattern = raytracer.NewCheckerPattern(raytracer.Colors["White"], raytracer.Colors["Gray"])

	backWall := raytracer.NewPlane()
	backWall.Transform = backWall.Transform.Compose(
		NewRotateX(math.Pi/2),
		NewTranslation(0, 0, 20),
	)
	backWall.Material.Pattern = raytracer.NewCheckerPattern(raytracer.Colors["White"], raytracer.Colors["Gray"])

	dat, err := ioutil.ReadFile("raytracer/files/utah_teapot_hires.obj") // ("raytracer/files/mit_teapot.obj")
	if err != nil {
		panic(err)
	}

	objFile := ParseObjFile(string(dat))
	group := objFile.ToGroup()
	group.Transform = group.Transform.Compose(
		NewRotateX(-math.Pi/2),
		NewUScale(0.75),
	)
	// TODO: color on group not working?
	// group.Material.Color = Colors["Red"]

	world := NewWorld()
	world.Objects = []Shape{
		floor,
		backWall,
		group,
	}
	world.Lights = []PointLight{
		NewPointLight(NewPoint(0, -10, -5), NewColor(1, 1, 1)),
		NewPointLight(NewPoint(0, 10, -5), NewColor(1, 1, 1)),
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
