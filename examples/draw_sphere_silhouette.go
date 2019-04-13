package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	"github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawSphereSilhouette(scaleX, scaleY, rotateZ, skew bool) {
	canvasSize := 250
	wallZ := 10.0
	wallSize := 7.0
	halfWallSize := float64(wallSize) / 2.0
	pixelSize := wallSize / float64(canvasSize)

	rayOrigin := raytracer.NewPoint(0, 0, -5)
	canvas := raytracer.NewCanvas(int(canvasSize), int(canvasSize))
	sphere := raytracer.NewSphere()
	transform := raytracer.IdentityMatrix()
	if scaleX {
		scale := raytracer.NewScale(1, 0.5, 1)
		transform = scale.Multiply(transform)
	}
	if scaleY {
		scale := raytracer.NewScale(0.5, 1, 1)
		transform = scale.Multiply(transform)
	}
	if rotateZ {
		rotate := raytracer.NewRotateZ(math.Pi / 4)
		transform = rotate.Multiply(transform)
	}
	if skew {
		shear := raytracer.NewShear(1, 0, 0, 0, 0, 0)
		transform = shear.Multiply(transform)
	}
	sphere.Transform = transform

	// for each row of pixels in the canvas
	for y := 0; y < canvasSize; y++ {
		// compute the world y coordinate (top = +half, bottom = -half)
		worldY := halfWallSize - (pixelSize * float64(y))
		// for each col of pixels in the canvas
		for x := 0; x < canvasSize; x++ {
			// compute the world x coordinate (left = -half, right = half)
			worldX := -halfWallSize + (pixelSize * float64(x))

			// describe the point on the wall that the ray will target
			targetPoint := raytracer.NewPoint(worldX, worldY, wallZ)
			rayDirection := targetPoint.Subtract(rayOrigin)
			r := raytracer.NewRay(rayOrigin, rayDirection)
			intersections := sphere.Intersect(r)

			if hit := intersections.Hit(); !hit.IsNull() {
				canvas.WritePixel(int(x), int(y), raytracer.Colors["Purple"])
			}
		}
	}

	fmt.Println("Generating PPM...")
	ppm := canvas.ToPpm()
	filename := "tmp/sphere_silhouette.ppm"
	ppmBytes := []byte(ppm)
	fmt.Printf("Saving scene to %s...\n", filename)
	if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
		panic(err)
	}
}
