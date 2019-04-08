package raytracer

import (
	"fmt"
	"sort"
)

type World struct {
	Objects []Shape
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
func (w *World) Contains(obj Shape) bool {
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
		xs = append(xs, obj.Intersect(r)...)
	}
	// Sort the  intersections by time, so we can get the first hit. (p 97)
	sort.Slice(xs, func(i, j int) bool { return xs[i].Time < xs[j].Time })

	return xs
}

// ShadeHit returns the color for the given computation's intersection.
func (w *World) ShadeHit(c Computation) Color {
	color := NewColor(0, 0, 0)

	for _, light := range w.Lights {
		isShadowed := w.IsShadowed(c.OverPoint)
		result := c.Object.Material.Lighting(light, c.OverPoint, c.EyeV, c.NormalV, isShadowed)
		color = color.Add(result)
	}

	return color
}

// ColorAt gets a ray's intersection in the world and returns that intersection's color.
func (w *World) ColorAt(r Ray) Color {
	var color Color

	// 	Call intersect_world to find the intersections of the given ray with the given world.
	is := w.Intersect(r)

	// 2. Find the hit from the resulting intersections.
	if hit := is.Hit(); hit.IsEqualTo(NullIntersection()) {
		// 3. Return the color black if there is no such intersection.
		color = Colors["Black"]
	} else {
		// 4. Otherwise, precompute the necessary values with prepare_computations.
		c := hit.PrepareComputations(r)

		// 5. Finally, call shade_hit to find the color at the hit.
		color = w.ShadeHit(c)
	}

	return color
}

func (w *World) IsShadowed(p Tuple) bool {
	// TODO enable for more than 1 world light
	v := w.Lights[0].Position.Subtract(p)
	distance := v.Magnitude()
	direction := v.Normalized()
	r := NewRay(p, direction)
	is := w.Intersect(r)

	if hit := is.Hit(); hit.IsEqualTo(NullIntersection()) {
		return false
	} else {
		if hit.Time < distance {
			return true // the intersection is between the point and the light
		} else {
			return false
		}
	}
}
