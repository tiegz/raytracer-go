package raytracer

func NewTranslation(x, y, z float64) Matrix {
	translation := IdentityMatrix()
	translation.Set(0, 3, x)
	translation.Set(1, 3, y)
	translation.Set(2, 3, z)
	return translation
}

func NewScale(x, y, z float64) Matrix {
	scale := NewMatrix(4, 4, []float64{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	})
	return scale
}
