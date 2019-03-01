package examples

import (
	"fmt"
	"io/ioutil"

	"github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawSphereSilhouette() {
	canvasSize := 250
	wallZ := 10.0
	wallSize := 7.0
	halfWallSize := float64(wallSize) / 2.0
	pixelSize := wallSize / float64(canvasSize)

	rayOrigin := raytracer.NewPoint(0, 0, -5)
	canvas := raytracer.NewCanvas(int(canvasSize), int(canvasSize))
	sphere := raytracer.NewSphere()

	// # for each row of pixels in the canvas
	for y := 0; y < canvasSize; y++ {
		// # compute the world y coordinate (top = +half, bottom = -half)
		worldY := halfWallSize - (pixelSize * float64(y))
		// for each col of pixels in the canvas
		for x := 0; x < canvasSize; x++ {
			// # compute the world x coordinate (left = -half, right = half)
			worldX := -halfWallSize + (pixelSize * float64(x))

			// # describe the point on the wall that the ray will target
			targetPoint := raytracer.NewPoint(worldX, worldY, wallZ)
			rayDirection := targetPoint.Subtract(rayOrigin)
			r := raytracer.NewRay(rayOrigin, rayDirection)
			intersections := r.Intersect(sphere)

			if hit := intersections.Hit(); !hit.IsEqualTo(raytracer.NullIntersection()) {
				fmt.Printf("(%d, %d) found a hit \n", x, y)
				canvas.WritePixel(int(x), int(y), raytracer.Colors["Purple"])
			}
		}
	}

	fmt.Println("Generating PPM...")
	ppm := canvas.ToPpm()
	filename := "tmp/sphere_silhouette.ppm"
	ppmBytes := []byte(ppm)
	fmt.Printf("Saving analog clock to %s...\n", filename)
	if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
		panic(err)
	}
}
