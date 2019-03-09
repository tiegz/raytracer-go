package raytracer

import (
	"fmt"
	"sort"
)

type World struct {
	Objects []Sphere
	Lights  []PointLight
}

// NewWorld instantiates a new World object.
func NewWorld() World {
	return World{}
}

// DefaultWorld returns a new world with some default settings:
//   * 1 unit sphere with color
//   * 1 smaller sphere inside ^
//   * a single white light
func DefaultWorld() World {
	w := NewWorld()

	defaultPointLight := NewPointLight(NewPoint(-10, 10, -10), Colors["White"])
	defaultObj1 := NewSphere()
	defaultObj1.Material.Color = NewColor(0.8, 1.0, 0.6)
	defaultObj1.Material.Diffuse = 0.7
	defaultObj1.Material.Specular = 0.2

	defaultObj2 := NewSphere()
	defaultObj2.Transform = NewScale(0.5, 0.5, 0.5)

	w.Objects = append(w.Objects, defaultObj1, defaultObj2)
	w.Lights = []PointLight{defaultPointLight}
	return w
}

func (w *World) String() string {
	return fmt.Sprintf("World( %d Objects %d Lights )", len(w.Objects), len(w.Lights))
}

// Contains returns true if the world contains obj.
func (w *World) Contains(obj Sphere) bool {
	for _, o := range w.Objects {
		if o.IsEqualTo(obj) {
			return true
		}
	}
	return false
}

func (w *World) Intersect(r Ray) Intersections {
	var xs Intersections

	for _, obj := range w.Objects {
		xs = append(xs, r.Intersect(obj)...)
	}
	// Sort the intersections by time, so we can get the first hit. (p 97)
	sort.Slice(xs, func(i, j int) bool { return xs[i].Time < xs[j].Time })

	return xs
}

func (w *World) ShadeHit(c Computation) Color {
	fmt.Printf("Lights are %v\n", w.Lights)
	color := NewColor(0, 0, 0)

	for _, light := range w.Lights {
		color = color.Add(c.Object.Material.Lighting(light, c.Point, c.EyeV, c.NormalV))
	}

	return color
}
