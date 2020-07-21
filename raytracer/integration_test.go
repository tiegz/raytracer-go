package raytracer

import (
	"math"
	"os"
	"testing"
)

/////////////
// Benchmarks
/////////////

func BenchmarkRenderIntegrationTestScene(b *testing.B) {
	for i := 0; i < b.N; i++ {
		renderIntegrationTestScene()
	}
}

// Renders a tiny JPG image of a sphere on top of a plane.
func renderIntegrationTestScene() {
	world := NewWorld()
	world.Lights = []AreaLight{
		NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
	}
	camera := NewCamera(10, 10, math.Pi/3)
	camera.SetTransform(NewViewTransform(
		NewPoint(0, 1, -2),
		NewPoint(0, 0, 0),
		NewVector(0, 1, 0),
	))

	floor := NewPlane()
	floor.Material.Color = NewColor(1, 0.9, 0.9)
	floor.Material.Specular = 0

	sphere := NewSphere()
	sphere.Material.Color = Colors["Red"]

	world.Objects = []Shape{
		floor,
		sphere,
	}

	filepath := "integration_test_scene.jpg"
	canvas := camera.Render(world)
	if err := canvas.SaveJPEG(filepath); err != nil {
		panic(err)
	}
	if err := os.Remove(filepath); err != nil {
		panic(err)
	}
}
