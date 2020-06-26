package examples

import (
	"fmt"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawProjectileExample() {
	initialPos := NewPoint(0, 1, 0)
	initialVel := NewVector(1, 1.8, 0)
	initialVel = initialVel.Normalized()
	initialVel = initialVel.Multiply(11.25)
	initialGrav := NewVector(0, -0.1, 0)
	initialWind := NewVector(-0.01, 0, 0)
	proj := projectile{initialPos, initialVel}
	env := environment{initialGrav, initialWind}
	c := NewCanvas(900, 550)
	color := NewColor(1, 0, 0)

	for proj.position.Y > 0 {
		proj = tick(env, proj)
		x, y := int(proj.position.X), c.Height-int(proj.position.Y)
		if y >= 0 && y < c.Height && x >= 0 && x < c.Width {
			c.WritePixel(x, y, color)
		}
	}

	if err := c.SavePpm("tmp/projectile.ppm"); err != nil {
		fmt.Printf("Something went wrong! %s\n", err)
	} else {
		fmt.Println("Saved to tmp/projectile.ppm")
	}
}
