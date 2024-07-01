package slices

import "slices"

// Merge behaves similarly to [slices.Concat] but more efficient if either a or
// b is empty.
func Merge[E any, S ~[]E](a, b S) S {
	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	return slices.Concat(a, b)
}
