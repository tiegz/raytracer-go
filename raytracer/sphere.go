package raytracer

type Sphere struct {
	Origin Tuple
	Radius float64
}

func NewSphere() Sphere { // o Tuple, r float64
	// hardcoding spheres for now
	return Sphere{NewPoint(0, 0, 0), 1.0}
}

// func (r *Ray) Position(time float64) Tuple {
// 	return r.Origin.Add(r.Direction.Multiply(time))
// }
