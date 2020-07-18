package raytracer

import (
	"fmt"
	"math"
)

type Intersection struct {
	Time   float64
	Object *Shape
	U      float64 // identifies where on 2D triangle that an intersection occurred, relative to corners.
	V      float64 // identifies where on 2D triangle that an intersection occurred, relative to corners.
}

type Intersections []*Intersection

type Computation struct {
	Time       float64 // the moment (in time units) at which the intersection happened
	Object     *Shape  // the object that was intersected
	Point      Tuple   // the point where intersection happened
	OverPoint  Tuple   // the Point value adjusted slightly to avoid "raytracer acne"
	UnderPoint Tuple   //
	EyeV       Tuple   // the vector from eye to the Point
	NormalV    Tuple   // the normal on the object at the given Point
	ReflectV   Tuple   // the vector of reflection
	Inside     bool    // was the intersection inside the object?
	N1         float64 // refractive index of object being exited at ray-object intersection
	N2         float64 // refractive index of object being entered at ray-object intersection
}

func NewIntersection(t float64, obj *Shape) *Intersection {
	return &Intersection{Time: t, Object: obj}
}

func NewIntersectionWithUV(t float64, obj *Shape, u, v float64) *Intersection {
	return &Intersection{Time: t, Object: obj, U: u, V: v}
}

func (i *Intersection) IsEqualTo(i2 *Intersection) bool {
	if i.Time != i2.Time {
		return false
	} else if !i.Object.IsEqualTo(i2.Object) {
		return false
	}
	return true
}

func (i *Intersection) String() string {
	return fmt.Sprintf(
		"Intersection(\nTime: %.3f\nObject: %v\nU: %v, V: %v\n)",
		i.Time,
		i.Object,
		i.U,
		i.V,
	)
}

func (xs Intersections) String() string {
	str := fmt.Sprintf("Intersections(\n")
	for _, x := range xs {
		str += fmt.Sprintf("  %v\n", x)
	}
	str += ")\n"
	return str
}

// TODO: should we just break out the shadow-checking Hit() into a new method?
func (is *Intersections) Hit(checkingForShadows bool) *Intersection {
	var minIntersection *Intersection
	for _, intersection := range *is {
		if intersection.Time > EPSILON && (!checkingForShadows || intersection.Object.Shadows) { // NB the book uses 0, but I had to use this instead to fix the minute difference that broke in TestTheRefractedColorWithARefractedRay's Sphere#LocalIntersect's i2 calcluation
			if minIntersection == nil || minIntersection.IsEqualTo(intersection) || intersection.Time < minIntersection.Time {
				minIntersection = intersection
			}
		}
	}

	return minIntersection
}

// r:  the ray that hit the intersection
// xs: "... the collection of all intersections, which can tell you where the hit is relative to the rest of the intersections ...""
func (i *Intersection) PrepareComputations(r Ray, xs ...*Intersection) *Computation {
	if len(xs) == 0 {
		xs = []*Intersection{i}
	}

	c := &Computation{}
	c.Time = i.Time
	c.Object = i.Object
	c.Point = r.Position(c.Time)
	c.EyeV = r.Direction.Negate()
	c.NormalV = c.Object.NormalAt(c.Point, i)
	c.ReflectV = r.Direction.Reflect(c.NormalV)            // TODO after negating the normal, if necessary
	c.OverPoint = c.Point.Add(c.NormalV.Multiply(EPSILON)) // to avoid "raytracer acne" with shadows
	c.UnderPoint = c.Point.Subtract(c.NormalV.Multiply(EPSILON))

	if c.NormalV.Dot(c.EyeV) < 0 {
		c.Inside = true
		c.NormalV = c.NormalV.Negate()
	} else {
		c.Inside = false
	}

	visitedShapes := []*Shape{}
	for _, intersection := range xs {
		isHit := intersection.IsEqualTo(i) // is this intersection the hit?

		if isHit {
			if len(visitedShapes) == 0 {
				c.N1 = 1.0
			} else {
				c.N1 = visitedShapes[len(visitedShapes)-1].Material.RefractiveIndex
			}
		}

		// TODO move somewhere else?
		indexOf := func(shape *Shape, shapes []*Shape) int {
			for idx, s := range shapes {
				if s.IsEqualTo(shape) {
					return idx
				}
			}
			return -1
		}

		if containerIdx := indexOf(intersection.Object, visitedShapes); containerIdx == -1 { // enter
			visitedShapes = append(visitedShapes, intersection.Object)
		} else { // exit
			visitedShapes = append(visitedShapes[:containerIdx], visitedShapes[containerIdx+1:]...)
		}

		if isHit {
			if len(visitedShapes) == 0 {
				c.N2 = 1.0
			} else {
				c.N2 = visitedShapes[len(visitedShapes)-1].Material.RefractiveIndex
			}
			break
		}
	}

	return c
}

// Schlick's equation is an approximation of Fresnel's.
// Returns the "reflectance", which represents what fraction of the light is reflected at the given hit.
// TODO rename "SchlickReflectance"?
// TODO read “Reflections and Refractions in Ray Tracing” paper to understand this.
func (c *Computation) Schlick() float64 {
	cos := c.EyeV.Dot(c.NormalV) // Cosine of angle between eye and normal vectors

	if c.N1 > c.N2 {
		n := c.N1 / c.N2
		sinSquared := (n * n) * (1 - (cos * cos))
		if sinSquared > 1.0 {
			return 1.0
		}
		cosT := math.Sqrt(1.0 - sinSquared) // Cosine of theta_t using trig identity
		cos = cosT                          // When N1 > N2 use cosT instead of cos
	}

	r0 := math.Pow((c.N1-c.N2)/(c.N1+c.N2), 2)
	return r0 + (1-r0)*math.Pow((1-cos), 5)
}
