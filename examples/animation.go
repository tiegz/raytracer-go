package examples

import (
	"fmt"
	"io/ioutil"
	"math"
	"os/exec"

	"github.com/tiegz/raytracer-go/raytracer"
)

func RunAnimation() {
	// camera := raytracer.NewCamera(160, 100, math.Pi/3)
	// camera := raytracer.NewCamera(320, 200, math.Pi/3)
	camera := raytracer.NewCamera(640, 480, math.Pi/3)
	// camera := raytracer.NewCamera(400, 200, math.Pi/3)
	// camera := raytracer.NewCamera(1000, 500, math.Pi/3)
	// camera := raytracer.NewCamera(1920, 1080, math.Pi/3)

	camera.Transform = raytracer.NewViewTransform(
		raytracer.NewPoint(0, 1.5, -5),
		raytracer.NewPoint(0, 1, 0),
		raytracer.NewVector(0, 1, 0),
	)

	floor := raytracer.NewPlane()
	floor.Material.Pattern = raytracer.NewCheckerPattern(raytracer.Colors["Black"], raytracer.Colors["White"])

	rightWall := raytracer.NewPlane()
	rightWall.Material.Ambient = 0.5
	rightWall.Material.Reflective = 0.5
	rightWall.Transform = rightWall.Transform.Multiply(raytracer.NewTranslation(0, 0, 5))
	rightWall.Transform = rightWall.Transform.Multiply(raytracer.NewRotateY(math.Pi / 4))
	rightWall.Transform = rightWall.Transform.Multiply(raytracer.NewRotateX(math.Pi / 2))
	rightWall.Material.Pattern = raytracer.NewCheckerPattern(raytracer.Colors["Blue"], raytracer.Colors["Green"])

	leftWall := raytracer.NewPlane()
	leftWall.Material.Ambient = 0.5
	leftWall.Material.Reflective = 0.2
	leftWall.Transform = leftWall.Transform.Multiply(raytracer.NewTranslation(0, 0, 5))
	leftWall.Transform = leftWall.Transform.Multiply(raytracer.NewRotateY(-math.Pi / 4))
	leftWall.Transform = leftWall.Transform.Multiply(raytracer.NewRotateX(math.Pi / 2))
	leftWall.Material.Pattern = raytracer.NewCheckerPattern(raytracer.Colors["Purple"], raytracer.Colors["Yellow"])

	midSphere := raytracer.NewSphere()
	midSphere.Transform = raytracer.NewTranslation(-0.5, 1, 0.5)
	midSphere.Material.Reflective = 1.0
	midSphere.Material.Diffuse = 0.1
	midSphere.Material.Color = raytracer.NewColor(0.75, 0.75, 0.75)

	rightSphere := raytracer.NewSphere()
	rightSphere.Transform = raytracer.NewTranslation(1.5, 0.5, -1.5)
	rightSphere.Transform = rightSphere.Transform.Multiply(raytracer.NewScale(0.5, 0.5, 0.5))
	rightSphere.Material.Pattern = raytracer.NewRingPattern(raytracer.Colors["Red"], raytracer.Colors["White"])
	rightSphere.Material.Pattern.Transform = raytracer.NewScale(0.23, 0.1, 0.23)

	leftSphere := raytracer.NewSphere()
	leftSphere.Transform = raytracer.NewTranslation(-1.5, 0.33, -0.75)
	leftSphere.Transform = leftSphere.Transform.Multiply(raytracer.NewScale(0.33, 0.33, 0.33))
	leftSphere.Material.Pattern = raytracer.NewRingPattern(raytracer.Colors["Blue"], raytracer.Colors["White"])
	leftSphere.Material.Pattern.Transform = raytracer.NewScale(0.23, 0.23, 0.23)
	leftSphere.Material.Color = raytracer.NewColor(1, 0.8, 0.1)

	world := raytracer.NewWorld()
	world.Objects = []raytracer.Shape{
		floor,
		rightWall,
		leftWall,
		midSphere,
		leftSphere,
		rightSphere,
	}
	world.Lights = []raytracer.PointLight{
		raytracer.NewPointLight(raytracer.NewPoint(-10, 10, -10), raytracer.NewColor(1, 1, 1)),
	}

	frameCount := 100
	leftSphereTranslation := raytracer.NewTranslation(0, 0.05, 0)
	leftSphereTranslation = leftSphereTranslation.Multiply(raytracer.NewRotateZ(math.Pi / 40))
	leftSphereTranslation = leftSphereTranslation.Multiply(raytracer.NewTranslation(0, 0.1, 0))
	leftSphereTranslation = leftSphereTranslation.Multiply(raytracer.NewRotateX(math.Pi / 20))
	midSphereTranslation := raytracer.NewTranslation(0.075, 0, 0)
	rightSphereTranslation := raytracer.NewTranslation(0, 0.05, 0.1)

	for i := 0; i < frameCount; i++ {
		canvas := camera.Render(world)

		world.Objects[3].Transform = midSphereTranslation.Multiply(world.Objects[3].Transform)
		world.Objects[4].Transform = leftSphereTranslation.Multiply(world.Objects[4].Transform)
		world.Objects[5].Transform = rightSphereTranslation.Multiply(world.Objects[5].Transform)

		fmt.Println("Generating PPM...")
		ppm := canvas.ToPpm()
		filename := fmt.Sprintf("tmp/sphere_silhouette_%03d.ppm", i)
		ppmBytes := []byte(ppm)
		fmt.Printf("Saving scene to %s...\n", filename)
		if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
			panic(err)
		}
	}

	if _, err := exec.Command("convert -delay 5 tmp/sphere_silhouette*ppm movie.gif").Output(); err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
}
