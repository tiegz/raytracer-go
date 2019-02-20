package examples

import (
	"fmt"

	"github.com/tiegz/raytracer-go/raytracer"
)

type projectile struct {
	position raytracer.Tuple // Point
	velocity raytracer.Tuple // Vector
}

type environment struct {
	gravity raytracer.Tuple // Vector
	wind    raytracer.Tuple // Vector
}

func tick(env environment, proj projectile) projectile {
	pos := proj.position.Add(proj.velocity)
	vel := proj.velocity.Add(env.gravity)
	vel = vel.Add(env.wind)
	return projectile{pos, vel}
}

func RunProjectileExample() {
	initialPos := raytracer.NewPoint(0, 1, 0)
	initialVel := raytracer.NewVector(1, 1, 0)
	initialVel = initialVel.Normalized()
	initialGrav := raytracer.NewVector(0, -0.1, 0)
	initialWind := raytracer.NewVector(-0.01, 0, 0)

	proj := projectile{initialPos, initialVel}
	env := environment{initialGrav, initialWind}

	fmt.Printf("Projectile Example==================\nThe projectile is initially %v.\n", proj)

	for proj.position.Y > 0 {
		proj = tick(env, proj)
		fmt.Printf("The projectile is now %v.\n", proj)
	}
}
