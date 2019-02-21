package raytracer

func NewTranslation(x, y, z float64) Matrix {
	translation := IdentityMatrix()
	translation.Set(0, 3, x)
	translation.Set(1, 3, y)
	translation.Set(2, 3, z)
	return translation
}
