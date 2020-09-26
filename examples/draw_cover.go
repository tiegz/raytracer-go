package examples

import (
	"fmt"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawCover(printProgress bool, jobs int) {
	ysf, err := ParseYamlSceneFile("raytracer/files/cover.yml")
	if err != nil {
		panic(err)
	}
	canvas := ysf.Camera.Render(ysf.World, jobs, true)

	if err := canvas.SaveJPEG("tmp/world.jpg"); err != nil {
		fmt.Printf("Something went wrong! %s\n", err)
	} else {
		fmt.Println("Saved to tmp/world.jpg")
	}
}
