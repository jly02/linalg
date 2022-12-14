// Package linalg vectors module provides utility methods for performing standard
// linear-algebraic operations on vectors. All numeric results are
// returned as 64-bit floats, or vectors of them, and all successful
// function calls will return err == nil.
package linalg

import (
	"errors"
	"math"
)

// EPSILON Arbitrary maximum permissible error in floating point calculations.
const EPSILON float64 = 0.00000001

// Dot Computes the dot product of two vectors.
func Dot[T int | float64, E int | float64](a []T, b []E) (float64, error) {
	// Check if lengths are equal
	if len(a) != len(b) {
		return 0, errors.New("length of each vector must be equal")
	}

	var dot float64 = 0
	for i := range a {
		dot += float64(a[i]) * float64(b[i])
	}

	return dot, nil
}

// IsOrthogonal Determines whether two vectors are orthogonal, I.E. their dot product is zero,
// with an error within EPSILON.
func IsOrthogonal[T int | float64, E int | float64](a []T, b []E) (bool, error) {
	if a == nil || b == nil {
		return false, errors.New("vectors cannot be nil")
	}

	dot, err := Dot(a, b)
	return math.Abs(dot-0) < EPSILON, err
}

// Cross Computes the cross product of two vectors. A cross product is exclusive to 3-dimensional
// vectors.
func Cross[T int | float64, E int | float64](a []T, b []E) ([]float64, error) {
	// Both vectors must be 3D
	if len(a) != 3 || len(b) != 3 {
		return nil, errors.New("length of both vectors must be 3")
	}

	res := make([]float64, 3)
	res[0] = float64(a[1])*float64(b[2]) - float64(a[2])*float64(b[1]) // (a2 * b3) - (a3 * b2)
	res[1] = float64(a[2])*float64(b[0]) - float64(a[0])*float64(b[2]) // (a3 * b1) - (a1 * b3)
	res[2] = float64(a[0])*float64(b[1]) - float64(a[1])*float64(b[0]) // (a1 * b2) - (a2 * b1)

	return res, nil
}

// Add Adds two vectors, returning the resulting vector.
func Add[T int | float64, E int | float64](a []T, b []E) ([]float64, error) {
	// Check if lengths are equal
	if len(a) != len(b) {
		return nil, errors.New("length of each vector must be equal")
	}

	res := make([]float64, len(a))
	for i := range a {
		res[i] = float64(a[i]) + float64(b[i])
	}

	return res, nil
}

// ScalarMult Multiplies a given vector by a scalar value.
func ScalarMult[T int | float64, E int | float64](vec []T, scalar E) ([]float64, error) {
	if len(vec) < 1 {
		return nil, errors.New("vector cannot be empty")
	}

	res := make([]float64, len(vec))
	if scalar == 0 {
		return res, nil
	}

	for i := range vec {
		res[i] = float64(vec[i]) * float64(scalar)
	}

	return res, nil
}
