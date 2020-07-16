package examples

import (
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithCylinderAndCone() {
	Draw("tmp/world.jpg", func(world *World, camera *Camera) {
		camera.SetSize(320, 200, math.Pi/3)

		camera.SetTransform(NewViewTransform(
			NewPoint(0, 3, -7),
			NewPoint(0, 1, 0),
			NewVector(0, 1, 0),
		))

		floor := NewPlane()
		floor.Material.Color = NewColor(1, 0.9, 0.9)
		floor.Material.Specular = 0

		cylinder := NewCylinder()
		cylinder.LocalShape.(*Cylinder).Closed = true
		cylinder.LocalShape.(*Cylinder).Minimum = 0
		cylinder.LocalShape.(*Cylinder).Maximum = 1
		cylinder.SetTransform(cylinder.Transform.Multiply(NewScale(0.5, 3, 0.5)))
		cylinder.Material.Color = Colors["Blue"]
		cylinder.Material.Diffuse = 0.7
		cylinder.Material.Specular = 0.3

		cone := NewCone()
		cone.LocalShape.(*Cone).Closed = true
		cone.LocalShape.(*Cone).Minimum = 0
		cone.LocalShape.(*Cone).Maximum = 1
		cone.SetTransform(NewTranslation(2, 1, 0))
		cone.SetTransform(cone.Transform.Multiply(NewScale(0.5, 1, 0.5)))
		cone.SetTransform(cone.Transform.Multiply(NewRotateX(math.Pi)))
		cone.Material.Color = Colors["Green"]
		cone.Material.Diffuse = 0.7
		cone.Material.Specular = 0.3

		iceCreamCone := NewCone()
		iceCreamCone.LocalShape.(*Cone).Closed = true
		iceCreamCone.LocalShape.(*Cone).Minimum = 0
		iceCreamCone.LocalShape.(*Cone).Maximum = 1
		iceCreamCone.SetTransform(iceCreamCone.Transform.Multiply(NewTranslation(-1, 0, -3)))
		iceCreamCone.SetTransform(iceCreamCone.Transform.Multiply(NewScale(0.5, 2, 0.5)))
		iceCreamCone.Material.Color = NewColor(0.95, 0.95, 0.85)
		// iceCreamCone.Material.Diffuse = 0.7
		// iceCreamCone.Material.Specular = 0.3

		iceCreamScoopOne := NewSphere()
		iceCreamScoopOne.SetTransform(iceCreamScoopOne.Transform.Multiply(NewTranslation(-1, 2.1, -3)))
		iceCreamScoopOne.SetTransform(iceCreamScoopOne.Transform.Multiply(NewUScale(0.45)))
		iceCreamScoopOne.Material.Color = Colors["Red"]

		iceCreamScoopTwo := NewSphere()
		iceCreamScoopTwo.SetTransform(iceCreamScoopTwo.Transform.Multiply(NewTranslation(-1.2, 2.3, -3)))
		iceCreamScoopTwo.SetTransform(iceCreamScoopTwo.Transform.Multiply(NewUScale(0.3)))
		iceCreamScoopTwo.Material.Color = Colors["Orange"]

		iceCreamScoopThree := NewSphere()
		iceCreamScoopThree.SetTransform(iceCreamScoopThree.Transform.Multiply(NewTranslation(-0.8, 2.3, -3)))
		iceCreamScoopThree.SetTransform(iceCreamScoopThree.Transform.Multiply(NewUScale(0.3)))
		iceCreamScoopThree.Material.Color = Colors["DarkRed"]

		world.Objects = []Shape{
			floor,
			cylinder,
			cone,
			iceCreamCone,
			iceCreamScoopOne,
			iceCreamScoopTwo,
			iceCreamScoopThree,
		}
		world.Lights = []AreaLight{
			NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
		}
	})
}
