package util

type Matrix [][]float64
type Vector []float64

func (m Matrix) Rows() int {
	return len(m)
}

func (m Matrix) Cols() int {
	return len(m[0])
}

func NewVector(v ...int) Vector {
	var rows Vector
	for _, val := range v {
		rows = append(rows, float64(val))
	}
	return rows
}

func Solve(m Matrix, v Vector) (Matrix, Vector) {
	for i := 0; i < m.Rows(); i++ {
		norm := m[i][i]
		for c := 0; c < m.Cols(); c++ {
			m[i][c] /= norm
		}
		v[i] /= norm
		for j := 0; j < m.Cols(); j++ {
			if i == j {
				continue
			}
			ratio := m[j][i] / m[i][i]
			for k := i; k < m.Cols(); k++ {
				m[j][k] -= ratio * m[i][k]
			}
			v[j] -= ratio * v[i]
		}
	}
	return m, v
}
