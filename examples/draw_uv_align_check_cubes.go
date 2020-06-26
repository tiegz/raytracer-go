package examples

import (
	"fmt"
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawUVAlignCheckCubes() {
	// camera := NewCamera(160, 70, math.Pi/3)
	camera := NewCamera(800, 400, math.Pi/3)
	// camera := NewCamera(1280, 800, math.Pi/3)

	camera.SetTransform(NewViewTransform(
		NewPoint(0, 0, -20),
		NewPoint(0, 0, 0),
		NewVector(0, 1, 0),
	))

	material := DefaultMaterial()
	material.Pattern = NewCubeMapPattern(
		NewUVAlignCheckPattern(Colors["Yellow"], Colors["Cyan"], Colors["Red"], Colors["Blue"], Colors["Brown"]),
		NewUVAlignCheckPattern(Colors["Cyan"], Colors["Red"], Colors["Yellow"], Colors["Brown"], Colors["Green"]),
		NewUVAlignCheckPattern(Colors["Red"], Colors["Yellow"], Colors["Purple"], Colors["Green"], Colors["White"]),
		NewUVAlignCheckPattern(Colors["Green"], Colors["Purple"], Colors["Cyan"], Colors["White"], Colors["Blue"]),
		NewUVAlignCheckPattern(Colors["Brown"], Colors["Cyan"], Colors["Purple"], Colors["Red"], Colors["Yellow"]),
		NewUVAlignCheckPattern(Colors["Purple"], Colors["Brown"], Colors["Green"], Colors["Blue"], Colors["White"]),
	)
	material.Ambient = 0.2
	material.Specular = 0
	material.Diffuse = 0.8

	cube1 := NewCube()
	cube1.Material = material
	cube1.SetTransform(cube1.Transform.Compose(
		NewRotateY(0.7854),
		NewRotateX(0.7854),
		NewTranslation(-6, 2, 0),
	))

	cube2 := NewCube()
	cube2.Material = material
	cube2.SetTransform(cube2.Transform.Compose(
		NewRotateY(2.3562),
		NewRotateX(0.7854),
		NewTranslation(-2, 2, 0),
	))

	cube3 := NewCube()
	cube3.Material = material
	cube3.SetTransform(cube3.Transform.Compose(
		NewRotateY(3.927),
		NewRotateX(0.7854),
		NewTranslation(2, 2, 0),
	))

	cube4 := NewCube()
	cube4.Material = material
	cube4.SetTransform(cube4.Transform.Compose(
		NewRotateY(5.4978),
		NewRotateX(0.7854),
		NewTranslation(6, 2, 0),
	))

	cube5 := NewCube()
	cube5.Material = material
	cube5.SetTransform(cube5.Transform.Compose(
		NewRotateY(0.7854),
		NewRotateX(-0.7854),
		NewTranslation(-6, -2, 0),
	))

	cube6 := NewCube()
	cube6.Material = material
	cube6.SetTransform(cube6.Transform.Compose(
		NewRotateY(2.3562),
		NewRotateX(-0.7854),
		NewTranslation(-2, -2, 0),
	))

	cube7 := NewCube()
	cube7.Material = material
	cube7.SetTransform(cube7.Transform.Compose(
		NewRotateY(3.927),
		NewRotateX(-0.7854),
		NewTranslation(2, -2, 0),
	))

	cube8 := NewCube()
	cube8.Material = material
	cube8.SetTransform(cube8.Transform.Compose(
		NewRotateY(5.4978),
		NewRotateX(-0.7854),
		NewTranslation(6, -2, 0),
	))

	world := NewWorld()
	world.Objects = []Shape{
		cube1,
		cube2,
		cube3,
		cube4,
		cube5,
		cube6,
		cube7,
		cube8,
	}

	world.Lights = []AreaLight{
		NewPointLight(NewPoint(0, 100, -100), NewColor(0.25, 0.25, 0.25)),
		NewPointLight(NewPoint(0, -100, -100), NewColor(0.25, 0.25, 0.25)),
		NewPointLight(NewPoint(-100, 0, -100), NewColor(0.25, 0.25, 0.25)),
		NewPointLight(NewPoint(100, 0, -100), NewColor(0.25, 0.25, 0.25)),
	}

	canvas := camera.RenderWithProgress(world)

	if err := canvas.SavePpm("tmp/world.ppm"); err != nil {
		fmt.Printf("Something went wrong! %s\n", err)
	} else {
		fmt.Println("Saved to tmp/world.ppm")
	}
}
