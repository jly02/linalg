// The vectors package provides utility methods
// for performing standard linear-algebraic
// operations on vectors.
package vectors

import "errors"

// Computes the dot product of two vectors.
//
// A successful Dot() call will return err == nil.
func Dot(first []int, second []int) (int, error) {
	// Check if lengths are equal
	if len(first) != len(second) {
		return 0, errors.New("length of each array must be equal")
	}

	var dot int = 0
	for i := 0; i < len(first); i++ {
		dot += first[i] * second[i]
	}

	return dot, nil
}

// Adds two vectors, returning the resulting vector.
//
// A successful Add() call will return err == nil.
func Add(first []int, second []int) ([]int, error) {
	// Check if lengths are equal
	if len(first) != len(second) {
		return nil, errors.New("length of each array must be equal")
	}

	res := make([]int, len(first))
	for i := 0; i < len(first); i++ {
		res[i] = first[i] + second[i]
	}

	return res, nil
}
