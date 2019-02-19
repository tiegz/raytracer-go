package examples

import (
  "github.com/tiegz/raytracer-go/raytracer"
  "fmt"
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

  fmt.Printf("Projectile Example\n==================\nThe projectile is initially %v.\n", proj)

  for proj.position.Y > 0 {
    proj = tick(env, proj)
    fmt.Printf("The projectile is now %v.\n", proj)
  }
}

// • A projectile has a position (a point) and a velocity (a vector). •
// An environment has gravity (a vector) and wind (a vector).
// Then, add a tick(environment, projectile) function which returns a
// new projectile, representing the given projectile after one unit of
// time has passed. (The actual units here don’t really matter—maybe they’re
//   seconds, or milliseconds. Whatever. We’ll just call them “ticks”.)
// In pseudocode, the tick() function should do the following:
// function tick(env, proj)
// position ← proj.position + proj.velocity
// velocity ← proj.velocity + env.gravity + env.wind return projectile(position, velocity)
// end function
// Now, initialize a projectile, and an environment. Use whatever values you want, but these might get you started:
// # projectile starts one unit above the origin.
// # velocity is normalized to 1 unit/tick.
// p ← projectile(point(0, 1, 0), normalize(vector(1, 1, 0)))
// # gravity -0.1 unit/tick, and wind is -0.01 unit/tick.
// e ← environment(vector(0, -0.1, 0), vector(-0.01, 0, 0))


//   }
