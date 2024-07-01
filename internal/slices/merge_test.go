package slices_test

import (
	"slices"
	"testing"

	. "github.com/adzil/eqtest/internal/slices"
)

func TestMerge(t *testing.T) {
	a := []string{"hello", "world"}
	b := []string{"world", "hello"}

	result := Merge(a, []string{})
	if !slices.Equal(a, result) {
		t.Error("resulting slice is not equal")
	}

	result = Merge([]string{}, b)
	if !slices.Equal(b, result) {
		t.Error("resulting slice is not equal")
	}

	result = Merge(a, b)
	if !slices.Equal(slices.Concat(a, b), result) {
		t.Error("resulting slice is not equal")
	}
}
