package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithDice() {
	// camera := NewCamera(150, 150, 0.9)
	camera := NewCamera(600, 600, math.Pi/3)

	camera.SetTransform(NewViewTransform(
		NewPoint(2, 3, -2), // NewPoint(0, 3, -4),
		NewPoint(0, 1, 0),
		NewVector(0, 1, 0),
	))

	room := NewCube()
	room.Transform = room.Transform.Compose(
		NewTranslation(0, 1, 0),
		NewUScale(10),
	)
	room.Material.Pattern = NewCheckerPattern(NewColor(1, 1, 1), NewColor(0.9, 0.9, 0.9))
	room.Material.Pattern.Transform = NewUScale(0.05)
	room.Material.Ambient = 0.1
	room.Material.Diffuse = 0.7
	room.Material.Reflective = 0.5

	createDice := func(c Color) *Shape {
		cube := NewCube()
		cube.Transform = cube.Transform.Compose(
			NewTranslation(0, 1, 0),
			NewUScale(1),
		)
		cube.Material.Color = c
		cube.Material.Diffuse = 0.7
		// cube.Material.Ambient = 0
		cube.Material.Specular = 0.3
		cube.Material.Shininess = 100
		cube.Material.Reflective = 0.3

		dots := []struct {
			locations []Tuple
			rotation  Matrix
		}{
			// One - front
			{
				[]Tuple{NewPoint(0, 0, 0)},
				IdentityMatrix(),
			},
			// Two - Left
			{
				[]Tuple{NewPoint(-2, -2, 0), NewPoint(2, 2, 0)},
				NewRotateY(math.Pi / 2),
			},
			// Three - Back
			{
				[]Tuple{NewPoint(0, 0, 9), NewPoint(-3, -3, 9), NewPoint(3, 3, 9)},
				IdentityMatrix(),
			},
			// Four - Right
			{
				[]Tuple{NewPoint(-2, 2, 0), NewPoint(2, 2, 0), NewPoint(-2, -2, 0), NewPoint(2, -2, 0)},
				NewRotateY((3 * math.Pi) / 2),
			},
			// Five - Top
			{
				[]Tuple{NewPoint(0, 5, 4.5), NewPoint(3, 5, 7.5), NewPoint(-3, 5, 1.5), NewPoint(-3, 5, 7.5), NewPoint(3, 5, 1.5)},
				IdentityMatrix(),
			},
			// Six - Bottom
			{
				[]Tuple{NewPoint(-2.5, -5, 4.5), NewPoint(-2.5, -5, 1.5), NewPoint(-2.5, -5, 7.5), NewPoint(2.5, -5, 4.5), NewPoint(2.5, -5, 1.5), NewPoint(2.5, -5, 7.5)},
				IdentityMatrix(),
			},
		}
		dotGroup := NewGroup()
		for _, face := range dots {
			for _, location := range face.locations {
				dot := NewSphere()
				dot.Transform = dot.Transform.Compose(
					NewTranslation(location.X, location.Y, location.Z),
					NewUScale(0.2),
					NewTranslation(0, 1, -0.9),
					face.rotation,
				)
				dot.Shadows = false
				dot.Material.Color = Colors["White"]
				dot.Transform = dot.Transform.Multiply(face.rotation)
				dotGroup.AddChildren(&dot)
			}
		}

		dice := NewCsg("difference", &cube, &dotGroup)
		// To flip the dice over
		// dice.Transform = dice.Transform.Compose(
		// 	NewRotateX(math.Pi),
		// 	NewTranslation(0, 2, 0),
		// )

		return &dice
	}

	d1 := createDice(NewColor(0.2, 0.9, 0.4))
	d2 := createDice(NewColor(0.9, 0.2, 0.4))
	d2.Transform = d2.Transform.Multiply(NewTranslation(-4, 0, 0))
	d3 := createDice(NewColor(0.2, 0.2, 0.9))
	d3.Transform = d3.Transform.Multiply(NewTranslation(0, 0, 4))

	world := NewWorld()
	world.Objects = []Shape{
		room,
		*d1,
		*d2,
		*d3,
	}
	world.Lights = []PointLight{
		NewPointLight(NewPoint(-2, 4, -5), NewColor(1, 1, 1)),
		NewPointLight(NewPoint(2, 2, 5), NewColor(1, 1, 1)),
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