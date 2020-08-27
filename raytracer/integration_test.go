package raytracer

import (
	"math"
	"os"
	"testing"
)

/////////////
// Benchmarks
/////////////

// As of this writing, GH Actions vms have 2 cores, so bench w/up to 2 jobs.
// (https://docs.github.com/en/actions/reference/virtual-environments-for-github-hosted-runners)
func BenchmarkRenderIntegrationTestScene(b *testing.B)  { benchmarkRenderIntegrationTestScene(1, b) }
func BenchmarkRenderIntegrationTestScene2(b *testing.B) { benchmarkRenderIntegrationTestScene(2, b) }

func benchmarkRenderIntegrationTestScene(jobs int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		renderIntegrationTestScene(jobs)
	}
}

// Renders a tiny JPG image of a sphere on top of a plane.
func renderIntegrationTestScene(jobs int) {
	world := NewWorld()
	world.Lights = []*AreaLight{
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

	world.Objects = []*Shape{
		floor,
		sphere,
	}

	filepath := "integration_test_scene.jpg"
	canvas := camera.Render(world, jobs, false)
	if err := canvas.SaveJPEG(filepath); err != nil {
		panic(err)
	}
	if err := os.Remove(filepath); err != nil {
		panic(err)
	}
}
