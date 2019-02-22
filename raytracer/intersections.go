package raytracer

type Intersection struct {
	T      float64 // aka "time"
	Object Sphere
}
type Intersections []Intersection

func NewIntersection(t float64, obj Sphere) Intersection {
	return Intersection{t, obj}
}
