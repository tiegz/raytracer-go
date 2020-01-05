package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawWorldWithSnowman() {
	camera := NewCamera(640, 480, math.Pi/3)
	// camera := NewCamera(1280, 960, math.Pi/3)

	camera.SetTransform(NewViewTransform(
		NewPoint(-2, 2.5, -6),
		NewPoint(0, 2.2, 0),
		NewVector(0, 1, 0),
	))

	snowMaterial := DefaultMaterial()
	snowMaterial.Specular = 0.1
	snowMaterial.Diffuse = 0.8
	snowMaterial.Ambient = 0.4

	floor := NewPlane()
	floor.Material = snowMaterial
	// floor.Material.Pattern = NewCheckerPattern(Colors["White"], Colors["Gray"])

	backWall := NewPlane()
	backWall.SetTransform(backWall.Transform.Compose(
		NewRotateX(math.Pi/2),
		NewTranslation(0, 0, 80),
	))
	backWall.Material.Color = Colors["White"]
	// backWall.Material.Pattern = NewCheckerPattern(Colors["White"], Colors["Gray"])

	createNose := func() *Shape {
		nose := NewCone()
		nose.Material.Color = Colors["Orange"]
		nose.SetTransform(nose.Transform.Compose(
			NewScale(0.05, 0.4, 0.05),
			NewRotateX(math.Pi/2),
			NewShear(0.1, 0, 0, 0, 0, 0),
			NewTranslation(-0.05, 3.4, -1),
		))
		nose.LocalShape.(*Cone).Closed = true
		nose.LocalShape.(*Cone).Minimum = 0.0
		nose.LocalShape.(*Cone).Maximum = 2.0
		return &nose
	}

	createHat := func() *Shape {
		hat := NewGroup()

		hatMaterial := DefaultMaterial()
		hatMaterial.Color = NewColor(0.1, 0.1, 0.3)
		hatMaterial.Ambient = 0.4
		hatMaterial.Specular = 0.3
		hatMaterial.Diffuse = 0.3

		topOne := NewCone()
		topOne.Material = hatMaterial
		topOne.SetTransform(topOne.Transform.Compose(
			NewScale(0.35, 1, 0.4),
			NewRotateX(math.Pi),
			NewTranslation(0, 2, 0),
		))
		topOne.LocalShape.(*Cone).Closed = true
		topOne.LocalShape.(*Cone).Minimum = 0.0
		topOne.LocalShape.(*Cone).Maximum = 2.0
		topTwo := NewSphere()
		topTwo.Material = hatMaterial
		topTwo.SetTransform(topTwo.Transform.Compose(
			NewTranslation(0, 1.7, 0),
			NewScale(1.2, 1, 1.4),
		))
		topTwo.Shadows = false
		top := NewCsg("difference", &topOne, &topTwo)

		band := NewCylinder()
		band.Material.Color = NewColor(0.7, 0, 0)
		band.Material.Specular = 0.75
		band.Material.Ambient = 0.50
		band.Material.RefractiveIndex = 0.5
		band.Material.Reflective = 0.5
		band.SetTransform(band.Transform.Compose(
			NewScale(0.65, 1, 0.75),
			NewTranslation(0, 0.12, 0),
		))
		band.LocalShape.(*Cylinder).Closed = true
		band.LocalShape.(*Cylinder).Minimum = 0.0
		band.LocalShape.(*Cylinder).Maximum = 0.15

		rim := NewCylinder()
		rim.Material = hatMaterial
		rim.LocalShape.(*Cylinder).Closed = true
		rim.LocalShape.(*Cylinder).Minimum = 0.0
		rim.LocalShape.(*Cylinder).Maximum = 0.02

		hat.AddChildren(&rim, &band, &top)
		return &hat
	}

	createArm := func() *Shape {
		arm := NewGroup()

		stickOne := NewCube()
		stickOne.Material.Color = NewColor(0.64, 0.16, 0.16)
		stickOne.SetTransform(stickOne.Transform.Compose(
			NewScale(0.6, 0.03, 0.03),
			NewRotateZ(-math.Pi/8),
			NewTranslation(-1, 3, 0),
		))
		stickTwo := NewCube()
		stickTwo.Material.Color = NewColor(0.54, 0.16, 0.16)
		stickTwo.SetTransform(stickTwo.Transform.Compose(
			NewScale(0.3, 0.03, 0.03),
			NewRotateZ(-math.Pi/4),
			NewTranslation(-1.7, 3.39, 0),
		))
		stickThree := NewCube()
		stickThree.Material.Color = NewColor(0.54, 0.16, 0.16)
		stickThree.SetTransform(stickThree.Transform.Compose(
			NewScale(0.2, 0.03, 0.03),
			NewRotateZ(math.Pi/4),
			NewTranslation(-1.6, 3.1, 0),
		))

		arm.AddChildren(&stickOne, &stickTwo, &stickThree)

		return &arm
	}

	createDot := func() *Shape {
		sphere := NewSphere()
		sphere.Shadows = false
		sphere.Material.Specular = 0.2
		sphere.Material.Color = NewColor(0.1, 0.1, 0.1)
		sphere.SetTransform(sphere.Transform.Compose(
			NewScale(0.05, 0.05, 0.05),
			NewTranslation(0, 1, 0),
		))
		return &sphere
	}

	createBall := func() *Shape {
		sphere := NewSphere()
		sphere.SetTransform(sphere.Transform.Compose(
			NewTranslation(0, 1, 0)),
		)
		sphere.Material.Color = NewColor(1, 1, 1)
		return &sphere
	}

	createSnowman := func() *Shape {
		snowman := NewGroup()
		bottomBall := createBall()
		bottomBall.Material = snowMaterial
		middleBall := createBall()
		middleBall.SetTransform(middleBall.Transform.Compose(
			NewUScale(0.8),
			NewTranslation(0, 1.5, 0),
		))
		middleBall.Material = snowMaterial
		topBall := createBall()
		topBall.SetTransform(topBall.Transform.Compose(
			NewUScale(0.6),
			NewTranslation(0, 2.8, 0),
		))
		topBall.Material = snowMaterial

		leftEye := createDot()
		leftEye.SetTransform(leftEye.Transform.Compose(
			NewTranslation(-0.2, 2.6, -0.53),
		))
		rightEye := createDot()
		rightEye.SetTransform(rightEye.Transform.Compose(
			NewTranslation(0.1, 2.6, -0.56),
		))

		mouthCoords := [][]float64{
			[]float64{-0.35, 2.35, -0.48},
			[]float64{-0.25, 2.25, -0.51},
			[]float64{-0.15, 2.20, -0.54},
			[]float64{-0.05, 2.20, -0.54},
			[]float64{0.05, 2.20, -0.54},
			[]float64{0.15, 2.25, -0.53},
			[]float64{0.25, 2.35, -0.52},
		}
		for _, coords := range mouthCoords {
			mouthDot := createDot()
			mouthDot.SetTransform(mouthDot.Transform.Compose(
				NewTranslation(coords[0], coords[1], coords[2]),
			))
			snowman.AddChildren(mouthDot)
		}

		buttonCoords := [][]float64{
			[]float64{0, 1.60, -0.66},
			[]float64{0, 1.40, -0.74},
			[]float64{0, 1.2, -0.77},
		}
		for _, coords := range buttonCoords {
			buttonDot := createDot()
			buttonDot.SetTransform(buttonDot.Transform.Compose(
				NewUScale(1.1),
				NewTranslation(coords[0], coords[1], coords[2]),
			))
			snowman.AddChildren(buttonDot)
		}

		snowman.AddChildren(bottomBall, middleBall, topBall, leftEye, rightEye)

		return &snowman
	}

	snowman := createSnowman()
	hat := createHat()
	hat.SetTransform(hat.Transform.Compose(
		NewUScale(0.5),
		NewRotateZ(math.Pi/8),
		NewTranslation(-0.22, 3.9, 0),
	))
	nose := createNose()

	leftArm := createArm()
	rightArm := createArm()
	rightArm.SetTransform(rightArm.Transform.Compose(
		NewRotateY(math.Pi),
	))

	world := NewWorld()
	world.Objects = []Shape{
		floor,
		backWall,
		*nose,
		*leftArm,
		*rightArm,
		*snowman,
		*hat,
	}
	world.Lights = []AreaLight{
		// NewPointLight(NewPoint(0, 30, 70), NewColor(1, 1, 1)),
		NewAreaLight(NewPoint(0, 5, -2), NewVector(4, 0, 0), 16, NewVector(0, 4, 0), 16, NewColor(1, 1, 1)),
	}

	canvas := camera.RenderWithProgress(world)

	fmt.Println("Generating PPM...")
	ppm := canvas.ToPpm()
	filename := "tmp/world.ppm"
	ppmBytes := []byte(ppm)
	fmt.Printf("Saving scene to %s...\n", filename)
	if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
		panic(err)
	}
}
