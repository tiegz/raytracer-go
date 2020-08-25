package main

import (
	"fmt"
	"math"
	"syscall/js"

	assets "github.com/tiegz/raytracer-go/cmd/wasm/assets"

	. "github.com/tiegz/raytracer-go/raytracer"
)

// GOOS=js GOARCH=wasm go build -o assets/raytracer.wasm cmd/wasm/main.go
func main() {
	fmt.Println("Here")

	// Scene
	sphere := NewSphere()

	c, err := NewCanvasFromPpm(assets.EarthMap1k())
	if err != nil {
		panic(err)
	}
	sphere.Material.Pattern = NewTextureMapPattern(
		NewUVImagePattern(c),
		SphericalMap,
	)
	sphere.SetTransform(sphere.Transform.Compose(
		// NewRotateY(0.1),
		// NewRotateX(-0.5),
		NewTranslation(0, 1.1, 0),
	))
	sphere.Material.Diffuse = 0.9
	sphere.Material.Specular = 0.1
	sphere.Material.Shininess = 10
	sphere.Material.Ambient = 0.1
	world := NewWorld()
	world.Lights = []*AreaLight{
		NewPointLight(NewPoint(-2, 2, -10), NewColor(1, 1, 1)),
		NewPointLight(NewPoint(2, 10, -20), NewColor(1, 1, 1)),
	}
	world.Objects = []*Shape{
		sphere,
	}
	camera := NewCamera(50, 50, math.Pi/3)
	camera.SetTransform(NewViewTransform(
		NewPoint(0, 1, -2.2),
		NewPoint(0, 1.1, 0),
		NewVector(0, 1, 0),
	))

	js.Global().Set("setSize", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		width, height := args[0].Int(), args[1].Int()
		camera.HSize = width
		camera.VSize = height
		return true
	}))
	js.Global().Set("moveLeft", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		sphere.SetTransform(sphere.Transform.Compose(NewRotateY(0.1)))
		return true
	}))
	js.Global().Set("moveRight", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		sphere.SetTransform(sphere.Transform.Compose(NewRotateY(-0.1)))
		return true
	}))
	js.Global().Set("renderPixel", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w, h, x, y := args[0].Int(), args[1].Int(), args[2].Int(), args[3].Int()
		camera.HSize = int(math.Max(math.Min(float64(w), 1000), 0))
		camera.VSize = int(math.Max(math.Min(float64(h), 1000), 0))
		r := camera.RayForPixel(x, y)
		color := world.ColorAt(r, 4) // DefaultMaximumReflections
		// fmt.Printf("Called renderPixel(), got %s for %s, %s.\n", color, camera.HSize, camera.VSize)
		return js.ValueOf([]interface{}{int(color.Red * 255), int(color.Green * 255), int(color.Blue * 255)})
	}))
	<-make(chan bool)
}
