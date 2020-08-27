package examples

import (
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

// Adapted from https://github.com/jamis/rtc-ocaml/blob/master/progs/chap12.ml
// TODO: include NewSolidPattern() in here
func RunDrawWorldWithTable(jobs int) {
	Draw(jobs, "tmp/world.jpg", func(world *World, camera *Camera) {
		camera.HSize = 1600
		camera.VSize = 800
		camera.FieldOfView = math.Pi / 4

		camera.SetTransform(NewViewTransform(
			NewPoint(8, 6, -8),
			NewPoint(0, 3, 0),
			NewVector(0, 1, 0),
		))

		floorCeiling := NewCube()
		floorCeiling.SetTransform(floorCeiling.Transform.Compose(
			NewTranslation(0, 1, 0),
			NewScale(20, 7, 20),
		))
		floorCeiling.Material.Pattern = NewCheckerPattern(
			NewColor(0, 0, 0),
			NewColor(0.25, 0.25, 0.25),
		)
		floorCeiling.Material.Pattern.SetTransform(floorCeiling.Material.Pattern.Transform.Multiply(NewUScale(0.07)))
		floorCeiling.Material.Ambient = 0.25
		floorCeiling.Material.Diffuse = 0.7
		floorCeiling.Material.Specular = 0.9
		floorCeiling.Material.Shininess = 300
		floorCeiling.Material.Reflective = 0.1

		walls := NewCube()
		walls.SetTransform(NewScale(10, 10, 10))
		walls.Material.Pattern = NewCheckerPattern(
			NewColor(0.4863, 0.3765, 0.2941),
			NewColor(0.3725, 0.2902, 0.2275),
		)
		walls.Material.Pattern.SetTransform(walls.Material.Pattern.Transform.Multiply(NewScale(0.05, 20, 0.05)))
		walls.Material.Ambient = 0.1
		walls.Material.Diffuse = 0.7
		walls.Material.Specular = 0.9
		walls.Material.Shininess = 300
		walls.Material.Reflective = 0.1
		walls.Material.Pattern.SetTransform(NewScale(0.05, 20, 0.05))

		tabletop := NewCube()
		tabletop.SetTransform(tabletop.Transform.Compose(NewScale(3, 0.1, 2), NewTranslation(0, 3.1, 0)))
		tabletop.Material.Pattern = NewStripePattern(NewColor(0.5529, 0.4235, 0.3255), NewColor(0.6588, 0.5098, 0.4000))
		tabletop.Material.Pattern.SetTransform(tabletop.Material.Pattern.Transform.Compose(NewRotateY(0.1), NewScale(0.05, 0.05, 0.05)))
		tabletop.Material.Ambient = 0.1
		tabletop.Material.Diffuse = 0.7
		tabletop.Material.Specular = 0.9
		tabletop.Material.Shininess = 300
		tabletop.Material.Reflective = 0.2

		legMaterial := DefaultMaterial()
		legMaterial.Color = NewColor(0.5529, 0.4235, 0.3255)
		legMaterial.Ambient = 0.2
		legMaterial.Diffuse = 0.7

		leg1 := NewCube()
		leg1.SetTransform(leg1.Transform.Compose(NewScale(0.1, 1.5, 0.1), NewTranslation(2.7, 1.5, -1.7)))
		leg1.Material = legMaterial

		leg2 := NewCube()
		leg2.SetTransform(leg2.Transform.Compose(NewScale(0.1, 1.5, 0.1), NewTranslation(2.7, 1.5, 1.7)))
		leg2.Material = legMaterial

		leg3 := NewCube()
		leg3.SetTransform(leg3.Transform.Compose(NewScale(0.1, 1.5, 0.1), NewTranslation(-2.7, 1.5, -1.7)))
		leg3.Material = legMaterial

		leg4 := NewCube()
		leg4.SetTransform(leg4.Transform.Compose(NewScale(0.1, 1.5, 0.1), NewTranslation(-2.7, 1.5, 1.7)))
		leg4.Material = legMaterial

		glassCube := NewCube()
		glassCube.SetTransform(glassCube.Transform.Compose(NewUScale(0.25), NewRotateY(0.2), NewTranslation(0, 3.450001, 0)))
		glassCube.Material.Color = NewColor(1, 1, 0.8)
		glassCube.Material.Diffuse = 0.3
		glassCube.Material.Ambient = 0
		glassCube.Material.Specular = 0.9
		glassCube.Material.Shininess = 300
		glassCube.Material.Reflective = 0.7
		glassCube.Material.Transparency = 0.7
		glassCube.Material.RefractiveIndex = 1.5

		littleCube1 := NewCube()
		littleCube1.SetTransform(littleCube1.Transform.Compose(NewUScale(0.15), NewRotateY(-0.4), NewTranslation(1, 3.35, -0.9)))
		littleCube1.Material.Color = NewColor(1, 0.5, 0.5)
		littleCube1.Material.Diffuse = 0.4
		littleCube1.Material.Reflective = 0.6

		littleCube2 := NewCube()
		littleCube2.SetTransform(littleCube2.Transform.Compose(NewScale(0.15, 0.07, 0.15), NewRotateY(0.4), NewTranslation(-1.5, 3.27, 0.3)))
		littleCube2.Material.Color = NewColor(1, 1, 0.5)

		littleCube3 := NewCube()
		littleCube3.SetTransform(littleCube3.Transform.Compose(NewScale(0.2, 0.05, 0.05), NewRotateY(0.4), NewTranslation(0, 3.25, 1)))
		littleCube3.Material.Color = NewColor(0.5, 1, 0.5)

		littleCube4 := NewCube()
		littleCube4.SetTransform(littleCube4.Transform.Compose(NewScale(0.05, 0.2, 0.05), NewRotateY(0.8), NewTranslation(-0.6, 3.4, -1)))
		littleCube4.Material.Color = NewColor(0.5, 0.5, 1)

		littleCube5 := NewCube()
		littleCube5.SetTransform(littleCube5.Transform.Compose(NewScale(0.05, 0.2, 0.05), NewRotateY(0.8), NewTranslation(2, 3.4, 1)))
		littleCube5.Material.Color = NewColor(0.5, 1, 1)

		frame1 := NewCube()
		frame1.SetTransform(frame1.Transform.Compose(NewScale(0.05, 1, 1), NewTranslation(-10, 4, 1)))
		frame1.Material.Color = NewColor(0.7098, 0.2471, 0.2196)
		frame1.Material.Diffuse = 0.6

		frame2 := NewCube()
		frame2.SetTransform(frame2.Transform.Compose(NewScale(0.05, 0.4, 0.4), NewTranslation(-10, 3.4, 2.7)))
		frame2.Material.Color = NewColor(0.2667, 0.2706, 0.6902)
		frame2.Material.Diffuse = 0.6

		frame3 := NewCube()
		frame3.SetTransform(frame3.Transform.Compose(NewScale(0.05, 0.4, 0.4), NewTranslation(-10, 4.6, 2.7)))
		frame3.Material.Color = NewColor(0.3098, 0.5961, 0.3098)
		frame3.Material.Diffuse = 0.6

		mirrorFrame := NewCube()
		mirrorFrame.SetTransform(mirrorFrame.Transform.Compose(NewScale(5, 1.5, 0.05), NewTranslation(-2, 3.5, 9.95)))
		mirrorFrame.Material.Color = NewColor(0.3882, 0.2627, 0.1882)
		mirrorFrame.Material.Diffuse = 0.7

		mirror := NewCube()
		mirror.SetTransform(mirror.Transform.Compose(NewScale(4.8, 1.4, 0.06), NewTranslation(-2, 3.5, 9.95)))
		mirror.Material.Color = Colors["Black"]
		mirror.Material.Ambient = 0
		mirror.Material.Diffuse = 0
		mirror.Material.Specular = 1
		mirror.Material.Shininess = 300
		mirror.Material.Reflective = 1

		world.Objects = []*Shape{
			floorCeiling,
			walls,
			tabletop,
			leg1,
			leg2,
			leg3,
			leg4,
			glassCube,
			littleCube1,
			littleCube2,
			littleCube3,
			littleCube4,
			littleCube5,
			frame1,
			frame2,
			frame3,
			mirrorFrame,
			mirror,
		}
		world.Lights = []*AreaLight{
			NewPointLight(NewPoint(0, 6.9, -5), NewColor(1, 1, 1)),
		}
	})
}
