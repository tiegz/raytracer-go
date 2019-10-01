package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	"github.com/tiegz/raytracer-go/raytracer"
	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithCylinderAndCone() {
	camera := raytracer.NewCamera(320, 200, math.Pi/3)
	// camera := raytracer.NewCamera(640, 480, math.Pi/3)
	// camera := raytracer.NewCamera(400, 200, math.Pi/3)
	// camera := raytracer.NewCamera(1000, 500, math.Pi/3)
	// camera := raytracer.NewCamera(1920, 1080, math.Pi/3)

	camera.Transform = raytracer.NewViewTransform(
		raytracer.NewPoint(0, 3, -7),
		raytracer.NewPoint(0, 1, 0),
		raytracer.NewVector(0, 1, 0),
	)

	floor := NewPlane()
	floor.Material.Color = NewColor(1, 0.9, 0.9)
	floor.Material.Specular = 0

	cylinder := NewCylinder()
	cylinder.LocalShape.(*Cylinder).Closed = true
	cylinder.LocalShape.(*Cylinder).Minimum = 0
	cylinder.LocalShape.(*Cylinder).Maximum = 1
	cylinder.Transform = cylinder.Transform.Multiply(NewScale(0.5, 3, 0.5))
	cylinder.Material.Color = Colors["Blue"]
	cylinder.Material.Diffuse = 0.7
	cylinder.Material.Specular = 0.3

	cone := NewCone()
	cone.LocalShape.(*Cone).Closed = true
	cone.LocalShape.(*Cone).Minimum = 0
	cone.LocalShape.(*Cone).Maximum = 1
	cone.Transform = NewTranslation(2, 1, 0)
	cone.Transform = cone.Transform.Multiply(NewScale(0.5, 1, 0.5))
	cone.Transform = cone.Transform.Multiply(NewRotateX(math.Pi))
	cone.Material.Color = Colors["Green"]
	cone.Material.Diffuse = 0.7
	cone.Material.Specular = 0.3

	iceCreamCone := NewCone()
	iceCreamCone.LocalShape.(*Cone).Closed = true
	iceCreamCone.LocalShape.(*Cone).Minimum = 0
	iceCreamCone.LocalShape.(*Cone).Maximum = 1
	iceCreamCone.Transform = iceCreamCone.Transform.Multiply(NewTranslation(-1, 0, -3))
	iceCreamCone.Transform = iceCreamCone.Transform.Multiply(NewScale(0.5, 2, 0.5))
	iceCreamCone.Material.Color = NewColor(0.95, 0.95, 0.85)
	// iceCreamCone.Material.Diffuse = 0.7
	// iceCreamCone.Material.Specular = 0.3

	iceCreamScoopOne := NewSphere()
	iceCreamScoopOne.Transform = iceCreamScoopOne.Transform.Multiply(NewTranslation(-1, 2.1, -3))
	iceCreamScoopOne.Transform = iceCreamScoopOne.Transform.Multiply(NewUScale(0.45))
	iceCreamScoopOne.Material.Color = Colors["Red"]

	iceCreamScoopTwo := NewSphere()
	iceCreamScoopTwo.Transform = iceCreamScoopTwo.Transform.Multiply(NewTranslation(-1.2, 2.3, -3))
	iceCreamScoopTwo.Transform = iceCreamScoopTwo.Transform.Multiply(NewUScale(0.3))
	iceCreamScoopTwo.Material.Color = Colors["Orange"]

	iceCreamScoopThree := NewSphere()
	iceCreamScoopThree.Transform = iceCreamScoopThree.Transform.Multiply(NewTranslation(-0.8, 2.3, -3))
	iceCreamScoopThree.Transform = iceCreamScoopThree.Transform.Multiply(NewUScale(0.3))
	iceCreamScoopThree.Material.Color = Colors["DarkRed"]

	world := raytracer.NewWorld()
	world.Objects = []raytracer.Shape{
		floor,
		cylinder,
		cone,
		iceCreamCone,
		iceCreamScoopOne,
		iceCreamScoopTwo,
		iceCreamScoopThree,
	}
	world.Lights = []raytracer.PointLight{
		raytracer.NewPointLight(raytracer.NewPoint(-10, 10, -10), raytracer.NewColor(1, 1, 1)),
	}

	canvas := camera.Render(world)

	fmt.Println("Generating PPM...")
	ppm := canvas.ToPpm()
	filename := "tmp/cube.ppm"
	ppmBytes := []byte(ppm)
	fmt.Printf("Saving scene to %s...\n", filename)
	if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
		panic(err)
	}
}
