package examples

import (
	"fmt"
	"math"
	"strings"

	. "github.com/tiegz/raytracer-go/raytracer"
)

// Calls drawFunc() with a World and Camera, and saves it to filepath.
// Example:
//   Draw("tmp/world.jpg", func(world *World, camera *Camera) {
//		 floor := NewPlane()
// 	   cube := NewCube()
//     cube.SetTransform(NewTranslation(0, 1, 1))
//     cube.Material.Color = Colors["Red"]
//     world.Objects = []Shape{floor, cube}
//   }
func Draw(filepath string, drawFunc func(*World, *Camera)) {
	world := NewWorld()
	world.Lights = []AreaLight{
		NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
	}
	camera := NewCamera(800, 400, math.Pi/3)
	camera.SetTransform(NewViewTransform(
		NewPoint(0, 1, -5),
		NewPoint(0, 1, 0),
		NewVector(0, 1, 0),
	))
	drawFunc(&world, &camera)
	canvas := camera.RenderWithProgress(world)

	var err error
	if strings.HasSuffix(filepath, ".jpg") || strings.HasSuffix(filepath, ".jpeg") {
		err = canvas.SaveJPEG(filepath)
	} else if strings.HasSuffix(filepath, ".png") {
		err = canvas.SavePNG(filepath)
	} else if strings.HasSuffix(filepath, ".gif") {
		err = canvas.SaveGIF(filepath)
	} else {
		fmt.Printf("Unsupported file extension: %s\n", filepath)
	}

	if err != nil {
		fmt.Printf("Something went wrong! %s\n", err)
	} else {
		fmt.Printf("Saved to %s\n", filepath)
	}
}
