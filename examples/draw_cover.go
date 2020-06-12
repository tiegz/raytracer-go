package examples

import (
	"fmt"
	"io/ioutil"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawCover() {
	ysf, err := ParseYamlSceneFile("raytracer/files/cover.yml")
	if err != nil {
		panic(err)
	}
	canvas := ysf.Camera.RenderWithProgress(ysf.World)
	fmt.Println("Generating PPM...")
	ppm := canvas.ToPpm()
	filename := "tmp/world.ppm"
	ppmBytes := []byte(ppm)
	fmt.Printf("Saving scene to %s...\n", filename)
	if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
		panic(err)
	}
}
