package raytracer

import (
	"fmt"
	"math"
	"sort"
)

// TODO ugh can we make this a default argument somehow?
const DefaultMaximumReflections int = 4 // recommended on page 148

type World struct {
	Objects []*Shape
	Lights  []*AreaLight
}

// NewWorld instantiates a new World object.
func NewWorld() *World {
	return &World{}
}

// DefaultWorld returns a new world with some default settings:
//   * 1 unit sphere with color
//   * 1 smaller sphere inside ^
//   * a single white light
func DefaultWorld() *World {
	w := NewWorld()

	defaultPointLight := NewPointLight(NewPoint(-10, 10, -10), Colors["White"])
	defaultObj1 := NewSphere()
	defaultObj1.Material.Color = NewColor(0.8, 1.0, 0.6)
	defaultObj1.Material.Diffuse = 0.7
	defaultObj1.Material.Specular = 0.2

	defaultObj2 := NewSphere()
	defaultObj2.SetTransform(NewScale(0.5, 0.5, 0.5))

	w.Objects = append(w.Objects, defaultObj1, defaultObj2)
	w.Lights = []*AreaLight{defaultPointLight}
	return w
}

func (w *World) String() string {
	return fmt.Sprintf("World(\n  %d Objects\n  %d Lights\n)", len(w.Objects), len(w.Lights))
}

// Contains returns true if the world contains obj.
func (w *World) Contains(obj *Shape) bool {
	for _, o := range w.Objects {
		if o.IsEqualTo(obj) {
			return true
		}
	}
	return false
}

func (w *World) Intersect(r *Ray) Intersections {
	var xs Intersections

	for _, obj := range w.Objects {
		xs = append(xs, obj.Intersect(r)...)
	}
	// Sort the  intersections by time, so we can get the first hit. (p 97)
	sort.Slice(xs, func(i, j int) bool { return xs[i].Time < xs[j].Time })

	return xs
}

// ShadeHit returns the color for the given computation's intersection.
func (w *World) ShadeHit(c *Computation, remainingReflections int) Color {
	color := NewColor(0, 0, 0)

	for _, light := range w.Lights {
		// isShadowed := w.IsShadowed(c.OverPoint, light)
		intensity := light.IntensityAt(c.OverPoint, w)
		surfaceColor := c.Object.Material.Lighting(c.Object, light, c.OverPoint, c.EyeV, c.NormalV, intensity)

		reflectedColor := w.ReflectedColor(c, remainingReflections)
		refractedColor := w.RefractedColor(c, remainingReflections)
		if c.Object.Material.Reflective > 0 && c.Object.Material.Transparency > 0 {
			reflectance := c.SchlickReflectance()
			color = color.
				Add(surfaceColor).
				Add(reflectedColor.Multiply(reflectance)).
				Add(refractedColor.Multiply(1 - reflectance))
		} else {
			color = color.
				Add(surfaceColor).
				Add(reflectedColor).
				Add(refractedColor)
		}
	}

	return color
}

func (w *World) ReflectedColor(c *Computation, remainingReflections int) Color {
	if remainingReflections < 1 {
		return Colors["Black"]
	} else if c.Object.Material.Reflective == 0 {
		return Colors["Black"]
	} else {
		reflectionRay := NewRay(c.OverPoint, c.ReflectV)
		return w.ColorAt(reflectionRay, remainingReflections-1).
			Multiply(c.Object.Material.Reflective)
	}
}

// ColorAt gets a ray's intersection in the world and returns that intersection's color.
func (w *World) ColorAt(r *Ray, remainingReflections int) Color {
	var color Color

	// 	Call intersect_world to find the intersections of the given ray with the given world.
	is := w.Intersect(r)

	// 2. Find the hit from the resulting intersections.
	if hit := is.Hit(false); hit == nil {

		// 3. Return the color black if there is no such intersection.
		color = Colors["Black"]
	} else {
		// 4. Otherwise, precompute the necessary values with prepare_computations.
		c := hit.PrepareComputations(r, is...)

		// 5. Finally, call shade_hit to find the color at the hit.
		color = w.ShadeHit(c, remainingReflections)
	}

	return color
}

func (w *World) RefractedColor(c *Computation, remaining int) Color { // remaining
	if remaining == 0 || c.Object.Material.Transparency == 0 {
		return Colors["Black"]
	}

	// Check for Total Internal Reflection using Snell's Law (p 157)
	// ... a phenomenon that occurs when light enters a new medium at a sufficiently acute angle, and the new medium has a lower refractive index than the old ...
	nRatio := c.N1 / c.N2                                 // Find the ratio of first index of refraction to the second.
	cosI := c.EyeV.Dot(c.NormalV)                         // cos(theta_i) is the same as the dot product of the two vectors
	sinSquared := (nRatio * nRatio) * (1 - (cosI * cosI)) // Find sin(theta_t)^2 via trigonometric identity
	if sinSquared > 1 {                                   // total I
		return Colors["Black"]
	}

	// ... Show that refracted_color() in all other cases will spawn a secondary ray in the correct direction, and return its color. ...
	cosT := math.Sqrt(1 - sinSquared) // Find cos(theta_t) via trigonometric identity

	normalScaled := c.NormalV.Multiply(nRatio*cosI - cosT)
	eyeScaled := c.EyeV.Multiply(nRatio)
	direction := normalScaled.Subtract(eyeScaled)   // Compute the direction of the refracted ray
	refractedRay := NewRay(c.UnderPoint, direction) // The refracted ray

	// Find the color of the refracted ray, making sure to multiply # by the transparency value to account for any opacity
	color := w.ColorAt(refractedRay, remaining-1.0).
		Multiply(c.Object.Material.Transparency)

	return color
}

func (w *World) IsShadowed(p Tuple, lightPosition Tuple) bool {
	// TODO: do this for multiple light sources?
	v := lightPosition.Subtract(p)
	distance := v.Magnitude()
	direction := v.Normalized()
	r := NewRay(p, direction)
	is := w.Intersect(r)

	if hit := is.Hit(true); hit == nil {
		return false
	} else {
		// is the intersection between the point and the light?
		return hit.Time < distance
	}
}
