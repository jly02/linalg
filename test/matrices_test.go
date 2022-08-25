package test

import (
	"testing"

	matrices "github.com/jly02/linalg/mat"
)

func TestEyeLessThanOne(t *testing.T) {
	mat, err := matrices.Eye(0)

	if err == nil {
		t.Fatalf(`Eye(0) = %f, %v, want 0`, mat, err)
	}
}
