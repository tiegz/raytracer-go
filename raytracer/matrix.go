package raytracer

type Matrix struct {
	Rows int
	Cols int
	Data []float64
}

func NewMatrix(rows, cols int, data []float64) Matrix {
	return Matrix{rows, cols, data}
}

func (m *Matrix) At(row, col int) float64 {
	return m.Data[row*m.Cols+col]
}

func (m *Matrix) Set(row, col int, val float64) {
	m.Data[row*m.Cols+col] = val
}

func (m *Matrix) IsEqualTo(m2 Matrix) bool {
	if m.Cols != m2.Cols || m.Rows != m2.Rows {
		return false
	}

	for col := 0; col < m.Cols; col = col + 1 {
		for row := 0; row < m.Rows; row = row + 1 {
			if m.At(row, col) != m2.At(row, col) {
				return false
			}
		}
	}

	return true
}

func (m *Matrix) Identity() Matrix {
	// Only implementing a 4x4 identity matrix.
	return NewMatrix(m.Rows, m.Cols, []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	})
}

func (m *Matrix) Transpose() Matrix {
	m2 := NewMatrix(m.Rows, m.Cols, make([]float64, m.Rows*m.Cols, m.Rows*m.Cols))

	for r := 0; r < m.Rows; r += 1 {
		for c := 0; c < m.Cols; c += 1 {
			m2.Set(r, c, m.At(c, r))
		}
	}

	return m2
}

func (m *Matrix) Multiply(m2 Matrix) Matrix {
	m3 := NewMatrix(m.Rows, m.Cols, make([]float64, m.Rows*m.Cols, m.Rows*m.Cols))

	for row := 0; row < m.Rows; row += 1 {
		for col := 0; col < m.Cols; col += 1 {

			// Note: this would only work with square matrices. (pg 29)
			sum := 0.0
			for i := 0; i < m.Rows; i += 1 {
				sum += m.At(row, i) * m2.At(i, col)
			}

			m3.Set(row, col, sum)
		}
	}

	return m3
}

func (m *Matrix) MutiplyByTuple(t Tuple) Tuple {
	// Similar to mulitplying matrices, except there's only one col in a Tuple.
	x := m.At(0, 0)*t.X +
		m.At(0, 1)*t.Y +
		m.At(0, 2)*t.Z +
		m.At(0, 3)*t.W

	y := m.At(1, 0)*t.X +
		m.At(1, 1)*t.Y +
		m.At(1, 2)*t.Z +
		m.At(1, 3)*t.W

	z := m.At(2, 0)*t.X +
		m.At(2, 1)*t.Y +
		m.At(2, 2)*t.Z +
		m.At(2, 3)*t.W

	w := m.At(3, 0)*t.X +
		m.At(3, 1)*t.Y +
		m.At(3, 2)*t.Z +
		m.At(3, 3)*t.W

	return Tuple{x, y, z, w}
}
