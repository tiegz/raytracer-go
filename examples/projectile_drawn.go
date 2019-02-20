package examples

import (
	"fmt"
	"io/ioutil"

	"github.com/tiegz/raytracer-go/raytracer"
)

func RunProjectileDrawnExample() {
	initialPos := raytracer.NewPoint(0, 1, 0)
	initialVel := raytracer.NewVector(1, 1.8, 0)
	initialVel = initialVel.Normalized()
	initialVel = initialVel.Multiply(11.25)
	initialGrav := raytracer.NewVector(0, -0.1, 0)
	initialWind := raytracer.NewVector(-0.01, 0, 0)
	proj := projectile{initialPos, initialVel}
	env := environment{initialGrav, initialWind}
	c := raytracer.NewCanvas(900, 550)
	color := raytracer.NewColor(1, 0, 0)

	for proj.position.Y > 0 {
		proj = tick(env, proj)
		x, y := int(proj.position.X), c.Height-int(proj.position.Y)
		if y >= 0 && y < c.Height && x >= 0 && x < c.Width {
			c.WritePixel(x, y, color)
		}
	}

	fmt.Println("Generating PPM...")
	ppm := c.ToPpm()
	filename := "projectile.ppm"
	ppmBytes := []byte(ppm)
	fmt.Printf("Saving projectile to %s...\n", filename)
	if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
		panic(err)
	}
}
