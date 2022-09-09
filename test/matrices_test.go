package test

import (
	"testing"

	matrices "github.com/jly02/linalg/pkg/mat"
)

// Compares equality between two float matrices based on arbitrary epsilon value.
func floatMatrixEquals(a, b [][]float64) bool {
	for i := range a {
		if !floatSliceEquals(a[i], b[i]) {
			return false
		}
	}

	return true
}

func TestEyeLessThanOne(t *testing.T) {
	mat, err := matrices.IdMatrix(0)

	if err == nil {
		t.Fatalf(`Eye(0) = %f, %v, want 0`, mat, err)
	}
}

func TestEyeOne(t *testing.T) {
	mat, err := matrices.IdMatrix(1)

	want := [][]float64{{1}}
	if err != nil || !floatMatrixEquals(mat, want) {
		t.Fatalf(`Eye(0) = %f, %v, want %f`, mat, err, want)
	}
}

func TestEyeSmall(t *testing.T) {
	mat, err := matrices.IdMatrix(4)

	want := [][]float64{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	if err != nil || !floatMatrixEquals(mat, want) {
		t.Fatalf(`Eye(4) = %f, %v, want %f`, mat, err, want)
	}
}

func TestTransposeEmpty(t *testing.T) {
	input := make([][]float64, 0)
	mat, err := matrices.Transpose(input)

	if err != nil || len(mat) > 0 {
		t.Fatalf(`Transpose([[]]) = %f, %v`, mat, err)
	}
}

func TestTransposeOne(t *testing.T) {
	input := [][]float64{{1}}
	mat, err := matrices.Transpose(input)

	want := [][]float64{{1}}
	if err != nil || !floatMatrixEquals(mat, want) {
		t.Fatalf(`Transpose([[1]]) = %f, %v`, mat, err)
	}
}

func TestTransposeSquare(t *testing.T) {
	input := [][]float64{
		{0, 1, 1, 1},
		{0, 0, 1, 1},
		{0, 0, 0, 1},
		{0, 0, 0, 0},
	}
	mat, err := matrices.Transpose(input)

	want := [][]float64{
		{0, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 1, 0, 0},
		{1, 1, 1, 0},
	}
	if err != nil || !floatMatrixEquals(mat, want) {
		t.Fatalf(`Transpose([4][4]) = %f, %v`, mat, err)
	}
}

func TestTransposeRectangle(t *testing.T) {
	input := [][]float64{
		{0, 0, 1, 1},
		{0, 0, 0, 1},
	}
	mat, err := matrices.Transpose(input)

	want := [][]float64{
		{0, 0},
		{0, 0},
		{1, 0},
		{1, 1},
	}
	if err != nil || !floatMatrixEquals(mat, want) {
		t.Fatalf(`Transpose([2][4]) = %f, %v`, mat, err)
	}
}
