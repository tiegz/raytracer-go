package raytracer

type PointLight struct {
	Position  Tuple
	Intensity Color
}

func NewPointLight(position Tuple, intensity Color) PointLight {
	return PointLight{position, intensity}
}

func (l *PointLight) IsEqualTo(l2 PointLight) bool {
	if !l.Position.IsEqualTo(l2.Position) {
		return false
	} else if !l.Intensity.IsEqualTo(l2.Intensity) {
		return false
	}
	return true
}
