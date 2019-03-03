package raytracer

type PointLight struct {
	Position  Tuple
	Intensity Color
}

func NewPointLight(position Tuple, intensity Color) PointLight {
	return PointLight{position, intensity}
}
