package examples

import (
	"fmt"
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithUVPattern() {
	// camera := NewCamera(160, 70, math.Pi/3)
	camera := NewCamera(320, 140, math.Pi/3)
	// camera := NewCamera(1280, 800, math.Pi/3)

	camera.SetTransform(NewViewTransform(
		NewPoint(0, 3, -10),
		NewPoint(0, 1, 0),
		NewVector(0, 1, 0),
	))

	floor := NewPlane()
	floor.Material.Pattern = NewTextureMapPattern(
		NewUVCheckerPattern(1, 1, Colors["White"], Colors["Gray"]),
		PlanarMap,
	)

	backWall := NewPlane()
	backWall.SetTransform(backWall.Transform.Compose(
		NewRotateX(math.Pi/2),
		NewTranslation(0, 0, 5),
	))
	backWall.Material.Pattern = NewCheckerPattern(Colors["White"], Colors["Gray"])

	sphere := NewSphere()
	sphere.SetTransform(NewTranslation(3, 1, 0))
	sphere.Material.Color = NewColor(0.1, 1, 0.5)
	sphere.Material.Diffuse = 0.7
	sphere.Material.Specular = 0.3
	sphere.Material.Pattern = NewTextureMapPattern(
		NewUVCheckerPattern(16, 16, Colors["Green"], Colors["Purple"]),
		SphericalMap,
	)

	cube := NewCube()
	cube.SetTransform(NewTranslation(0, 1, 0))
	cube.Material.Diffuse = 0.7
	cube.Material.Specular = 0.3
	cube.Material.Pattern = NewTextureMapPattern(
		NewUVCheckerPattern(16, 16, Colors["Green"], Colors["Purple"]),
		SphericalMap,
	)

	cone := NewCone()
	cone.LocalShape.(*Cone).Closed = true
	cone.LocalShape.(*Cone).Minimum = 0
	cone.LocalShape.(*Cone).Maximum = 1
	cone.SetTransform(cone.Transform.Compose(
		NewRotateX(math.Pi),
		NewTranslation(5, 1, 0),
		NewScale(1, 2, 1),
	))
	cone.Material.Diffuse = 0.7
	cone.Material.Specular = 0.3
	cone.Material.Pattern = NewTextureMapPattern(
		NewUVCheckerPattern(16, 16, Colors["Green"], Colors["Purple"]),
		SphericalMap,
	)

	cyl := NewCylinder()
	cyl.LocalShape.(*Cylinder).Closed = true
	cyl.LocalShape.(*Cylinder).Minimum = 0
	cyl.LocalShape.(*Cylinder).Maximum = 1
	cyl.Material.Diffuse = 0.7
	cyl.Material.Specular = 0.3
	cyl.Material.Pattern = NewTextureMapPattern(
		NewUVCheckerPattern(16, 16, Colors["Green"], Colors["Purple"]),
		CylindricalMap,
	)
	cyl.SetTransform(cyl.Transform.Compose(
		NewRotateX(math.Pi/2),
		NewRotateY(-math.Pi/2.5),
		NewTranslation(-1.2, 1, -3),
		NewScale(2, 1, 1),
	))

	world := NewWorld()
	world.Objects = []Shape{
		floor,
		backWall,
		sphere,
		cube,
		cone,
		cyl,
	}
	world.Lights = []AreaLight{
		//		NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
		NewAreaLight(NewPoint(3, 5, -5), NewVector(4, 0, 0), 4, NewVector(0, 4, 0), 4, NewColor(1, 1, 1)),
	}

	canvas := camera.RenderWithProgress(world)

	if err := canvas.SaveJPEG("tmp/world.jpg"); err != nil {
		fmt.Printf("Something went wrong! %s\n", err)
	} else {
		fmt.Println("Saved to tmp/world.jpg")
	}
}
