package examples

import (
	"fmt"

	. "github.com/tiegz/raytracer-go/pkg/raytracer"
)

type projectile struct {
	position Tuple // Point
	velocity Tuple // Vector
}

type environment struct {
	gravity Tuple // Vector
	wind    Tuple // Vector
}

func tick(env environment, proj projectile) projectile {
	pos := proj.position.Add(proj.velocity)
	vel := proj.velocity.Add(env.gravity)
	vel = vel.Add(env.wind)
	return projectile{pos, vel}
}

func RunProjectileExample() {
	initialPos := NewPoint(0, 1, 0)
	initialVel := NewVector(1, 1, 0)
	initialVel = initialVel.Normalized()
	initialGrav := NewVector(0, -0.1, 0)
	initialWind := NewVector(-0.01, 0, 0)

	proj := projectile{initialPos, initialVel}
	env := environment{initialGrav, initialWind}

	fmt.Printf("Projectile Example==================\nThe projectile is initially %v.\n", proj)

	for proj.position.Y > 0 {
		proj = tick(env, proj)
		fmt.Printf("The projectile is now %v.\n", proj)
	}
}
