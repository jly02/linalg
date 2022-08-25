package linalg

import "errors"

// Produces an identity matrix of a given size.
func Eye(size int) ([][]float64, error) {
	if size < 1 {
		return nil, errors.New("matrix cannot have size < 1")
	}

	mat := make([][]float64, size)
	for i := range mat {
		mat[i] = make([]float64, size)
		mat[i][i] = 1
	}

	return mat, nil
}
