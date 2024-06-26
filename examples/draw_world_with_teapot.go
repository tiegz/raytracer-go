package examples

import (
	"io/ioutil"
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithTeapot(printProgress bool, jobs int) {
	Draw(printProgress, jobs, "tmp/world.jpg", func(world *World, camera *Camera) {
		camera.HSize = 640
		camera.VSize = 400
		camera.FieldOfView = math.Pi / 3

		camera.SetTransform(NewViewTransform(
			NewPoint(0, 20, -20),
			NewPoint(2, 5, 0),
			NewVector(0, 1, 0),
		))

		floor := NewPlane()
		floor.Material.Pattern = NewCheckerPattern(Colors["White"], Colors["Gray"])

		backWall := NewPlane()
		backWall.SetTransform(backWall.Transform.Compose(
			NewRotateX(math.Pi/2),
			NewTranslation(0, 0, 20),
		))
		backWall.Material.Pattern = NewCheckerPattern(Colors["White"], Colors["Gray"])

		dat, err := ioutil.ReadFile("raytracer/files/utah_teapot_hires.obj") // ("raytracer/files/mit_teapot.obj")
		if err != nil {
			panic(err)
		}

		objFile := ParseObjFile(string(dat))
		group := objFile.ToGroup()
		group.SetTransform(group.Transform.Compose(
			NewRotateX(-math.Pi/2),
			NewUScale(0.75),
		))

		m := DefaultMaterial()
		m.Color = Colors["Purple"]
		group.SetMaterialRecursively(m)

		world.Objects = []*Shape{
			floor,
			backWall,
			group,
		}
		world.Lights = []*AreaLight{
			NewPointLight(NewPoint(0, -10, -5), NewColor(1, 1, 1)),
			NewPointLight(NewPoint(0, 10, -5), NewColor(1, 1, 1)),
		}
	})
}
