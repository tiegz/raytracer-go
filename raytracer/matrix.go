package raytracer

type Matrix struct {
	Cols int
	Rows int
	Data []float64
}

func NewMatrix(cols, rows int, data []float64) Matrix {
	return Matrix{cols, rows, data}
}

func (m *Matrix) At(row, col int) float64 {
	return m.Data[row*m.Cols+col]
}
